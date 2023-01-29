test:
	go test -v --race ./...

generate:
	go generate -x && \
	cd web/dashboard && \
	npm run build
