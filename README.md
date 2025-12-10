# log-provider with OPA policies

Started with an empty directory and executed

- `go mod init github.com/barkhabansal/log-provider`
- `go get github.com/hashicorp/terraform-plugin-framework@v1.11.0`
- `go get github.com/hashicorp/terraform-plugin-log@v0.9.0`
- `go mod tidy`
- `go install .`

Create a provider directory
Linux/macOS: ~/.terraform.d/plugins/
Windows: %APPDATA%\terraform.d\plugins\

For macOS to publish provider locally
```
GOOS=$(uname -s | tr '[:upper:]' '[:lower:]')
GOARCH=$(uname -m | sed 's/x86_64/amd64/; s/aarch64/arm64/')
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/local/log-provider/0.0.1/${GOOS}_${GOARCH}
go build -o ~/.terraform.d/plugins/registry.terraform.io/local/log-provider/0.0.1/${GOOS}_${GOARCH}/terraform-provider-log-provider_v0.0.1
chmod +x ~/.terraform.d/plugins/registry.terraform.io/local/log-provider/0.0.1/darwin_arm64/*
```
