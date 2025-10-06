
tst:
	@make test | grep -E '(FAIL|    --- PASS)' --color=always

test:
	go clean -testcache
	go test -v ./...

run:
	go run ./cmd/globe.go

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build \
		-o ./build/globe \
		-gcflags -m=2 \
		./cmd/ 

install:
	ln "$(realpath ./build/globe)" -s ~/.local/bin/globe

hub_update:
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./build/globe)"

