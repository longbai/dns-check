all: get_api
	bash -c "export GOPATH=`pwd` && go install -v ./..."
	cp src/github.com/slene/iploc/iploc.dat bin/

get_api:
	bash -c "export GOPATH=`pwd` && go get github.com/miekg/dns && go get github.com/slene/iploc"
