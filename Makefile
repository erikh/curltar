IMAGE_NAME ?= erikh/curltar

image:
	docker build -t $(IMAGE_NAME) .
