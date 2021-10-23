package main

import (
	"github.com/lyft/clutch/backend/cmd/assets"
	"github.com/lyft/clutch/backend/gateway"
	ecsmod "github.com/lyft/clutch/backend/module/ecs"
)

func main() {
	flags := gateway.ParseFlags()
	components := gateway.CoreComponentFactory
	components.Modules[ecsmod.Name] = ecsmod.New

	gateway.Run(flags, components, assets.VirtualFS)
}
