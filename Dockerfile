FROM golang:alpine
WORKDIR /usr/src/app
COPY . .
EXPOSE 8080
# Run application
CMD ["go", "run", "main.go"]