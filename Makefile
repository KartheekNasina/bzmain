# Makefile

.PHONY: all build tag push

# Image and repository details
IMAGE_NAME = bz-main
REPO_NAME = registry.digitalocean.com/bz-repo
IMAGE_TAG = v1.1

all: build tag push

build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f ./bz-main.dockerfile .

# Authenticate with DigitalOcean's Docker registry
login:
	@echo "Logging in to DigitalOcean Container Registry..."
	doctl registry login --never-expire

# Push the image to DigitalOcean's Docker registry
push: login
	@echo "Pushing image to DigitalOcean Container Registry..."
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(REPO_NAME)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(REPO_NAME)/$(IMAGE_NAME):$(IMAGE_TAG)

# All-in-one command to build, login, and push
all: build push
