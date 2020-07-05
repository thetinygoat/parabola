build:
	go build -o bin/memdb ./src/
	go build -o bin/client ./utils/client/main.go

run:
	./bin/memdb