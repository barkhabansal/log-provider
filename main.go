package main

import (
	"context"
	"fmt"
	"github.com/barkhabansal/log-provider/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var version = "0.0.1"

const address = "registry.terraform.io/local/log-provider"

func main() {
	fmt.Println("Terraform log provider")

	providerserver.Serve(
		context.Background(),
		provider.New(version),
		providerserver.ServeOpts{
			Address: address,
		})
}
