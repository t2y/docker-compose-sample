FROM golang:alpine3.19
RUN apk add --no-cache git make net-tools bind-tools bash
WORKDIR /app
COPY . .
ENTRYPOINT [ "go", "run", "main.go" ]
