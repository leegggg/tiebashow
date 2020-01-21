GOOS=windows GOARCH=386 \
  CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc \
  go build -o ./target/