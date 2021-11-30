# build & test automation

build:
	go build main.go

test: build
	@echo Test 1 - 3 enemies - Nice
	./main 3

clean:
	rm -rf main