package ccunits

import (
	"fmt"
	"strings"
)

type unit struct {
	prefix     Prefix
	measure    Measure
	divMeasure Measure
}

type Unit interface {
	Valid() bool
	String() string
	Short() string
	AddUnitDenominator(div Measure)
	getPrefix() Prefix
	getMeasure() Measure
	getUnitDenominator() Measure
	setPrefix(p Prefix)
}

var INVALID_UNIT = NewUnit("foobar")

// Check whether a unit is a valid unit. It requires at least a prefix and a measure. The unit denominator is optional
func (u *unit) Valid() bool {
	return u.prefix != InvalidPrefix && u.measure != InvalidMeasure
}

// Get the long string for the unit like 'KiloHertz'
func (u *unit) String() string {
	if u.divMeasure != InvalidMeasure {
		return fmt.Sprintf("%s%s/%s", u.prefix.String(), u.measure.String(), u.divMeasure.String())
	} else {
		return fmt.Sprintf("%s%s", u.prefix.String(), u.measure.String())
	}
}

// Get the short string for the unit like 'KHz' (recommened over String()!)
func (u *unit) Short() string {
	if u.divMeasure != InvalidMeasure {
		return fmt.Sprintf("%s%s/%s", u.prefix.Prefix(), u.measure.Short(), u.divMeasure.Short())
	} else {
		return fmt.Sprintf("%s%s", u.prefix.Prefix(), u.measure.Short())
	}
}

//
func (u *unit) AddUnitDenominator(div Measure) {
	u.divMeasure = div
}

func (u *unit) getPrefix() Prefix {
	return u.prefix
}

func (u *unit) setPrefix(p Prefix) {
	u.prefix = p
}

func (u *unit) getMeasure() Measure {
	return u.measure
}

func (u *unit) getUnitDenominator() Measure {
	return u.divMeasure
}

// This creates the default conversion function between two prefixes
func GetPrefixPrefixFactor(in Prefix, out Prefix) func(value interface{}) interface{} {
	var factor = 1.0
	var in_prefix = float64(in)
	var out_prefix = float64(out)
	factor = in_prefix / out_prefix
	conv := func(value interface{}) interface{} {
		switch v := value.(type) {
		case float64:
			return v * factor
		case float32:
			return float32(float64(v) * factor)
		case int:
			return int(float64(v) * factor)
		case int32:
			return int32(float64(v) * factor)
		case int64:
			return int64(float64(v) * factor)
		case uint:
			return uint(float64(v) * factor)
		case uint32:
			return uint32(float64(v) * factor)
		case uint64:
			return uint64(float64(v) * factor)
		}
		return value
	}
	return conv
}

// This is the conversion function between temperatures in Celsius to Fahrenheit
func convertTempC2TempF(value interface{}) interface{} {
	switch v := value.(type) {
	case float64:
		return (v * 1.8) + 32
	case float32:
		return (v * 1.8) + 32
	case int:
		return int((float64(v) * 1.8) + 32)
	case int32:
		return int32((float64(v) * 1.8) + 32)
	case int64:
		return int64((float64(v) * 1.8) + 32)
	case uint:
		return uint((float64(v) * 1.8) + 32)
	case uint32:
		return uint32((float64(v) * 1.8) + 32)
	case uint64:
		return uint64((float64(v) * 1.8) + 32)
	}
	return value
}

// This is the conversion function between temperatures in Fahrenheit to Celsius
func convertTempF2TempC(value interface{}) interface{} {
	switch v := value.(type) {
	case float64:
		return (v - 32) / 1.8
	case float32:
		return (v - 32) / 1.8
	case int:
		return int(((float64(v) - 32) / 1.8))
	case int32:
		return int32(((float64(v) - 32) / 1.8))
	case int64:
		return int64(((float64(v) - 32) / 1.8))
	case uint:
		return uint(((float64(v) - 32) / 1.8))
	case uint32:
		return uint32(((float64(v) - 32) / 1.8))
	case uint64:
		return uint64(((float64(v) - 32) / 1.8))
	}
	return value
}

// If we only have strings with the prefixes, this is a handy wrapper
func GetPrefixStringPrefixStringFactor(in string, out string) func(value interface{}) interface{} {
	var i Prefix = NewPrefix(in)
	var o Prefix = NewPrefix(out)
	return GetPrefixPrefixFactor(i, o)
}

// Get the conversion function and resulting unit for a unit and a prefix
func GetUnitPrefixFactor(in Unit, out Prefix) (func(value interface{}) interface{}, Unit) {
	outUnit := NewUnit(in.Short())
	if outUnit.Valid() {
		outUnit.setPrefix(out)
		conv := GetPrefixPrefixFactor(in.getPrefix(), out)
		return conv, outUnit
	}
	return nil, INVALID_UNIT
}

// Get the conversion function and resulting unit for a unit and a prefix as string
func GetUnitPrefixStringFactor(in Unit, out string) (func(value interface{}) interface{}, Unit) {
	var o Prefix = NewPrefix(out)
	return GetUnitPrefixFactor(in, o)
}

// Get the conversion function and resulting unit for a unit and a prefix when both are only string representations
func GetUnitStringPrefixStringFactor(in string, out string) (func(value interface{}) interface{}, Unit) {
	var i = NewUnit(in)
	return GetUnitPrefixStringFactor(i, out)
}

// Get the conversion function and (maybe) error for unit to unit conversion
func GetUnitUnitFactor(in Unit, out Unit) (func(value interface{}) interface{}, error) {
	if in.getMeasure() == TemperatureC && out.getMeasure() == TemperatureF {
		return convertTempC2TempF, nil
	} else if in.getMeasure() == TemperatureF && out.getMeasure() == TemperatureC {
		return convertTempF2TempC, nil
	} else if in.getMeasure() != out.getMeasure() || in.getUnitDenominator() != out.getUnitDenominator() {
		return func(value interface{}) interface{} { return 1.0 }, fmt.Errorf("invalid measures in in and out Unit")
	}
	return GetPrefixPrefixFactor(in.getPrefix(), out.getPrefix()), nil
}

// Create a new unit out of a string representing a unit like 'Mbyte/s' or 'GHz'.
func NewUnit(unitStr string) Unit {
	u := &unit{
		prefix:     InvalidPrefix,
		measure:    InvalidMeasure,
		divMeasure: InvalidMeasure,
	}
	matches := prefixRegex.FindStringSubmatch(unitStr)
	if len(matches) > 2 {
		pre := NewPrefix(matches[1])
		measures := strings.Split(matches[2], "/")
		m := NewMeasure(measures[0])
		// Special case for prefix 'p' or 'P' (Peta) and measures starting with 'p' or 'P'
		// like 'packets' or 'percent'. Same for 'e' or 'E' (Exa) for measures starting with
		// 'e' or 'E' like 'events'
		if m == InvalidMeasure {
			switch pre {
			case Peta, Exa:
				t := NewMeasure(matches[1] + measures[0])
				if t != InvalidMeasure {
					m = t
					pre = Base
				}
			}
		}
		div := InvalidMeasure
		if len(measures) > 1 {
			div = NewMeasure(measures[1])
		}

		switch m {
		// Special case for 'm' as prefix for Bytes and some others as thers is no unit like MilliBytes
		case Bytes, Flops, Packets, Events, Cycles, Requests:
			if pre == Milli {
				pre = Mega
			}
		// Special case for percentage. No/ignore prefix
		case Percentage:
			pre = Base
		}
		if pre != InvalidPrefix && m != InvalidMeasure {
			u.prefix = pre
			u.measure = m
			if div != InvalidMeasure {
				u.divMeasure = div
			}
		}
	}
	return u
}
