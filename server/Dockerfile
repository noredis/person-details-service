FROM golang:1.24-alpine AS builder

RUN apk --no-cache add bash git gcc gettext musl-dev

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /usr/local/src

COPY ["app/go.mod", "app/go.sum", "./"]

COPY pds.env ./

RUN go mod download

COPY app ./

RUN swag init --dir ./ --output ./docs --generalInfo ./cmd/app/app.go --parseDependency --parseInternal

RUN go build -o ./bin/app cmd/app/app.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /

COPY --from=builder /usr/local/src/pds.env /

CMD ["/app"]
