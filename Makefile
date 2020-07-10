build:
	go build -o bin/dictX ./src/
	go build -o bin/client ./utils/client/main.go

run:
	./bin/dictX

run_client:
	./bin/client