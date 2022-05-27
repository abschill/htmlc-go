all: clr run-0

clr:
	clear

run-0:
	go run ./v2

run-1:
	go run ./v2 ./.spec
