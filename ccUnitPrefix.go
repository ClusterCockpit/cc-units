package ccunits

import "regexp"

type Prefix float64

const (
	InvalidPrefix Prefix = iota
	Base                 = 1
	Yotta                = 1e24
	Zetta                = 1e21
	Exa                  = 1e18
	Peta                 = 1e15
	Tera                 = 1e12
	Giga                 = 1e9
	Mega                 = 1e6
	Kilo                 = 1e3
	Milli                = 1e-3
	Micro                = 1e-6
	Nano                 = 1e-9
	Kibi                 = 1024
	Mebi                 = 1024 * 1024
	Gibi                 = 1024 * 1024 * 1024
	Tebi                 = 1024 * 1024 * 1024 * 1024
	Pebi                 = 1024 * 1024 * 1024 * 1024 * 1024
	Exbi                 = 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	Zebi                 = 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	Yobi                 = 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024
)
const prefixRegexStr = `^([kKmMgGtTpP]?[i]?)(.*)`

var prefixRegex = regexp.MustCompile(prefixRegexStr)

func (s *Prefix) String() string {
	switch *s {
	case InvalidPrefix:
		return "Inval"
	case Base:
		return ""
	case Kilo:
		return "Kilo"
	case Mega:
		return "Mega"
	case Giga:
		return "Giga"
	case Tera:
		return "Tera"
	case Peta:
		return "Peta"
	case Exa:
		return "Exa"
	case Zetta:
		return "Zetta"
	case Yotta:
		return "Yotta"
	case Milli:
		return "Milli"
	case Micro:
		return "Micro"
	case Nano:
		return "Nano"
	case Kibi:
		return "Kibi"
	case Mebi:
		return "Mebi"
	case Gibi:
		return "Gibi"
	case Tebi:
		return "Tebi"
	default:
		return "Unkn"
	}
}

func (s *Prefix) Prefix() string {
	switch *s {
	case InvalidPrefix:
		return "<inval>"
	case Base:
		return ""
	case Kilo:
		return "K"
	case Mega:
		return "M"
	case Giga:
		return "G"
	case Tera:
		return "T"
	case Peta:
		return "P"
	case Exa:
		return "E"
	case Zetta:
		return "Z"
	case Yotta:
		return "Y"
	case Milli:
		return "m"
	case Micro:
		return "u"
	case Nano:
		return "n"
	case Kibi:
		return "Ki"
	case Mebi:
		return "Mi"
	case Gibi:
		return "Gi"
	case Tebi:
		return "Ti"
	default:
		return "<unkn>"
	}
}

func NewPrefix(prefix string) Prefix {
	switch prefix {
	case "k":
		return Kilo
	case "K":
		return Kilo
	case "m":
		return Milli
	case "M":
		return Mega
	case "g":
		return Giga
	case "G":
		return Giga
	case "t":
		return Tera
	case "T":
		return Tera
	case "p":
		return Peta
	case "P":
		return Peta
	case "e":
		return Exa
	case "E":
		return Exa
	case "z":
		return Zetta
	case "Z":
		return Zetta
	case "y":
		return Yotta
	case "Y":
		return Yotta
	case "u":
		return Micro
	case "n":
		return Nano
	case "ki":
		return Kibi
	case "Ki":
		return Kibi
	case "Mi":
		return Mebi
	case "gi":
		return Gibi
	case "Gi":
		return Gibi
	case "Ti":
		return Tebi
	case "":
		return Base
	default:
		return InvalidPrefix
	}
}
