
helloWorld:
	8g -o  helloWorld.8 helloWorld.go
build: helloWorld
	8l -o helloWorld helloWorld.8
run: build
	./helloWorld

.PHONY : clean
clean :
		-rm helloWorld *.8
		ls
