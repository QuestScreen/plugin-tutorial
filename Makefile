# TODO: rename PluginTemplate.so

all: PluginTemplate.so

WEBFILES = \
  web/html/myplugin.html\
	web/js/myplugin.js

generated:
	mkdir generated

generated/data.go: generated ${WEBFILES}
	${GOPATH}/bin/go-bindata -o generated/data.go -pkg generated web/html web/js

PluginTemplate.so: generated/data.go
	go build -buildmode=plugin .

.PHONY: PluginTemplate.so