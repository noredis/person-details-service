FROM golang:1.24-alpine AS builder

RUN apk --no-cache add bash git gcc gettext musl-dev

WORKDIR /usr/local/src

COPY ["app/go.mod", "app/go.sum", "./"]

COPY pds.env ./

RUN go mod download

COPY app ./

COPY configs ./configs

RUN go build -o ./bin/app cmd/app/app.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /

COPY --from=builder /usr/local/src/pds.env /

COPY --from=builder /usr/local/src/configs /configs

CMD ["/app"]
