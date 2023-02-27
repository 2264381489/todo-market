gen:
	goctl api go -api todoMarket.api -dir .

doc:
	rm -f ./todoMarket.md
	go run ../../goctl-doc/main.go -api todoMarket.api

format:
	goctl api format --dir ./todoMarket.api -declare

all:
	make gen && make doc && make format