.PHONY: deps

UNAME := $(shell echo $(shell uname) | tr '[:upper:]' '[:lower:]')
KAKASI_VERSION=2.3.6
KAKASI_DIR=kakasi-$(KAKASI_VERSION)
KAKASI_ARCHIVE=kakasi-$(KAKASI_VERSION).tar.gz
BUILD_DIR=$(PWD)/.build
KAKASI_INC=$(BUILD_DIR)/kakasi/include
KAKASI_LIBDIR=$(BUILD_DIR)/kakasi/lib
DEPS_DIR=$(PWD)/deps

ifeq ($(UNAME), darwin)
	include tools/darwin.mk
endif
ifeq ($(UNAME), linux)
	include tools/linux.mk
endif

all: deps $(UNAME) static clean

deps:
	if [ -d "$(BUILD_DIR)" ]; then\
		rm -rf $(BUILD_DIR);\
	fi
	mkdir $(BUILD_DIR)

	cd $(BUILD_DIR) && \
		curl -o $(KAKASI_ARCHIVE) http://kakasi.namazu.org/stable/$(KAKASI_ARCHIVE) && \
		tar xvfz $(KAKASI_ARCHIVE) && \
		patch -u $(KAKASI_DIR)/src/kanjiio.c < ../tools/kanjiio.diff && \
		patch -u $(KAKASI_DIR)/configure < ../tools/configure.diff

	if [ -f "$(BUILD_DIR)/$(KAKASI_DIR)/Makefile" ]; then\
		make clean;\
	fi\

	cd $(BUILD_DIR)/$(KAKASI_DIR) && \
		./configure --prefix=$(BUILD_DIR)/kakasi && \
		make && \
		make install

static:
	rm -rf dict
	statik -src=deps/share -p=dict -dest=. --include="*"

test:
	go test

lint:
	golangci-lint run

clean:
	rm -rf $(BUILD_DIR)
