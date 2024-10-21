bin_name=cloudflareips
target=cmd/main.go

all: build

run:
	go run $(target)

build:
	go build -trimpath -ldflags="-s" -o build/$(bin_name) $(target)

clean:
	rm -rf build
