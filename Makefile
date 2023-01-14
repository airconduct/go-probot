test:
	go test -v --race ./...

generate:
	cd web/dashboard && \
	npm run build
