package roller

import (
	"fmt"
	"regexp"
)

var ptn = regexp.MustCompile(`(\d+)?d\d+`)

func ParseArgs(args ...string) ([]*Die, error) {
	dice := make([]*Die, 0)

	for i, arg := range args {
		if !ptn.MatchString(arg) {
			return nil, fmt.Errorf("arg %s at pos: %d is invalid", arg, i+1)
		}

		dice = append(dice, DiceFromArg(arg)...)
	}

	return dice, nil
}
