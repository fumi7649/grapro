.PHONY: install-istio
install-istio:
	helm install istio-base istio/base -n istio-system --create-namespace
	helm install istiod istio/istiod -n istio-system --values ./oss/istio/values.yaml 
	helm install istio-ingressgateway istio/gateway -n istio-system  --values ./oss/istio/values.yaml 

.PHONY: uninstall-istio
uninstall-istio:
		helm uninstall istio-base istio/base -n istio-system 
		helm uninstall istiod istio/istiod -n istio-system 
		helm uninstall istio-ingressgateway istio/gateway -n istio-system 

.PHONY: apply-go-server
apply-go-server:
	kubectl apply -f ./application/manifest/server.yaml

.PHONY: apply-gateway
apply-gateway:
	kubectl apply -f ./application/manifest/istio-gateway.yaml

.PHONY: create-cluster
create-cluster:
	kind create cluster --config cluster.yaml

.PHONY: build-image
build-image:
	docker build ./application -t application-go:latest

.PHONY: push-image-to-kind
push-image-to-kind:
	kind load docker-image application-go:latest
