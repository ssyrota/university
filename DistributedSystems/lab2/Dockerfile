FROM golang:1.21-bullseye as dev
WORKDIR /app

COPY go.mod .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN go env -w CGO_ENABLED=0
RUN ./build.dev.sh
RUN curl -fsSL https://raw.githubusercontent.com/tj/n/master/bin/n | bash -s lts
CMD ["/go/bin/app"]