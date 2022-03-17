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
	AddDivisorUnit(div Measure)
	getPrefix() Prefix
	getMeasure() Measure
	getDivMeasure() Measure
	setPrefix(p Prefix)
}

var INVALID_UNIT = NewUnit("foobar")

func (u *unit) Valid() bool {
	return u.measure != InvalidMeasure
}

func (u *unit) String() string {
	if u.divMeasure != InvalidMeasure {
		return fmt.Sprintf("%s%s/%s", u.prefix.String(), u.measure.String(), u.divMeasure.String())
	} else {
		return fmt.Sprintf("%s%s", u.prefix.String(), u.measure.String())
	}
}

func (u *unit) Short() string {
	if u.divMeasure != InvalidMeasure {
		return fmt.Sprintf("%s%s/%s", u.prefix.Prefix(), u.measure.Short(), u.divMeasure.Short())
	} else {
		return fmt.Sprintf("%s%s", u.prefix.Prefix(), u.measure.Short())
	}
}

func (u *unit) AddDivisorUnit(div Measure) {
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

func (u *unit) getDivMeasure() Measure {
	return u.divMeasure
}

func GetPrefixPrefixFactor(in Prefix, out Prefix) func(value float64) float64 {
	var factor = 1.0
	var in_prefix = float64(in)
	var out_prefix = float64(out)
	factor = in_prefix / out_prefix
	return func(value float64) float64 { return factor }
}

func GetPrefixStringPrefixStringFactor(in string, out string) func(value float64) float64 {
	var i Prefix = NewPrefix(in)
	var o Prefix = NewPrefix(out)
	return GetPrefixPrefixFactor(i, o)
}

func GetUnitPrefixFactor(in Unit, out Prefix) (func(value float64) float64, Unit) {
	outUnit := NewUnit(in.Short())
	if outUnit.Valid() {
		outUnit.setPrefix(out)
		conv := GetPrefixPrefixFactor(in.getPrefix(), out)
		return conv, outUnit
	}
	return nil, INVALID_UNIT
}

func GetUnitPrefixStringFactor(in Unit, out string) (func(value float64) float64, Unit) {
	var o Prefix = NewPrefix(out)
	return GetUnitPrefixFactor(in, o)
}

func GetUnitStringPrefixStringFactor(in string, out string) (func(value float64) float64, Unit) {
	var i = NewUnit(in)
	return GetUnitPrefixStringFactor(i, out)
}

func GetUnitUnitFactor(in Unit, out Unit) (func(value float64) float64, error) {
	if in.getMeasure() == TemperatureC && out.getMeasure() == TemperatureF {
		return func(value float64) float64 { return (value * 1.8) + 32 }, nil
	} else if in.getMeasure() == TemperatureF && out.getMeasure() == TemperatureC {
		return func(value float64) float64 { return (value - 32) / 1.8 }, nil
	} else if in.getMeasure() != out.getMeasure() || in.getDivMeasure() != out.getDivMeasure() {
		return func(value float64) float64 { return 1.0 }, fmt.Errorf("invalid measures in in and out Unit")
	}
	return GetPrefixPrefixFactor(in.getPrefix(), out.getPrefix()), nil
}

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
