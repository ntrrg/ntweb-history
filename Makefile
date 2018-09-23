include config.mk

.PHONY: all
all: build

.PHONY: build
build:
	$(MAKE) -s "hugo ''"

.PHONY: build-docker
build-docker:
	docker build -t $(docker_image) .

.PHONY: clean
clean:
	rm -rf public resources
	docker rm -f $(lint_container) > /dev/null 2> /dev/null || true
	docker rm -f $(lint_container)-watch > /dev/null 2> /dev/null || true

.PHONY: hugo%
hugo%:
	@docker run --rm -it \
		-e PORT=$(hugo_port) \
		-p $(hugo_port):$(hugo_port) \
		-u $$(id -u $$USER) \
		-v "$$PWD":/site/ \
		ntrrg/hugo:$(hugo_version) $*

.PHONY: lint
lint:
	@docker run --name $(lint_container) -it \
		-v "$$PWD":/files/ \
		ntrrg/md-linter 2> /dev/null || \
	docker start -ai $(lint_container)

.PHONY: lint-watch
lint-watch:
	@docker run --name $(lint_container)-watch -it \
		-v "$$PWD":/files/ \
		ntrrg/md-linter:watch 2> /dev/null || \
	docker start -ai $(lint_container)-watch

.PHONY: run
run:
	$(MAKE) -s "hugo "

