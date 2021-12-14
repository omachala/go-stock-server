build: 
	go build -o server main.go

run: build
	./server

compress:
	go build -ldflags="-s -w" -o server main.go
	upx --brute server

watch:
	ulimit -n 1000
	reflex -s -r '\.go$$' make run
