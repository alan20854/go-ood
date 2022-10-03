package main

import (
	_ "embed"
	"github.com/ronna-s/go-ood/pkg/pnpdev"

	"github.com/ronna-s/go-ood/pkg/pnp"
	engine "github.com/ronna-s/go-ood/pkg/pnp/engine/simple"
)

func main() {
	game := pnp.New(pnpdev.NewMinion(), pnpdev.NewGopher(), pnpdev.NewRubyist())
	//game := pnp.New()
	game.Run(engine.Engine{})
}
