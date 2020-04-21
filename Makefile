PLUGIN_NAME=discworld

all: ${PLUGIN_NAME}.so

WEBFILES = \
	web/html/templates.html\
	web/js/controllers.js

generated:
	mkdir generated

generated/data.go: generated ${WEBFILES}
	${GOPATH}/bin/go-bindata -o generated/data.go -pkg generated web/html web/js

${PLUGIN_NAME}.so: generated/data.go
	go build -buildmode=plugin -o ${PLUGIN_NAME}.so .

${PLUGIN_NAME}_debug.so: generated/data.go
	go build -buildmode=plugin -o ${PLUGIN_NAME}_debug.so -gcflags='all=-N -l' .

.PHONY: ${PLUGIN_NAME}.so ${PLUGIN_NAME}_debug.so
