include config.mk

.PHONY: all
all: build

.PHONY: build
build:
	@CMD="-v" $(MAKE) -s hugo

.PHONY: build-docker
build-docker:
	@docker build -t ntrrg/site .

.PHONY: clean
clean:
	docker rm -f site-lint 2> /dev/null || echo
	rm -rf public

.PHONY: hugo
hugo:
	@docker run --rm -it \
		-u $$(id -u $$USER) \
		-v "$$PWD":/site/ \
	ntrrg/hugo:$(hugo_version) $$CMD

.PHONY: lint
lint:
	@docker run --name site-lint -it \
		-v "$$PWD":/files/ \
	ntrrg/md-linter:watch 2> /dev/null || docker start -ai site-lint

.PHONY: run
run:
	@docker run --rm -it \
		-e PORT=$(hugo_port) \
		-p $(hugo_port):$(hugo_port) \
		-v "$$PWD":/site/ \
	ntrrg/hugo:$(hugo_version)

