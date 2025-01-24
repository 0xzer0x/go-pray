package common

import (
	"fmt"
	"strings"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/util"
)

var Methods = map[string]calc.CalculationMethod{
	"MWL":       calc.MUSLIM_WORLD_LEAGUE,
	"EGYPT":     calc.EGYPTIAN,
	"KARACHI":   calc.KARACHI,
	"UAQ":       calc.UMM_AL_QURA,
	"DUBAI":     calc.DUBAI,
	"ISNA":      calc.NORTH_AMERICA,
	"KUWAIT":    calc.KUWAIT,
	"QATAR":     calc.QATAR,
	"SINGAPORE": calc.SINGAPORE,
	"UOIF":      calc.UOIF,
	"OTHER":     calc.OTHER,
}

func CalculationMethod(name string) (calc.CalculationMethod, error) {
	if val, ok := Methods[name]; ok {
		return val, nil
	}

	return 0, fmt.Errorf(
		"invalid method '%s', valid methods are: %s",
		name,
		strings.Join(util.MapKeys(Methods), ", "),
	)
}
