package korname

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenerate(T *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		fmt.Println(Generate(Male))
	}
	for i := 0; i < 50; i++ {
		fmt.Println(Generate(Female))
	}
}
