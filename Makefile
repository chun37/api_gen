.PHONY: dev
dev: export LOCAL_ENV = 1
dev:
	go run cmd/*.go

gen:
	go generate ./...

.PHONY: init
init: bootstrap installdeps
	direnv allow
	cp ./scripts/pre-commit .git/hooks/

.PHONY: bootstrap
bootstrap:
	(cd $(mktemp -d); GO111MODULE=on \
			go get \
			github.com/golang/mock/gomock \
			github.com/golang/mock/mockgen \
			github.com/golang/protobuf/protoc-gen-go \
			github.com/favadi/protoc-go-inject-tag \
			github.com/moznion/go-errgen/cmd/errgen \
	)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0

.PHONY: installdeps
installdeps:
	go mod download
	mkdir -p ./bin
	go build -o ./bin/ ./tools/...

.PHONY: test
test:
	go test ./... -v

.PHONY: server_generator
server_generator:
	go build -o ./bin/server_generator ./server_generator

.PHONY: client_generator
client_generator:
	go build -o ./bin/client_generator ./client_generator

sample_server_generator: server_generator
	cd server_generator/sample && ../../bin/server_generator .

sample_client_generator: client_generator
	cd client_generator/sample && ../../bin/client_generator ../../server_generator/sample

sample: sample_server_generator sample_client_generator

sample_server_generator_run: sample_server_generator
	go run server_generator/sample/cmd/server/main.go

build-release:
	$(shell bash ./scripts/build_release.sh)

.PHONY: cmd
cmd:
	go build -o ./bin/api_gen ./cmd/*.go

gen_samples: cmd
	bash ./samples/generate.sh
