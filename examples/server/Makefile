#
# Run a server to access the examples
#
PORT := 8200
WEB := /tmp/wasm
GOROOT := $(shell go env GOROOT)
TINYGOROOT := $(shell tinygo env TINYGOROOT)

all: dir hello
	go run server.go --port=${PORT} --base=${WEB}

hello:
	mkdir -p ${WEB}/hello
	cp ../hello/index.html ${WEB}/hello
	make hello-go # or hello-tinygo
	echo Hello world at http://localhost:${PORT}/hello

hello-tinygo:
	(cd ../hello; tinygo build -target wasm -o ${WEB}/hello.wasm)
	cp -f ${TINYGOROOT}/targets/wasm_exec.js ${WEB}

hello-go:
	(cd ../hello; GOOS=js GOARCH=wasm go build -o ${WEB}/hello.wasm)
	cp ${GOROOT}/misc/wasm/wasm_exec.js ${WEB}

dir:
	mkdir -p ${WEB}

clean:
	rm -rf ${WEB}
