all: run-0 run-1


run-0:
	go run ./v2

run-1:
	go run ./v2 ./.spec
