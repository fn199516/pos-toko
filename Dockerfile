FROM golang:1.23

WORKDIR /app/go/pos-toko

COPY . .

RUN go build -o main cmd/pos-toko/main.go

CMD ["./main"]