IMAGE=b1ame/k8s-go-app
TAG?=v1
PLATFORM=linux/amd64

build:
	docker buildx build --platform ${PLATFORM} -t ${IMAGE}:${TAG} --push .
