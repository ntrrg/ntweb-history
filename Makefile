hugo_version := 0.52
hugo_port := 1313
docker_image := ntrrg/site
lint_container := $(subst /,-,$(docker_image))-lint

.PHONY: all
all: build

.PHONY: build
build:
	@docker run --rm -it \
		-u $$(id -u $$USER) \
		-v "$$PWD":/site/ \
		ntrrg/hugo:$(hugo_version)

.PHONY: build-docker
build-docker:
	docker build -t $(docker_image) .

.PHONY: clean
clean:
	rm -rf public resources
	docker rm -f $(lint_container) > /dev/null 2> /dev/null || true
	docker rm -f $(lint_container)-watch > /dev/null 2> /dev/null || true

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
	@docker run --rm -it \
		-p $(hugo_port):$(hugo_port) \
		-u $$(id -u $$USER) \
		-v "$$PWD":/site/ \
		ntrrg/hugo:$(hugo_version) server \
			-DEF --bind=0.0.0.0 --port=$(hugo_port) --baseUrl=/ --appendPort=false

