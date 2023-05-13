compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o bin/jrpc-linux-amd64 main.go
	GOOS=linux GOARCH=arm go build -o bin/jrpc-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/jrpc-linux-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/jrpc-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/jrpc-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/jrpc-darwin-arm64 main.go