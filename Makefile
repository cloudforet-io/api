TARGET = all
VERSION = v1
DOCKER_NAME = spaceone/api_builder

.PHONY: help
.PHONY: all python go gateway json
.PHONY: clean

define banner
	@echo "========================================================================"
	@echo " $(1)"
	@echo "========================================================================"
endef

define set_build_env
	git submodule init
	git submodule update
	docker build -t ${DOCKER_NAME} .
endef

define build
	$(call set_build_env)
	$(call banner, "Start the build protobuf")
	docker run -i --rm -v ${PWD}:/opt ${DOCKER_NAME} python3 build.py ${TARGET} -c$(1)
endef

help:
	@echo "Make Options:"
	@echo " all                 - Generate all"
	@echo " python              - Generate python protobuf"
	@echo " go                  - Generate go protobuf"
	@echo " gateway             - Generate gateway protobuf"
	@echo " json                - Generate API artifact json files"
	@echo " clean               - Clean up dist directory"

all:
	$(call banner, "Generate all : ${TARGET}")
	$(call build, "all")

json:
	$(call banner, "Generate API json files")
	$(call build, "json")

python:
	$(call banner, "Generate python protobuf")
	$(call build, "python")

go:
	$(call banner, "Generate go protobuf")
	$(call build, "go")

gateway:
	$(call banner, "Generate gateway protobuf")
	$(call build, "gateway")

clean:
	$(call banner, "Clean up dist directory")
	sudo rm -rf dist
