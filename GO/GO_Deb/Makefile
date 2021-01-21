GOMODULE             = $(shell grep module go.mod | cut -d" " -f2)
VERSION              = $(shell head -1 debian/changelog | sed 's/.*(\(.*\)).*/\1/')

GO                   = GO111MODULE=on CGO_ENABLED=0 go
CGO                  = GO111MODULE=on CGO_ENABLED=1 go
LDFLAGS             += -s -w
GOFLAGS              = -ldflags "$(LDFLAGS)"
GOFLAGS              = -ldflags "$(LDFLAGS)" -mod=vendor

build: build/$(ARCH)/update

build/$(ARCH)/update:
	mkdir -p `dirname $@`
	$(CGO) build -tags  $(GOFLAGS) -o $@ ./cmd/main.go
	cp $@ $@-without-upx
	upx $@