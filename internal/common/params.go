package common

import (
	"fmt"
	"strings"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/util"
)

var Madhabs map[string]calc.AsrJuristicMethod = map[string]calc.AsrJuristicMethod{
	"HANAFI": calc.HANAFI,
	"SHAFI":  calc.SHAFI_HANBALI_MALIKI,
}

var HighLatitudeRules map[string]calc.HighLatitudeRule = map[string]calc.HighLatitudeRule{
	"MIDDLE":   calc.MIDDLE_OF_THE_NIGHT,
	"SEVENTH":  calc.SEVENTH_OF_THE_NIGHT,
	"TWILIGHT": calc.TWILIGHT_ANGLE,
}

func Madhab(name string) (calc.AsrJuristicMethod, error) {
	if val, ok := Madhabs[name]; ok {
		return val, nil
	}

	return 0, fmt.Errorf(
		"invalid madhab '%s', valid madhabs are: %s",
		name,
		strings.Join(util.MapKeys(Madhabs), ", "),
	)
}

func HighLatitudeRule(name string) (calc.HighLatitudeRule, error) {
	if val, ok := HighLatitudeRules[name]; ok {
		return val, nil
	}

	return 0, fmt.Errorf(
		"invalid high latitude rule '%s', valid values are: %s",
		name,
		strings.Join(util.MapKeys(HighLatitudeRules), ", "),
	)
}
