
REMOVESYMBOL := -w -s
ifeq (true, $(DEBUG))
        REMOVESYMBOL =
        GCFLAGS=-gcflags=all="-N -l "
endif
LDFLAGS += $(REMOVESYMBOL)

default: all

clean: 

vet:
	go vet ./...

all: clean  build

debug:
	@echo $(GCFLAGS)
	@echo $(LDFLAGS)

build:
	go build -tags="$(BUILD_TAGS)" -ldflags "$(LDFLAGS)" $(GCFLAGS) ./cmd/...

.PHONY: all build clean debug vet 
