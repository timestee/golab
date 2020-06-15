package patch

import (
	"fmt"

	"reflect"
	_ "unsafe"

	"sample/player"

	"bou.ke/monkey"
)

func FixWorld(r *player.Player) {
	fmt.Println("Wonderful World ", r.Name)
}

func Patch() {
	var d *player.Player
	fmt.Println("patch exec")
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "World", FixWorld)
}
