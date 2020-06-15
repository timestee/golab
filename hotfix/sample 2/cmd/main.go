package main

import (
	"sample/player"
	"sample/plugin"
)

func main() {
	p := plugin.NewPlugin()
	err := p.Build("/tmp/patch.so", &plugin.Config{
		Name:    "patch",
		Type:    "patch",
		Path:    "sample/patch",
		NewFunc: "Patch",
	})

	if err != nil {
		panic(err)
	}

	pp := &player.Player{}
	pp.World()
	c, err := plugin.Load("/tmp/patch.so")
	if err != nil {
		panic(err)
	}
	err = plugin.Init(c)
	if err != nil {
		panic(err)
	}
	pp.World()
}
