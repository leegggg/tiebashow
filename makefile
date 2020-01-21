default:
	go build -x 

cross-win64:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -x -o ./target/tiebashow.exe

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CXX=g++ CC=gcc go build -x -o ./target/tiebashow

win64:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=clang++ CC=clang go build -x -o ./target/tiebashow.exe

.PHONY: clean

clean:
	rm target/*