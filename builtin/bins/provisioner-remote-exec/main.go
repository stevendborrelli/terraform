package main

import (
	"github.com/hashicorp/terraform/builtin/provisioners/remote-exec"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(new(remoteexec.ResourceProvisioner))
}
