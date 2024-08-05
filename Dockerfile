FROM golang:1.22-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main 

EXPOSE 7777
 
CMD ["./main"]