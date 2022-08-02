# Step 1: Modules caching
FROM golang:1.17.1-alpine3.14

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go install -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" github.com/golang-migrate/migrate/v4/cmd/migrate@latest 
RUN go build -o /cmd/scaffold_service /app/cmd/main.go

CMD ["/cmd/scaffold_service"]