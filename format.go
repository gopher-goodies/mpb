package mpb

import (
	"fmt"
	"strings"
)

const (
	_          = iota
	bytesInKiB = 1 << (iota * 10)
	bytesInMiB
	bytesInGiB
	bytesInTiB
)

type Units uint

const (
	UnitNone Units = iota
	UnitBytes
)

func Format(i int) *formatter {
	return &formatter{n: i}
}

type formatter struct {
	n     int
	unit  Units
	width int
}

func (f *formatter) To(unit Units) *formatter {
	f.unit = unit
	return f
}

func (f *formatter) Width(width int) *formatter {
	f.width = width
	return f
}

func (f *formatter) String() string {
	switch f.unit {
	case UnitBytes:
		return formatBytes(f.n)
	default:
		return fmt.Sprintf(fmt.Sprintf("%%%dd", f.width), f.n)
	}
}

func formatBytes(i int) (result string) {
	switch {
	case i > bytesInTiB:
		result = fmt.Sprintf("%.02fTiB", float64(i)/bytesInTiB)
	case i > bytesInGiB:
		result = fmt.Sprintf("%.02fGiB", float64(i)/bytesInGiB)
	case i > bytesInMiB:
		result = fmt.Sprintf("%.02fMiB", float64(i)/bytesInMiB)
	case i > bytesInKiB:
		result = fmt.Sprintf("%.02fKiB", float64(i)/bytesInKiB)
	default:
		result = fmt.Sprintf("%db", i)
	}
	result = strings.Trim(result, " ")
	return
}