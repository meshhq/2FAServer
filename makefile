# This how we want to name the binary output
BINARY=2FAServer

.DEFAULT_GOAL: $(BINARY)

# Builds the project
$(BINARY):
	dep ensure
	go-bindata public/...	
	go build -o ${BINARY} .

install:
	dep ensure
	go-bindata public/...	
	go install -o ${BINARY} .

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install