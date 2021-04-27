SHELL := /bin/bash

# VARIABLES used
export GOBIN = $(shell pwd)/bin
export BINARY := ggd-cli

# Include all makefiles used in the project
include build/Makefile.help
include build/Makefile.dev
include build/Makefile.build
include build/Makefile.deps
