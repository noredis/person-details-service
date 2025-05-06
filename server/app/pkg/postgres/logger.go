package postgres

type Logger interface {
	Error(msg string, args ...any)
}

type FakeLogger struct{}

func (FakeLogger) Error(msg string, args ...any) {}
