mainprogram=go-sample-master_web
build_commit=$(shell git rev-parse HEAD)
build_version=$(shell git describe --tags 2> /dev/null || echo "dev-$(shell git rev-parse HEAD)")

COMPOSE := docker-compose -f docker-compose.yml

# Commands for running docker compose
RUN_COMPOSE = $(COMPOSE) up --build

.PHONY: version
version:
	echo $(build_version)
	echo $(build_commit)
	
.PHONY: build
build:
	@go build -v

.PHONY: run
run:
	@$(RUN_COMPOSE)

.PHONY: test
test:
# install richgo if not exists
ifeq (, $(shell which richgo))
	$(info "Install rich go")
	@go get -u github.com/kyoh86/richgo
endif
# run test
	richgo test -v ./...
