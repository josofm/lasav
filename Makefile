version     ?= latest
devimg       = lasavdev
GOPATH      ?= $(HOME)/go
packagename  = github.com/josofm/lasav
workdir      = /go/src/$(packagename)
rundev       = docker run --rm -v `pwd`:$(workdir) --workdir $(workdir) $(devimg)
cov          = coverage.out
covhtml      = coverage.html

all: check analyze

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

release: guard-version
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

imagedev:
	docker build -t $(devimg) -f ./hack/Dockerfile.dev .

static-analysis: imagedev
	$(rundev) golangci-lint run ./...

check: imagedev
	$(rundev) go test -timeout 20s -race -coverprofile=$(cov) ./...

coverage: imagedev
	$(rundev) go tool cover -html=$(cov) -o=$(covhtml)

coverage-show: imagedev
	$(rundev) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

