package ccunits

import "regexp"

type Measure int

const (
	InvalidMeasure Measure = iota
	Bytes
	Flops
	Percentage
	TemperatureC
	TemperatureF
	Rotation
	Hertz
	Time
	Watt
	Joule
	Cycles
	Requests
	Packets
	Events
)

// String returns the long string for the measure like 'Percent' or 'Seconds'
func (m *Measure) String() string {
	switch *m {
	case Bytes:
		return "byte"
	case Flops:
		return "Flops"
	case Percentage:
		return "Percent"
	case TemperatureC:
		return "DegreeC"
	case TemperatureF:
		return "DegreeF"
	case Rotation:
		return "RPM"
	case Hertz:
		return "Hertz"
	case Time:
		return "Seconds"
	case Watt:
		return "Watts"
	case Joule:
		return "Joules"
	case Cycles:
		return "Cycles"
	case Requests:
		return "Requests"
	case Packets:
		return "Packets"
	case Events:
		return "Events"
	case InvalidMeasure:
		return "Invalid"
	default:
		return "Unknown"
	}
}

// Short returns the short string for the measure like 'B' (Bytes), 's' (Time) or 'W' (Watt). Is is recommened to use Short() over String().
func (m *Measure) Short() string {
	switch *m {
	case Bytes:
		return "B"
	case Flops:
		return "Flops"
	case Percentage:
		return "Percent"
	case TemperatureC:
		return "degC"
	case TemperatureF:
		return "degF"
	case Rotation:
		return "RPM"
	case Hertz:
		return "Hz"
	case Time:
		return "s"
	case Watt:
		return "W"
	case Joule:
		return "J"
	case Cycles:
		return "cyc"
	case Requests:
		return "requests"
	case Packets:
		return "packets"
	case Events:
		return "events"
	case InvalidMeasure:
		return "Invalid"
	default:
		return "Unknown"
	}
}

const bytesRegexStr = `^([bB][yY]?[tT]?[eE]?[sS]?)`
const flopsRegexStr = `^([fF][lL]?[oO]?[pP]?[sS]?)`
const percentRegexStr = `^(%|[pP]ercent)`
const degreeCRegexStr = `^(deg[Cc]|°[cC])`
const degreeFRegexStr = `^(deg[fF]|°[fF])`
const rpmRegexStr = `^([rR][pP][mM])`
const hertzRegexStr = `^([hH][eE]?[rR]?[tT]?[zZ])`
const timeRegexStr = `^([sS][eE]?[cC]?[oO]?[nN]?[dD]?[sS]?)`
const wattRegexStr = `^([wW][aA]?[tT]?[tT]?[sS]?)`
const jouleRegexStr = `^([jJ][oO]?[uU]?[lL]?[eE]?[sS]?)`
const cyclesRegexStr = `^([cC][yY][cC]?[lL]?[eE]?[sS]?)`
const requestsRegexStr = `^([rR][eE][qQ][uU]?[eE]?[sS]?[tT]?[sS]?)`
const packetsRegexStr = `^([pP][aA]?[cC]?[kK][eE]?[tT][sS]?)`
const eventsRegexStr = `^([eE][vV]?[eE]?[nN][tT][sS]?)`

var bytesRegex = regexp.MustCompile(bytesRegexStr)
var flopsRegex = regexp.MustCompile(flopsRegexStr)
var percentRegex = regexp.MustCompile(percentRegexStr)
var degreeCRegex = regexp.MustCompile(degreeCRegexStr)
var degreeFRegex = regexp.MustCompile(degreeFRegexStr)
var rpmRegex = regexp.MustCompile(rpmRegexStr)
var hertzRegex = regexp.MustCompile(hertzRegexStr)
var timeRegex = regexp.MustCompile(timeRegexStr)
var wattRegex = regexp.MustCompile(wattRegexStr)
var jouleRegex = regexp.MustCompile(jouleRegexStr)
var cyclesRegex = regexp.MustCompile(cyclesRegexStr)
var requestsRegex = regexp.MustCompile(requestsRegexStr)
var packetsRegex = regexp.MustCompile(packetsRegexStr)
var eventsRegex = regexp.MustCompile(eventsRegexStr)

// NewMeasure creates a new measure out of a string representing a measure like 'Bytes', 'Flops' and 'precent'.
// It uses regular expressions for matching.
func NewMeasure(unit string) Measure {
	var match []string
	match = bytesRegex.FindStringSubmatch(unit)
	if match != nil {
		return Bytes
	}
	match = flopsRegex.FindStringSubmatch(unit)
	if match != nil {
		return Flops
	}
	match = percentRegex.FindStringSubmatch(unit)
	if match != nil {
		return Percentage
	}
	match = degreeCRegex.FindStringSubmatch(unit)
	if match != nil {
		return TemperatureC
	}
	match = degreeFRegex.FindStringSubmatch(unit)
	if match != nil {
		return TemperatureF
	}
	match = rpmRegex.FindStringSubmatch(unit)
	if match != nil {
		return Rotation
	}
	match = hertzRegex.FindStringSubmatch(unit)
	if match != nil {
		return Hertz
	}
	match = timeRegex.FindStringSubmatch(unit)
	if match != nil {
		return Time
	}
	match = cyclesRegex.FindStringSubmatch(unit)
	if match != nil {
		return Cycles
	}
	match = wattRegex.FindStringSubmatch(unit)
	if match != nil {
		return Watt
	}
	match = jouleRegex.FindStringSubmatch(unit)
	if match != nil {
		return Joule
	}
	match = requestsRegex.FindStringSubmatch(unit)
	if match != nil {
		return Requests
	}
	match = packetsRegex.FindStringSubmatch(unit)
	if match != nil {
		return Packets
	}
	match = eventsRegex.FindStringSubmatch(unit)
	if match != nil {
		return Events
	}
	return InvalidMeasure
}
