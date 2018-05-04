.PHONY: all

OPTIKON = $(GOPATH)/src/github.com/optikon/api
VERSION = v0

all:
	rm -rf ../models
	rm -rf ../server/restapi/operations
	mkdir -p ../client;
	docker run -v $(OPTIKON):/go/src/github.com/optikon/api \
		 	-w /go/src/github.com/optikon/api/api/$(VERSION)/gen \
			quay.io/goswagger/swagger:0.13.0 generate client swagger.yml \
	    -t /go/src/github.com/optikon/api/api/$(VERSION);
	mkdir -p ../server;
	docker run -v $(OPTIKON):/go/src/github.com/optikon/api \
			-w /go/src/github.com/optikon/api/api/$(VERSION)/gen \
	    quay.io/goswagger/swagger:0.13.0 generate server swagger.yml \
	   	-t /go/src/github.com/optikon/api/api/$(VERSION)/server \
			--existing-models github.com/optikon/api/api/$(VERSION)/models \
			--exclude-main \
			--flag-strategy=pflag;
	rm -rf ../server/cmd;
