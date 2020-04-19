all: discworld.so

WEBFILES = \
	web/html/templates.html\
	web/js/controllers.js

generated:
	mkdir generated

generated/data.go: generated ${WEBFILES}
	${GOPATH}/bin/go-bindata -o generated/data.go -pkg generated web/html web/js

discworld.so: generated/data.go
	go build -buildmode=plugin .

.PHONY: discworld.so
