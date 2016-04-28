all:
	go run main.go api.go

dev:
	gin -a 8080 run
