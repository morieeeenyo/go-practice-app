FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/volatiletech/sqlboiler@latest
RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
RUN go install github.com/cosmtrek/air@latest

RUN go mod download && go mod verify

CMD ["air", "-c", ".air.toml"]