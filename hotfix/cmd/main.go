package main

import (
	"github.com/timestee/golab/hotfix/sample/player"
	"github.com/timestee/golab/hotfix/sample/plugin"
)

func main() {
	p := plugin.NewPlugin()
	err := p.Build("/tmp/patch.so", &plugin.Config{
		Name:    "patch",
		Type:    "patch",
		Path:    "github.com/timestee/golab/hotfix/sample/patch",
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
	pp.Set("test1")
	pp.World()
}
