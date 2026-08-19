.PHONY: generate generate-user generate-product generate-order generate-billing

generate: generate-user generate-product generate-order generate-billing

generate-user:
	protoc \
		--go_out=user_service \
		--go_opt=paths=source_relative \
		--go-grpc_out=user_service \
		--go-grpc_opt=paths=source_relative \
		proto_contracts/user.proto

generate-product:
	protoc \
		--go_out=product_service \
		--go_opt=paths=source_relative \
		--go-grpc_out=product_service \
		--go-grpc_opt=paths=source_relative \
		proto_contracts/product.proto

generate-order:
	protoc \
		--go_out=order_service \
		--go_opt=paths=source_relative \
		--go-grpc_out=order_service \
		--go-grpc_opt=paths=source_relative \
		proto_contracts/order.proto

generate-billing:
	protoc \
		--go_out=billing_service \
		--go_opt=paths=source_relative \
		--go-grpc_out=billing_service\
		--go-grpc_opt=paths=source_relative \
		proto_contracts/billing.proto