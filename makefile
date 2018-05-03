GO_META_LINTER := $(GOPATH)/bin/gometalinter.v2

$(GO_META_LINTER):
	go get -u gopkg.in/alecthomas/gometalinter.v2	 
	gometalinter.v2 --install &> /dev/null

build: install-dep
	# make lint	 
	make test
	go build

just-build: install-dep
	# make lint	 
	go build	

# Downloads the 'dep' package and uses it to 
# install all the dependencies for the project
install-dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure 

# Runs the linter over the project and runs
# the automated tests for the project
test: just-test

# Runs the automated tests for the project
just-test:
	go list ./... | grep -v /vendor | xargs go test -p 1 -timeout=20s

# Runs the linter over the project and
# reports any linting errors
lint: $(GO_META_LINTER)
	test -z $(gofmt -s -l $GO_FILES)
	gometalinter.v2 ./... --deadline=60s --vendor
