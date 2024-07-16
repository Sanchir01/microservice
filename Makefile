build:
	go build -o ./.bin/main ./cmd/main/main.go

gql:
	go get github.com/99designs/gqlgen@latest && go run github.com/99designs/gqlgen generate

run: build
	./.bin/main