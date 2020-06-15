package race

import (
	"fmt"
	"sync"
	"testing"
)

type Data struct {
	Name string
}

var dataPool sync.Pool

func init() {
	dataPool.New = func() interface{} {
		return &Data{
			Name: "",
		}
	}
}
func TestRace(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			data := dataPool.Get().(*Data)
			defer dataPool.Put(data)
			data.Name = fmt.Sprintf("name_%d", index)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
