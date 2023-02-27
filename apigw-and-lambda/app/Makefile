build:
	if [ -d bin ]; then rm -rf bin/main.zip; fi;
	GOARCH=amd64 GOOS=linux go build -o bin/main main.go
	zip bin/main.zip bin/main
	rm bin/main
