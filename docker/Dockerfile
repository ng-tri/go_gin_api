FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go
# RUN go build -o migrate ./cmd/migrate/main.go
# RUN go build -o worker ./cmd/worker/main.go 

EXPOSE 8080

CMD ["./server"]
