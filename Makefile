build_and_push:
	docker build -t registry.hop.io/capy/api:latest . --platform=linux/amd64
	docker push registry.hop.io/capy/api

build_and_push_staging:
	docker build -t registry-stg.hop.io/capy/api:latest . --platform=linux/amd64
	docker push registry-stg.hop.io/capy/api