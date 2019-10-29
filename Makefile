hugo_version := 0.59.0
hugo_port := 1313

.PHONY: all
all: build

.PHONY: build
build:
	$(MAKE) -s "hugo ''"

.PHONY: clean
clean:
	rm -rf public resources

.PHONY: hugo%
hugo%:
	hugo $*

.PHONY: run
run:
	$(MAKE) -s "hugo server -DEF --noHTTPCache --port $(hugo_port)"

# Docker

.PHONY: docker-build
docker-build:
	$(MAKE) -s "docker-hugo ''"

.PHONY: docker-hugo%
docker-hugo%:
	@docker run --rm -it \
		-e PORT=$(hugo_port) \
		-p $(hugo_port):$(hugo_port) \
		-u $$(id -u $$USER) \
		-v "$${TMPDIR:-/tmp}":/tmp/ \
		-v "$$PWD":/site/ \
		ntrrg/hugo:$(hugo_version) $*

.PHONY: docker-run
docker-run:
	$(MAKE) -s "docker-hugo server -DEF --noHTTPCache --port $(hugo_port)"

