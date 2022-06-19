package roller

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())

type Die struct {
	Sides int
}

func (d *Die) Roll() int {
	r := rand.New(seed)

	fmt.Println("Rolling die", d.Sides)

	return r.Intn(d.Sides) + 1
}

func DiceFromArg(arg string) []*Die {
	if arg[0] == 'd' {
		arg = "1" + arg
	}

	parts := strings.Split(arg, "d")

	dice := make([]*Die, 0)

	for i := 0; i < mustConvertInt(parts[0]); i++ {
		dice = append(dice, &Die{Sides: mustConvertInt(parts[1])})
	}

	return dice
}

func mustConvertInt(val string) int {
	number, err := strconv.Atoi(val)

	if err != nil {
		panic(err)
	}

	return number
}
