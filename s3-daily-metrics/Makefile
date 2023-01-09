build:
	go build main.go

function:
#!/bin/bash
	GOARCH=amd64 GOOS=linux go build -o bootstrap
	chmod +x bootstrap
	zip -r function.zip bootstrap
	rm bootstrap