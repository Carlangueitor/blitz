test:
	go test ./...

watch-n-test:
	fswatch -rx \
		--event=Updated --event=Created --event=Removed --event=Renamed \
		./*/**.go \
		2>/dev/null \
		| xargs -n1 -I{} make test

serve:
	go run cmd/blitz/main.go

watch-n-serve:
	go run cmd/blitz/main.go & echo $$! > server.pid
	fswatch -rx \
		--event=Updated --event=Created --event=Removed --event=Renamed \
		./**/*.go 2>/dev/null \
		| xargs -I{} sh -c \
		'if [ -f server.pid ]; then pkill -P $$(cat server.pid); fi; go run cmd/blitz/main.go & echo $$! > server.pid'
