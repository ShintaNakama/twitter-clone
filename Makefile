.PHONY: codegen
codegen: go_gen pb_gen

.PHONY: go_gen
go_gen: ## Goコードを生成する
	go generate ./...

.PHONY: pb_gen
pb_gen: ## protobuf => Goコード生成
	protoc \
		--go_out=plugins=grpc,paths=source_relative:pb \
		--proto_path=protobuf \
		$(wildcard protobuf/*.proto)

.PHONY: gql_gen
gql_gen: ## GraphQLスキーマ => Goコード生成
	cd ./bff && go run github.com/99designs/gqlgen --config=gqlgen.yml


.PHONY: list
list: ## タスク一覧を表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "} {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
