SHELL :=/bin/bash
CWD := $(PWD)

build-local:
	@docker build \
	-t app/admin \
    --target dev .

build-dev:
	@docker build \
	-t app/admin \
    --target dev .

build-prod:
	@docker build \
	-t app/admin \
	--target prod .

run-local:
	@docker run \
	-p 3000:3000 \
  	--env-file=.env \
  	app/admin
	  
build-app:
	@exec go build

.PHONY: all
