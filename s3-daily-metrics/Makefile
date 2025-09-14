build:
	go build -o main .

function:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bootstrap .
	chmod +x bootstrap
	zip -r function.zip bootstrap
	rm bootstrap

clean:
	rm -f main bootstrap function.zip

test:
	go test ./...

.PHONY: build function clean test