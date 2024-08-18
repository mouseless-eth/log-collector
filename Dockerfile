FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE ${PORT}
CMD ["./main"]
