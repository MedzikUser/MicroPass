GO := go

build:
	$(GO) build  \
		-o target/bytepass-$$($(GO) env GOOS)-$$($(GO) env GOARCH) \
		-ldflags '-X github.com/bytepass/server/config.BuildType=production' \
		.

release:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=freebsd GOARCH=amd64 make build
	GOOS=freebsd GOARCH=arm64 make build

clean:
	$(GO) clean
	rm -rf target
