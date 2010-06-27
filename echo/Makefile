all: run

echo:
	8g -o  echo.8 echo.go
build: echo
	8l -o echo echo.8
run: build
	./echo 	ceci est un test

.PHONY : clean
clean :
		-rm echo *.8
		ls
