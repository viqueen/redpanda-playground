.PHONY: build mocks

build:
	docker run --rm \
		--volume "${PWD}":/go/src \
		--volume /var/run/docker.sock:/var/run/docker.sock \
		--workdir /go/src \
		goreleaser/goreleaser:v2.4.4 build --clean --snapshot


mocks:
	docker run --rm \
		--volume "${PWD}":/go/src \
		--workdir /go/src \
		vektra/mockery:v2.46 --all

rpk-connect-create:
	 docker run --rm \
	 	docker.redpanda.com/redpandadata/connect \
	 	create "$@" > connect.yaml

rpk-connect-run:
	 docker run --rm -it \
	 	--network="host" \
		-v "${PWD}"/connect.yaml:/connect.yaml \
		docker.redpanda.com/redpandadata/connect \
		run