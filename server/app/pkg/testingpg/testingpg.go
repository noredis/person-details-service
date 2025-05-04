package testingpg

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

type TestingT interface {
	require.TestingT

	Cleanup(f func())
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Failed() bool
}

const defaultPostgresURL = "postgresql://postgres:postgres@localhost:32260/postgres?sslmode=disable"

func NewWithIsolatedDatabase(t TestingT) *Postgres {
	return newPostgres(t, defaultPostgresURL).cloneFromReference()
}

func NewWithIsolatedSchema(t TestingT) *Postgres {
	return newPostgres(t, defaultPostgresURL).createSchema(t)
}

func NewWithTransactionalCleanup(t TestingT) interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
} {
	const databaseName = "transaction"

	postgres := newPostgres(t, defaultPostgresURL)
	postgres = postgres.replaceDBName(databaseName)

	ctx, done := context.WithCancel(context.Background())
	t.Cleanup(done)

	tx, err := postgres.DB().BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadWrite,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, tx.Rollback(ctx))
	})

	return tx
}

type Postgres struct {
	t TestingT

	url string
	ref string

	sqlDB     *pgxpool.Pool
	sqlDBOnce sync.Once
}

func newPostgres(t TestingT, defaultPostgresURL string) *Postgres {
	urlStr := os.Getenv("TESTING_DB_URL")
	if urlStr == "" {
		urlStr = defaultPostgresURL

		const format = "env TESTING_DB_URL is empty, used default value: %s"

		t.Logf(format, urlStr)
	}

	refDatabase := os.Getenv("TESTING_DB_REF")
	if refDatabase == "" {
		refDatabase = "reference"
	}

	return &Postgres{
		t: t,

		url: urlStr,
		ref: refDatabase,
	}
}

func (p *Postgres) URL() string {
	return p.url
}

func (p *Postgres) DB() *pgxpool.Pool {
	p.sqlDBOnce.Do(func() {
		p.sqlDB = open(p.t, p.URL())
	})

	return p.sqlDB
}

func (p *Postgres) createSchema(t TestingT) *Postgres {
	schemaName := newUniqueHumanReadableDatabaseName(p.t)

	schemaName = strings.ToLower(schemaName)

	ctx, done := context.WithCancel(context.Background())
	t.Cleanup(done)

	{
		sql := fmt.Sprintf(`CREATE SCHEMA "%s";`, schemaName)

		_, err := p.DB().Exec(ctx, sql)
		require.NoError(t, err)
	}

	t.Cleanup(func() {
		sql := fmt.Sprintf(`DROP SCHEMA "%s" CASCADE;`, schemaName)

		_, err := p.DB().Exec(ctx, sql)
		require.NoError(t, err)
	})

	pgurl := setSearchPath(t, p.URL(), schemaName)

	return &Postgres{
		t:   p.t,
		ref: p.ref,
		url: pgurl.String(),
	}
}

func (p *Postgres) cloneFromReference() *Postgres {
	newDBName := newUniqueHumanReadableDatabaseName(p.t)

	p.t.Log("database name for this test:", newDBName)

	sql := fmt.Sprintf(
		`CREATE DATABASE %q WITH TEMPLATE %q;`,
		newDBName,
		p.ref,
	)

	_, err := p.DB().Exec(context.Background(), sql)
	require.NoError(p.t, err)

	p.t.Cleanup(func() {
		sql := fmt.Sprintf(`DROP DATABASE %q WITH (FORCE);`, newDBName)

		ctx, done := context.WithTimeout(context.Background(), time.Minute)
		defer done()

		_, err := p.DB().Exec(ctx, sql)
		require.NoError(p.t, err)
	})

	return p.replaceDBName(newDBName)
}

func (p *Postgres) replaceDBName(newDBName string) *Postgres {
	o := p.clone()
	o.url = replaceDBName(p.t, p.URL(), newDBName)

	return o
}

func (p *Postgres) clone() *Postgres {
	return &Postgres{
		t: p.t,

		url: p.url,
		ref: p.ref,
	}
}

func newUniqueHumanReadableDatabaseName(t TestingT) string {
	output := strings.Builder{}

	const maxIdentifierLengthBytes = 63

	uid := genUnique8BytesID(t)
	maxHumanReadableLenBytes := maxIdentifierLengthBytes - len(uid)

	lastSymbolIsHyphen := false

	for _, r := range t.Name() {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			output.WriteRune(r)

			lastSymbolIsHyphen = false
		} else {
			if !lastSymbolIsHyphen {
				output.WriteRune('-')
			}

			lastSymbolIsHyphen = true
		}

		if output.Len() >= maxHumanReadableLenBytes {
			break
		}
	}

	output.WriteString(uid)

	return output.String()
}

func genUnique8BytesID(t TestingT) string {
	bs := make([]byte, 6)

	_, err := rand.Read(bs)
	require.NoError(t, err)

	return base64.RawURLEncoding.EncodeToString(bs)
}

func replaceDBName(t TestingT, dataSourceURL, dbname string) string {
	r, err := url.Parse(dataSourceURL)
	require.NoError(t, err)

	r.Path = dbname

	return r.String()
}

func open(t TestingT, dataSourceURL string) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), dataSourceURL)
	require.NoError(t, err)

	t.Cleanup(func() {
		db.Close()
	})

	return db
}

func setSearchPath(t TestingT, pgURL string, schemaName string) *url.URL {
	pgurl, err := url.Parse(pgURL)
	require.NoError(t, err)

	query := pgurl.Query()
	query.Set("search_path", schemaName)
	pgurl.RawQuery = query.Encode()

	return pgurl
}
