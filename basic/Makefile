# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
DIR_NAME := $(shell pwd)
PACKAGE_NAME := $(shell basename $(DIR_NAME))
TARGET_DIR_BASE := $(DIR_NAME)/target
TARGET_DIR := $(TARGET_DIR_BASE)

# define functions - These can be reused in multiple targets. Any exported VARs are available in the function
define compile 
	for COMMAND in $(DIR_NAME)/cmd/* ; do \
		if [[ -z "$${GOOS}" ]]; then \
			GOOS=$${OSTYPE};\
		fi; \
		if [[ -d "$${COMMAND}" ]]; then \
			go build -trimpath -v -o "$(TARGET_DIR)/bin/$$(basename $${COMMAND})-$${GOOS}" "$${COMMAND}/..."; \
		fi \
	done
endef

define copy-support
	mkdir -p $(TARGET_DIR)
endef

# Full build commands
all: clean-target mods lint test build
release: clean mods lint test package-ansible

# Code cleanup commands
mods:
	go mod tidy
	go mod verify
format: $(SOURCE_LIST)
	find . -name \*.go -not -path vendor -not -path target -exec $(GOPATH)/bin/goimports -w {} \;
lint: format
	${GOPATH}/bin/golangci-lint run

# Build commands
build: 
	$(call compile)
test: 
	$(GOTEST) -v ./...
clean:
	go clean -modcache -testcache -cache
	rm -rf $(TARGET_DIR_BASE)
clean-target:
	rm -rf $(TARGET_DIR_BASE)

# Cross compilation
build-all: build-osx build-linux build-freebsd

build-osx: export CGO_ENABLED=0
build-osx: export GOOS=darwin
build-osx: export GOARCH=amd64
build-osx: export TARGET_DIR = $(TARGET_DIR_BASE)/$(GOOS)
build-osx:
	$(call copy-support)
	$(call compile)

build-linux: export CGO_ENABLED=0
build-linux: export GOOS=linux
build-linux: export GOARCH=amd64
build-linux: export TARGET_DIR = $(TARGET_DIR_BASE)/$(GOOS)
build-linux:
	$(call copy-support)
	$(call compile)

build-freebsd: export CGO_ENABLED=0
build-freebsd: export GOOS=freebsd
build-freebsd: export GOARCH=amd64
build-freebsd: export TARGET_DIR = $(TARGET_DIR_BASE)/$(GOOS)
build-freebsd:
	$(call copy-support)
	$(call compile)

# Packaging
package-binaries: package-osx package-linux package-freebsd
package-osx: build-osx
	tar -zcvf $(TARGET_DIR_BASE)/$(PACKAGE_NAME)-$(GOOS).tar.gz $(TARGET_DIR_BASE)/$(GOOS)/*
package-linux: build-linux
	tar -zcvf $(TARGET_DIR_BASE)/$(PACKAGE_NAME)-$(GOOS).tar.gz $(TARGET_DIR_BASE)/$(GOOS)/*
package-freebsd: build-freebsd
	tar -zcvf $(TARGET_DIR_BASE)/$(PACKAGE_NAME)-$(GOOS).tar.gz $(TARGET_DIR_BASE)/$(GOOS)/*
package-ansible: clean-target build-all
	cp -r deployments/ansible target/ansible
	cp -r $(TARGET_DIR_BASE)/**/bin/* $(TARGET_DIR_BASE)/ansible/files
	cd target/ansible; tar -zcvf ../$(PACKAGE_NAME)_ansible.tar.gz .
