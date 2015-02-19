package bracmeister

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestSimple(t *testing.T) {
	pp.Println("E", Calc(168, 88, 56, true))
	pp.Println("E", Calc(168, 88, 56, false))

	pp.Println("G", Calc(162, 90, 58, true))
	pp.Println("G", Calc(162, 90, 58, false))

	pp.Println("E", Calc(162, 88, 58, true))
	pp.Println("E", Calc(162, 88, 58, false))

	pp.Println("E", Calc(161, 86, 58, true))
	pp.Println("E", Calc(161, 86, 58, false))

	pp.Println("D", Calc(160, 84, 57, true))
	pp.Println("D", Calc(160, 84, 57, false))

	pp.Println("F", Calc(170, 92, 62, true))
	pp.Println("F", Calc(170, 92, 62, false))

	pp.Println("G", Calc(164, 93, 59, true))
	pp.Println("G", Calc(164, 93, 59, false))
}
