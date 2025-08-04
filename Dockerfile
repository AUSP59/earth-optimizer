
FROM golang:1.22

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o optiterra ./cmd/earthopt

EXPOSE 8080
CMD ["./optiterra"]
