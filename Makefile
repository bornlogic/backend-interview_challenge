name = ibcc
img = golang:latest
src = github.com/iuryfukuda/$(name)
workdir = /go/src/$(src)
run = docker run --net=host -it -v $(PWD):$(workdir) -w $(workdir) --rm $(img)

shell:
	$(run) sh

build:
	$(run) go build -o testMatrix cmd/testMatrix/main.go

run:
	$(run) go run cmd/api/main.go $(args)

test:
ifeq ($(origin args), undefined)
	$(run) go test ./...
else
	$(run) go test $(args)
endif
