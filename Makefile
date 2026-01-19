NAME=nekoapi-ruijietoolkit-executor
BINDIR=bin
VERSION=0.0.1

TYPE:=release

BUILDTIME=$(shell date -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags \
		'-s -w -extldflags "-static" -buildid= \
		-X "shelltool/shelltool/constant.Version=$(VERSION)-$(TYPE)" \
		-X "shelltool/shelltool/constant.BuildTime=$(BUILDTIME)" \
		-X "shelltool/shelltool/constant.AppType=$(TYPE)"'

UPX_FLAGS := --best --lzma

PLATFORM_LIST = \
	linux-amd64 \
	linux-arm64 \
	linux-armv5 \
	linux-armv6 \
	linux-armv7 \
	linux-mips \
	linux-mipsle \
	linux-mips64 \
	linux-mips64le

all: normal_build

normal_build: $(PLATFORM_LIST)

linux-amd64:
	GOARCH=amd64 GOOS=linux GOAMD64=v3 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-arm64:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv5:
	GOARCH=arm GOOS=linux GOARM=5 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv6:
	GOARCH=arm GOOS=linux GOARM=6 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv7:
	GOARCH=arm GOOS=linux GOARM=7 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mips:
	GOARCH=mips GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mipsle:
	GOARCH=mipsle GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mips64:
	GOARCH=mips64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mips64le:
	GOARCH=mips64le GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

upx: normal_build
	mkdir -p $(BINDIR)/upx
	-rm $(BINDIR)/upx/*

	@for arch in $(PLATFORM_LIST); do \
		file="$(BINDIR)/$(NAME)-$$arch"; \
		if [ -f "$$file" ]; then \
			echo "UPX: $$file"; \
			upx $(UPX_FLAGS) "$$file" -o $(BINDIR)/upx/$(NAME)-$$arch-uxp >/dev/null 2>&1 || \
				echo "  -> skip (unsupported or error)"; \
		fi; \
	done