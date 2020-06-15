package player

import "fmt"

type Player struct {
	Name string
}

func (r *Player) Hello() {
	fmt.Print("Hello-->")
}

func (r *Player) World() {
	fmt.Println("World")
}

func (r *Player) setName(n string) error {
	r.Name = n
	return nil
}

func (r *Player) Set(n string) error {
	return r.setName(n)
}
