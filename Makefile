SHELL := /bin/bash
TARGETS = stardust

# http://docs.travis-ci.com/user/languages/go/#Default-Test-Script
test:
	go get -d && go test -v

bench:
	go test -bench=.

imports:
	goimports -w .

fmt:
	go fmt ./...

vet:
	go vet ./...

all: fmt test
	go build

install:
	go install

clean:
	go clean
	rm -f coverage.out
	rm -f $(TARGETS)
	rm -f stardust-*.x86_64.rpm
	rm -f debian/stardust*.deb
	rm -rf debian/stardust/usr

cover:
	go get -d && go test -v	-coverprofile=coverage.out
	go tool cover -html=coverage.out

stardust:
	go build cmd/stardust/stardust.go

# ==== packaging

deb: $(TARGETS)
	mkdir -p debian/stardust/usr/sbin
	cp $(TARGETS) debian/stardust/usr/sbin
	cd debian && fakeroot dpkg-deb --build stardust .

REPOPATH = /usr/share/nginx/html/repo/CentOS/6/x86_64

publish: rpm
	cp stardust-*.rpm $(REPOPATH)
	createrepo $(REPOPATH)

rpm: $(TARGETS)
	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
	cp ./packaging/stardust.spec $(HOME)/rpmbuild/SPECS
	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
	./packaging/buildrpm.sh stardust
	cp $(HOME)/rpmbuild/RPMS/x86_64/stardust*.rpm .
