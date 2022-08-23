build:
	go build -o ./out/myapp ./app/cmd/main.go
run:
	go run app/cmd/main.go
start:
	./out/myapp
image:
	docker build -t beamuuuu/redis-getset -f Dockerfile .

