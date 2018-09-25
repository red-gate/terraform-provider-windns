package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/red-gate/terraform-provider-windns/windns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: windns.Provider,
	})
}
