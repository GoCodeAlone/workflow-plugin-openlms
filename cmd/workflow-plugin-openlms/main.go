package main

import (
	"github.com/GoCodeAlone/workflow-plugin-openlms/internal"
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

var version = "dev"

func main() {
	sdk.Serve(internal.NewOpenLMSPlugin(), sdk.WithBuildVersion(sdk.ResolveBuildVersion(internal.Version)))
}
