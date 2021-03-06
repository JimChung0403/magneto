GOPATH:=$(CURDIR)/../../../../
export GOPATH

export CGO_ENABLED=0

all: bld

bld: magneto

magneto:
	govendor sync
	go build -o bin/magneto github.com/invite0403/magneto/cmd/magneto

clean:
	@rm -f init/magneto
	@rm -rf status
	@rm -f  log/*log*
	@rm -rf ./output

cleanlog:
	@rm -f log/*log*
