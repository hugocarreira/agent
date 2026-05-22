GO        = $(HOME)/.local/go/bin/go
GOROOT    = $(HOME)/.local/go
NAME      = agentrc
PLATFORMS = linux-amd64 linux-arm64 darwin-amd64 darwin-arm64

.PHONY: build install cross-build clean

build:
	GOROOT=$(GOROOT) $(GO) build -o bin/$(NAME) ./cmd/$(NAME)

install:
	GOROOT=$(GOROOT) $(GO) install ./cmd/$(NAME)

cross-build: $(PLATFORMS)
	@echo "✅ All builds complete"

$(PLATFORMS):
	@mkdir -p bin
	@os=$$(echo $@ | cut -d- -f1); \
	arch=$$(echo $@ | cut -d- -f2); \
	echo "  Building $(NAME)-$$os-$$arch"; \
	GOROOT=$(GOROOT) GOOS=$$os GOARCH=$$arch $(GO) build \
		-o bin/$(NAME)-$$os-$$arch ./cmd/$(NAME)

clean:
	rm -rf bin/
