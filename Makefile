VERSION=0.0.1

.PHONY: proto
proto:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
    	--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
    	api/grpc/v1/usage.proto
	protoc \
		-I . \
		--grpc-gateway_out . \
    	--grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
    	--grpc-gateway_opt generate_unbound_methods=true \
        api/grpc/v1/usage.proto

.PHONY: server
server:
	go run cmd/server/main.go

.PHONY: gateway
gateway:
	go run cmd/gateway/main.go

.PHONY: boavizta
boavizta:
	docker compose up -d boavizta

.PHONY: test
test:
	go test -v $(shell go list ./... | grep -v /e2e)

.PHONY: test-e2e
test-e2e:
	go clean -testcache
	go test -v $(shell go list ./e2e/...)

.PHONY: vm-install
vm-install:
	cd ansible && ansible-playbook -i inventory/all.yml install.yml

.PHONY: vm-deploy
vm-deploy:
	cd ansible && ansible-playbook -i inventory/all.yml deploy.yml

.PHONY: dev-nginx
dev-nginx:
	NGINX_CONF_DIR=./nginx/conf-dev/ docker compose up nginx -d

.PHONY: dev-up
dev-up:
	NGINX_CONF_DIR=./nginx/conf-dev/ docker compose up -d
