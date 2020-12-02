build:
        mkdir bin
	go build -o bin/tbot src/main.go
run:
	go run src/main.go
#compile:
#	echo "Compiling for every OS and Platform"
#	GOOS=windows GOARCH=amd64 go build -o bin/tbot-win-amd64.exe src/main.go
#	GOOS=linux GOARCH=386 go build -o bin/tbot-linux-x86 src/main.go
#	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
