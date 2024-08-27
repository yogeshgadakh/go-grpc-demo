# Define variables
PROJECT_NAME=go-grpc-demo
PROTO_IMAGE_NAME=$(PROJECT_NAME)-proto
SERVER_IMAGE_NAME=$(PROJECT_NAME)-server
CLIENT_IMAGE_NAME=$(PROJECT_NAME)-client
GOLANG_IMAGE_NAME=golang:1.23.0-alpine3.19


HOST_WORKDIR=$(PWD)
CONTAINER_WORKDIR=/home/dsci/$(PROJECT_NAME)
CONTAINER_PROTO_WORKDIR=/opt


# Run the Docker container
run: train_ticket.pb.go docker-build-server docker-build-client
	@docker-compose up


##
# Clean
# used for cleaning up make targets and docker images, run make clean if you need to rebuild your docker image
##
clean:
	-@rm -f docker-build*

clean-all:
	@make clean
	@make clean-images

clean-images:
	@docker system prune -a --filter label=project-name="$(PROJECT_NAME)" -f


###
# Building/updating the protobuf files
# You shouldn't need to run this for the interview take home task
###
docker-build-proto:Dockerfile.proto
	@docker build \
		--build-arg BUILDER_IMAGE_NAME=$(GOLANG_IMAGE_NAME) \
		--build-arg PROJECT_NAME=$(PROJECT_NAME) \
		--tag $(PROTO_IMAGE_NAME) \
		--file Dockerfile.proto \
		.
	@touch $@	

train_ticket.pb.go: docker-build-proto $(HOST_WORKDIR)/protos/train_ticket.proto
	@docker run \
		--rm \
		--volume $(HOST_WORKDIR):$(CONTAINER_PROTO_WORKDIR) \
		$(PROTO_IMAGE_NAME) \
		protoc --go_out=$(CONTAINER_PROTO_WORKDIR)/protos/ --go-grpc_out=$(CONTAINER_PROTO_WORKDIR)/protos/ --proto_path=$(CONTAINER_PROTO_WORKDIR)/protos/ train_ticket.proto 

docker-build-server: Dockerfile.server
	@docker build --no-cache \
		--build-arg BUILDER_IMAGE_NAME=$(GOLANG_IMAGE_NAME) \
		--build-arg PROJECT_NAME=$(PROJECT_NAME) \
		--tag $(SERVER_IMAGE_NAME) \
		--file Dockerfile.server \
		--no-cache \
		--target deploy \
		.
	@touch $@

docker-build-client: Dockerfile.client
	@docker build --no-cache \
		--build-arg BUILDER_IMAGE_NAME=$(GOLANG_IMAGE_NAME) \
		--build-arg PROJECT_NAME=$(PROJECT_NAME) \
		--tag $(CLIENT_IMAGE_NAME) \
		--file Dockerfile.client \
		--no-cache \
		--target deploy \
		.
	@touch $@
