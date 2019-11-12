hugo_version := 0.59.1
hugo_port := 1313

.PHONY: all
all: build

.PHONY: build
build:
	hugo

.PHONY: bump-version-hugo
bump-version-hugo:
	@grep -lR "$(hugo_version)" . | \
		grep -v "^\./\.git/" | \
		grep -v "\.swp\$$" | \
		grep -v "^\./go\.sum" | \
		grep -v "^\./content/"

.PHONY: clean
clean:
	rm -rf public resources

.PHONY: run
run:
	hugo server -DEF --noHTTPCache --baseUrl / \
		--bind 0.0.0.0 --port $(hugo_port) --appendPort=false

# Docker

.PHONY: docker-build
docker-build:
	@docker run --rm -it \
		-u $$(id -u $$USER) \
		-v "$${TMPDIR:-/tmp}":/tmp/ \
		-v "$$PWD":/site/ \
		ntrrg/hugo:$(hugo_version)

.PHONY: docker-run
docker-run:
	@docker run --rm -it \
		-e PORT=$(hugo_port) \
		-p $(hugo_port):$(hugo_port) \
		-u $$(id -u $$USER) \
		-v "$${TMPDIR:-/tmp}":/tmp/ \
		-v "$$PWD":/site/ \
		ntrrg/hugo:$(hugo_version) server -DEF --noHTTPCache --baseUrl / \
		--bind 0.0.0.0 --port $(hugo_port) --appendPort=false

