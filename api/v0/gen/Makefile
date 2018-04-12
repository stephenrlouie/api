.PHONY: all

OPTIKON = $(GOPATH)/src/wwwin-github.cisco.com/edge/optikon-api
VERSION = v0

all:
	rm -rf ../models
	rm -rf ../server/restapi/operations
	mkdir -p ../client;
	docker run -v $(OPTIKON):/go/src/wwwin-github.cisco.com/edge/optikon-api \
		 	-w /go/src/wwwin-github.cisco.com/edge/optikon-api/api/$(VERSION)/gen \
			quay.io/goswagger/swagger:0.13.0 generate client swagger.yml \
	    -t /go/src/wwwin-github.cisco.com/edge/optikon-api/api/$(VERSION);
	mkdir -p ../server;
	docker run -v $(OPTIKON):/go/src/wwwin-github.cisco.com/edge/optikon-api \
			-w /go/src/wwwin-github.cisco.com/edge/optikon-api/api/$(VERSION)/gen \
	    quay.io/goswagger/swagger:0.13.0 generate server swagger.yml \
	   	-t /go/src/wwwin-github.cisco.com/edge/optikon-api/api/$(VERSION)/server \
			--existing-models wwwin-github.cisco.com/edge/optikon-api/api/$(VERSION)/models \
			--exclude-main \
			--flag-strategy=pflag;
	rm -rf ../server/cmd;
