package patch

import (
	"fmt"
	"reflect"
	_ "unsafe"

	"github.com/timestee/golab/hotfix/sample/player"

	"bou.ke/monkey"
)

//go:linkname  setName github.com/timestee/golab/hotfix/sample/player.(*Player).setName
func setName(r *player.Player, n string) error

func FixWorld(r *player.Player) {
	setName(r, "patch ok")
	fmt.Println("Wonderful World ", r.Name)
}

func Patch() {
	var d *player.Player
	fmt.Println("patch exec")
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "World", FixWorld)
}
