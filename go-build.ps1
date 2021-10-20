go build -ldflags "-H windowsgui -X main.version=v1.1.1 -X main.revision=$(git rev-parse --short HEAD)"
