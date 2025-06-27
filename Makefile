APP_NAME=golang-local-kubernetes
IMAGE_NAME=$(APP_NAME):latest
KIND_CLUSTER_NAME=dev-cluster

.PHONY: build docker kind-create kind-load deploy

build:
	cd app && go build -o main .

docker:
	docker build -t $(IMAGE_NAME) ./app

kind-create:
	kind create cluster --name $(KIND_CLUSTER_NAME) --config kind/kind-config.yaml || true

kind-load:
	kind load docker-image $(IMAGE_NAME) --name $(KIND_CLUSTER_NAME)

deploy:
	kubectl apply -k deploy/base

all: build docker kind-create kind-load deploy