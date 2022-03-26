package switch_learn

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	var num1 = 100
	switch num1 {
	case 98, 99:
		fmt.Println("it is equery to 98")
	case 100:
		fmt.Println("it is equery to 100")
	default:
		fmt.Println("is is not equery 98 or 100")
	}
}
