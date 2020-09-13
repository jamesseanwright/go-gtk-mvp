build:
	@mkdir -p dist
	@glib-compile-resources .gresource.xml --target=dist/resources

	# -ldflags "-s" strips any debug symbols from the output
	@go build -o dist/main -ldflags "-s"
