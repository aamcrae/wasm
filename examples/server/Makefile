#
# Run a server to access the examples
#
PORT := 8200
WEB := /tmp/wasm
GOROOT := $(shell go env GOROOT)
TINYGOROOT := $(shell tinygo env TINYGOROOT)

all: dir hello compose
	go run server.go --port=${PORT} --base=${WEB}

hello:
	mkdir -p ${WEB}/hello
	cp ../hello/index.html ${WEB}/hello
	make hello-go # hello-tinygo
	echo Hello world at http://localhost:${PORT}/hello

compose:
	mkdir -p ${WEB}/compose
	cp ../compose/index.html ${WEB}/compose
	make compose-go
	echo Compose examples at http://localhost:${PORT}/compose

hello-tinygo: tiny-exec
	(cd ../hello; tinygo build -target wasm -o ${WEB}/hello.wasm)

hello-go: go-exec
	(cd ../hello; GOOS=js GOARCH=wasm go build -o ${WEB}/hello.wasm)

compose-go:
	(cd ../compose; GOOS=js GOARCH=wasm go build -o ${WEB}/compose.wasm)

tiny-exec:
	cp -f ${TINYGOROOT}/targets/wasm_exec.js ${WEB}

go-exec:
	cp ${GOROOT}/misc/wasm/wasm_exec.js ${WEB}

dir:
	mkdir -p ${WEB}

clean:
	rm -rf ${WEB}
