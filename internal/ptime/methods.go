package ptime

import (
	"fmt"
	"strings"

	"github.com/mnadev/adhango/pkg/calc"
)

var methods = map[string]calc.CalculationMethod{
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
}

func CalculationMethod(name string) (calc.CalculationMethod, error) {
	if val, ok := methods[name]; ok {
		return val, nil
	}
	var methodNames []string = make([]string, 0, len(methods))
	for k := range methods {
		methodNames = append(methodNames, k)
	}
	return 0, fmt.Errorf(
		"invalid method '%s', valid methods are: %s",
		name,
		strings.Join(methodNames, ", "),
	)
}
