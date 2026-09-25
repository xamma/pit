BINARY  := pit
CMD     := ./cmd/pit
DIST    := dist

# Overridden by the release workflow, which passes the git tag as VERSION.
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

# Release targets.
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

export CGO_ENABLED := 0

.PHONY: build run fmt vet install dist clean

build: ## build the binary for the host platform into ./$(BINARY)
	go build -trimpath -ldflags '-s -w' -o $(BINARY) $(CMD)

run: build ## build and run it
	./$(BINARY)

fmt:
	gofmt -l -w .

vet:
	go vet ./...

install: ## install into $GOBIN
	go install -trimpath -ldflags '-s -w' $(CMD)

dist: ## cross-compile every release artifact into ./$(DIST)
	@rm -rf $(DIST) && mkdir -p $(DIST)
	@for platform in $(PLATFORMS); do \
		os=$${platform%/*}; arch=$${platform#*/}; \
		ext=''; [ "$$os" = windows ] && ext='.exe'; \
		name="$(BINARY)_$(VERSION)_$${os}_$${arch}"; \
		echo "  building $$name"; \
		mkdir -p "$(DIST)/$$name"; \
		GOOS=$$os GOARCH=$$arch go build -trimpath -ldflags '-s -w' \
			-o "$(DIST)/$$name/$(BINARY)$$ext" $(CMD) || exit 1; \
		cp Readme.md LICENSE.md "$(DIST)/$$name/" 2>/dev/null || true; \
		if [ "$$os" = windows ]; then \
			(cd $(DIST) && zip -qr "$$name.zip" "$$name"); \
		else \
			tar -czf "$(DIST)/$$name.tar.gz" -C $(DIST) "$$name"; \
		fi; \
		rm -rf "$(DIST)/$$name"; \
	done
	@cd $(DIST) && sha256sum * > checksums.txt && echo && cat checksums.txt

clean:
	rm -rf $(DIST) $(BINARY)
