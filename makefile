all: clr run-0

clr:
	rm htmlc-go && clear

run-0:
	go build && ./htmlc-go
