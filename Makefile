native:
	go build

linux:
	GOOS=linux GOARCH=amd64 go build -o printtables

windows:
	GOOS=windows GOARCH=amd64 go build -o printtables.exe

clean:
	rm -rf printtables printtables.exe

get-deps:
	go get -u github.com/go-sql-driver/mysql
	go get -u golang.org/x/crypto
	go get -u golang.org/x/sys
	go get -u golang.org/x/term
