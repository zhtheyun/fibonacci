# Fibonacci API
#
#
# Targets:
# 	build: Builds the code
# 	fmt: Formats the source files
# 	clean: cleans the code
#       lint: lint the code
#	test: Runs the tests
#       vendor_get: Install the 3rd party jar file
#
#
#
# Go parameters

BUILD_DATE ?= $(shell date +%Y-%m-%d\T%H:%M)
MAJOR_VERSION ?= 1
MINOR_VERSION ?= 0
PATCH_VERSION ?= 0
BUILD_VERSION ?= 0
VERSION ?= $(MAJOR_VERSION).$(MINOR_VERSION).$(PATCH_VERSION).$(BUILD_VERSION)
# GOBUILD_VERSION_ARGS := -ldflags "-X cmd.Version=$(VERSION) -X main.BuildDate='$(BUILD_DATE)'"
GOBUILD_VERSION_ARGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE) "

GOCMD=cd $(MKFILE_DIR) && go
GODOCCMD=cd $(MKFILE_DIR) && godoc
GOBUILD=$(GOCMD) build $(GOBUILD_VERSION_ARGS)
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOLIST=$(GOCMD) list
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))

# Set the GOPATH to ensure we worked in an isolated environment
GOPATH:=$(MKFILE_DIR)../../../..
export GOPATH
GOLINT:=$(GOPATH)/bin/golint
GOCOV:=$(GOPATH)/bin/gocov
GODEP:=$(GOPATH)/bin/dep
GOCOVXML:=$(GOPATH)/bin/gocov-xml

SERVICE_NAME := $(strip $(notdir $(patsubst %/,%,$(MKFILE_DIR))))

#TODO: We need add some checking for go env.
check_env:
	echo "It make me mad due to network issue. run dep status yourself!"
	echo "System environment validation pass"


# List of pkgs for the project
PKGS:=$(shell export GOPATH=$(GOPATH) && $(GOLIST) ./... | grep -v vendor)
#PKGS:=$(shell export GOPATH=$(GOPATH) && $(GOLIST))
ALL_LIST:=$(PKGS)



BUILD_LIST = $(foreach int, $(ALL_LIST), $(int)_build)
CLEAN_LIST = $(foreach int, $(ALL_LIST), $(int)_clean)
INSTALL_LIST = $(foreach int, $(ALL_LIST), $(int)_install)
TEST_LIST = $(foreach int, $(ALL_LIST), $(int)_test)
FMT_TEST = $(foreach int, $(ALL_LIST), $(int)_fmt)
LINT_LIST = $(foreach int, $(ALL_LIST), $(int)_lint)
VET_LIST = $(foreach int, $(ALL_LIST), $(int)_vet)
DOC_LIST = $(foreach int, $(ALL_LIST), $(int)_doc)


DIST_DIR := $(GOPATH)/dist
BIN_DIR := $(GOPATH)/bin
COVERAGE_DIR := $(GOPATH)/coverage

PACKAGE_DIR := $(GOPATH)/pkg


# All are .PHONY for now because dependencyness is hard
.PHONY: $(VET_LIST) $(DOC_LIST) $(LINT_LIST) $(CLEAN_LIST) $(TEST_LIST) $(INSTALL_LIST) $(BUILD_LIST) build doc fmt lint test clean vet dist check_env cover FORCE

.DEFAULT_GOAL := all


all: build test cover dist
build: $(GODEP) $(BUILD_LIST)
clean: $(CLEAN_LIST)
	-rm -rf $(DIST_DIR)
	-rm -rf $(BIN_DIR)
	-rm -rf $(COVERAGE_DIR)
	-rm -rf $(PACKAGE_DIR)
	-rm -rf $(MKFILE_DIR)/$(SERVICE_NAME)

install: $(GODEP) $(INSTALL_LIST)
test: $(GODEP) $(TEST_LIST)
fmt: $(FMT_TEST)
lint: $(LINT_LIST)
doc: $(DOC_LIST)
vet: $(GODEP) $(VET_LIST)
dist: build
cover: $(GODEP) $(COVERAGE_DIR)/cobertura-coverage.xml

#We introduce several environment variable for our test environment, we need to check it before we run test to avoid the confusion


$(BIN_DIR)/$(SERVICE_NAME): install

$(LINT_LIST): %_lint: $(GOLINT)
	$(GOLINT) $*
$(BUILD_LIST): %_build:
	$(GOBUILD) $*
$(CLEAN_LIST): %_clean:
	$(GOCLEAN) $*
$(INSTALL_LIST): %_install: %_build
	$(GOINSTALL) $*
$(TEST_LIST): %_test: %_lint %_fmt %_vet
	$(GOTEST) $*
$(FMT_TEST): %_fmt:
	$(GOFMT) $*
$(VET_LIST): %_vet:
	$(GOVET) $*
$(DOC_LIST): %_doc:
	$(GODOCCMD) $*


# Coverage output: coverage/$PKG/coverage.out
COVPKGS:=$(addsuffix /coverage.out,$(addprefix $(COVERAGE_DIR)/,$(PKGS)))

$(COVERAGE_DIR)/all.out: $(COVPKGS)
	mkdir -p $(dir $@)
	echo "mode: set" >$@
	grep -hv "mode: set" $(wildcard $^) >>$@

$(COVPKGS): FORCE $(GODEP)
	mkdir -p $(dir $@)
	$(GOTEST) -coverprofile $@ $(patsubst $(COVERAGE_DIR)/%/coverage.out,%,$@)


$(COVERAGE_DIR)/cobertura-coverage.xml: $(COVERAGE_DIR)/all.out $(GOCOV) $(GOCOVXML)
	$(GOCOV) convert $< | $(GOCOVXML) >$@

$(GOCOV):
	go get -v github.com/axw/gocov/...
$(GOCOVXML):
	go get -v github.com/AlekSi/gocov-xml
$(GOLINT):
	go get -v github.com/golang/lint/golint
$(GODEP):
	mkdir -p $(BIN_DIR) || true
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
FORCE:
