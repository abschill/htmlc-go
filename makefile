all: clr run-0

clr:
	clear

run-0:
	go run ./lib

run-1:
	go run ./lib ./.spec
