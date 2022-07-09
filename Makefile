build_and_push:
	docker build -t registry.hop.io/capy/capy:latest . --platform=linux/amd64
	docker push registry.hop.io/capy/capy