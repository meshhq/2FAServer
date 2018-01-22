# This how we want to name the binary output
BINARY=2FAServer

# These are the values we want to pass for VERSION and BUILD
VERSION=1.0.0
BUILD=`git rev-parse HEAD`

.DEFAULT_GOAL: $(BINARY)

# Builds the project
$(BINARY):
	dep ensure
	go-bindata public/...
	go build ${LDFLAGS} -o ${BINARY} .

install:
	go install ${LDFLAGS} -o ${BINARY} .

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install