go-gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/oms.proto

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...    


# helm-install: ## Install all Helm charts in the specified namespace.
#     helm install omsv3-gateway deployments/charts/gateway --namespace omsv3
#     helm install omsv3-payments deployments/charts/payments --namespace omsv3
#     helm install omsv3-orders deployments/charts/orders --namespace omsv3
#     helm install omsv3-stock deployments/charts/stock --namespace omsv3    