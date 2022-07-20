build:
	go build -o bin/weproov cmd/weproov/weproov.go

run:
	./bin/weproov

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/weproov/weproov.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 cmd/weproov/weproov.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 cmd/weproov/weproov.go

all : build run