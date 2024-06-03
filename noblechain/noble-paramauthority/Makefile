DOCKER := $(shell which docker)

# containerProtoVer=v0.2
# containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
# containerProtoGen=cosmos-sdk-proto-gen-$(containerProtoVer)
# containerProtoGenSwagger=cosmos-sdk-proto-gen-swagger-$(containerProtoVer)
# containerProtoFmt=cosmos-sdk-proto-fmt-$(containerProtoVer)


# proto-all: proto-format proto-lint proto-gen

# proto-gen:
# 	@echo "Generating Protobuf files"
# 	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
# 		sh ./scripts/protocgen.sh; fi
# 	@go mod tidy


BUF_VERSION=1.28.1


protoVer=0.13.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(shell pwd):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-lint proto-gen

proto-format:
	@echo "🤖 Running protobuf formatter..."
	@docker run --rm --volume "$(shell pwd)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) format --diff --write
	@echo "✅ Completed protobuf formatting!"

# proto-gen:
# 	@echo "🤖 Generating code from protobuf..."
# 	@docker run --rm --volume "$(shell pwd)":/workspace --workdir /workspace \
# 		noble-param-proto sh ./proto/protocgen.sh
# 	@echo "✅ Completed code generation!"

proto-lint:
	@echo "🤖 Running protobuf linter..."
	@docker run --rm --volume "$(shell pwd)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) lint
	@echo "✅ Completed protobuf linting!"

proto-setup:
	@echo "🤖 Setting up protobuf environment..."
	@docker build --rm --tag noble-param-proto:latest --file proto/Dockerfile .
	@echo "✅ Setup protobuf environment!"

proto-gen:
	@echo "🤖 Generating code from protobuf..."
	@$(protoImage) sh ./proto/protocgen.sh
	@echo "✅ Completed code generation!"