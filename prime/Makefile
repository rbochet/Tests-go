all: run

prime:
	8g -o  prime.8 prime.go
build: prime
	8l -o prime prime.8
run: build
	./prime

.PHONY : clean
clean :
		-rm prime *.8
		ls
