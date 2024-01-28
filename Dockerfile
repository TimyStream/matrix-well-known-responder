FROM golang:1.21-alpine
WORKDIR /usr/src/app
COPY . .
CMD ["go", "run", "app.go"]
