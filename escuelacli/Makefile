current_path = $(shell pwd)
current_dir = $(shell basename "$(PWD)")
gopath = $(GOPATH)

test:
	go test ./pkg/db ./pkg/table

build:
	go build -o cli main.go

clean:
	rm cli

coverage:
	go test -coverprofile=coverage.out ./pkg/db ./pkg/table
	go tool cover -html=coverage.out

docs:
	rm -rf $(gopath)/src/godoc/src
	mkdir -p $(gopath)/src/godoc/src
	rsync -a $(current_path) $(gopath)/src/godoc/src
	(cd $(gopath)/src/godoc; godoc -goroot .)
	cd $(current_path)
