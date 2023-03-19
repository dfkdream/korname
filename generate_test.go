package korname_test

import (
	"fmt"
	"github.com/dfkdream/korname"
	"math/rand"
	"testing"
	"time"
)

func TestGenerate(T *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		fmt.Println(korname.Generate(korname.Male))
	}
	for i := 0; i < 50; i++ {
		fmt.Println(korname.Generate(korname.Female))
	}
}

func Example() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(korname.Generate(korname.Male))
	fmt.Println(korname.Generate(korname.Female))
	// output: random korean male name
	// random korean female name
}
