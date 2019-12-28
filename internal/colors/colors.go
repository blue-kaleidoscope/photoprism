/*
This package encapsulates session storage.

Additional information can be found in our Developer Guide:

https://github.com/photoprism/photoprism/wiki
*/
package colors

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

type ColorPerception struct {
	Colors    Colors
	MainColor Color
	Luminance LightMap
	Chroma    Chroma
}

type Color uint16
type Colors []Color

type Chroma uint8
type Luminance uint8
type LightMap []Luminance

const (
	Black Color = iota
	Brown
	Grey
	White
	Purple
	Gold
	Blue
	Cyan
	Teal
	Green
	Lime
	Yellow
	Magenta
	Orange
	Red
	Pink
)

var All = Colors{
	Red,
	Magenta,
	Pink,
	Orange,
	Gold,
	Yellow,
	Lime,
	Green,
	Teal,
	Cyan,
	Blue,
	Purple,
	Brown,
	White,
	Grey,
	Black,
}

var Names = map[Color]string{
	Black:   "dark",    // 0
	Brown:   "brown",   // 1
	Grey:    "grey",    // 2
	White:   "bright",  // 3
	Purple:  "purple",  // 4
	Gold:    "gold",    // 5
	Blue:    "blue",    // 6
	Cyan:    "cyan",    // 7
	Teal:    "teal",    // 8
	Green:   "green",   // 9
	Lime:    "lime",    // A
	Yellow:  "yellow",  // B
	Magenta: "magenta", // C
	Orange:  "orange",  // D
	Red:     "red",     // E
	Pink:    "pink",    // F
}

var Weights = map[Color]uint16{
	Black:   2,
	Brown:   2,
	Grey:    1,
	White:   2,
	Purple:  4,
	Gold:    4,
	Blue:    3,
	Cyan:    4,
	Teal:    4,
	Green:   3,
	Lime:    5,
	Yellow:  5,
	Magenta: 5,
	Orange:  4,
	Red:     4,
	Pink:    4,
}

func (c Color) Name() string {
	return Names[c]
}

func (c Color) Hex() string {
	return fmt.Sprintf("%X", c)
}

func (c Colors) Hex() (result string) {
	for _, indexedColor := range c {
		result += indexedColor.Hex()
	}

	return result
}

func (c Colors) List() []map[string]string {
	result := make([]map[string]string, 0, len(c))

	for _, c := range c {
		result = append(result, map[string]string{"name": c.Name(), "label": strings.Title(c.Name()), "example": ColorExamples[c]})
	}

	return result
}

func (c Chroma) Hex() string {
	return fmt.Sprintf("%X", c)
}

func (c Chroma) Uint() uint {
	return uint(c)
}

func (c Chroma) Int() int {
	return int(c)
}

func (l Luminance) Hex() string {
	return fmt.Sprintf("%X", l)
}

func (m LightMap) Hex() (result string) {
	for _, luminance := range m {
		result += luminance.Hex()
	}

	return result
}

var ColorExamples = map[Color]string{
	Red:     "#E57373",
	Magenta: "#FF00FF",
	Pink:    "#F06292",
	Orange:  "#FFB74D",
	Brown:   "#A1887F",
	Gold:    "#FFD54F",
	Yellow:  "#FFF176",
	Lime:    "#DCE775",
	Green:   "#81C784",
	Teal:    "#4DB6AC",
	Cyan:    "#4DD0E1",
	Blue:    "#64B5F6",
	Purple:  "#BA68C8",
	White:   "#F5F5F5",
	Grey:    "#BDBDBD",
	Black:   "#333333",
}

var ColorMap = map[color.RGBA]Color{
	{0x00, 0x00, 0x00, 0xff}: Black,
	{0xa1, 0x88, 0x7f, 0xff}: Brown,
	{0x8d, 0x6e, 0x63, 0xff}: Brown,
	{0xa0, 0x7f, 0x6c, 0xff}: Brown,
	{0x9b, 0x7b, 0x5b, 0xff}: Brown,
	{0x75, 0x64, 0x5b, 0xff}: Brown,
	{0x79, 0x55, 0x48, 0xff}: Brown,
	{0x6d, 0x4c, 0x41, 0xff}: Brown,
	{0x5d, 0x40, 0x37, 0xff}: Brown,
	{0x9b, 0x61, 0x36, 0xff}: Brown,
	{0xc1, 0xa4, 0x87, 0xff}: Brown,
	{0xaa, 0x80, 0x62, 0xff}: Brown,
	{0x6b, 0x55, 0x46, 0xff}: Brown,
	{0xb4, 0xb5, 0x9c, 0xff}: Brown,
	{0xb2, 0xb4, 0x9b, 0xff}: Green,
	{0xe0, 0xe0, 0xe0, 0xff}: Grey,
	{0x9E, 0x9E, 0x9E, 0xff}: Grey,
	{0x75, 0x75, 0x75, 0xff}: Grey,
	{0x61, 0x61, 0x61, 0xff}: Grey,
	{0x42, 0x42, 0x42, 0xff}: Grey,
	{0x84, 0x7a, 0x72, 0xff}: Grey,
	{0xdf, 0xe0, 0xe1, 0xff}: Grey,
	{0xFF, 0xFF, 0xFF, 0xff}: White,
	{0xe4, 0xe4, 0xe4, 0xff}: White,
	{0xe7, 0xe7, 0xe7, 0xff}: White,
	{0xf3, 0xe5, 0xf5, 0xff}: Purple,
	{0xe1, 0xbe, 0xe7, 0xff}: Purple,
	{0xce, 0x93, 0xd8, 0xff}: Purple,
	{0xba, 0x68, 0xc8, 0xff}: Purple,
	{0xab, 0x47, 0xbc, 0xff}: Purple,
	{0x9c, 0x27, 0xb0, 0xff}: Purple,
	{0x9b, 0x31, 0x8f, 0xff}: Purple,
	{0x86, 0x00, 0x7e, 0xff}: Purple,
	{0x8e, 0x24, 0xaa, 0xff}: Purple,
	{0x7b, 0x1f, 0xa2, 0xff}: Purple,
	{0x6a, 0x1b, 0x9a, 0xff}: Purple,
	{0x4a, 0x14, 0x8c, 0xff}: Purple,
	{0xaa, 0x00, 0xff, 0xff}: Purple,
	{0xed, 0xe7, 0xf6, 0xff}: Purple,
	{0xd1, 0xc4, 0xe9, 0xff}: Purple,
	{0xb3, 0x9d, 0xdb, 0xff}: Purple,
	{0x95, 0x75, 0xcd, 0xff}: Purple,
	{0x7e, 0x57, 0xc2, 0xff}: Purple,
	{0x5e, 0x35, 0xb1, 0xff}: Purple,
	{0x67, 0x3a, 0xb7, 0xff}: Purple,
	{0x51, 0x2d, 0xa8, 0xff}: Purple,
	{0x45, 0x27, 0xa0, 0xff}: Purple,
	{0x31, 0x1b, 0x92, 0xff}: Purple,
	{0xb3, 0x88, 0xff, 0xff}: Purple,
	{0x7c, 0x4d, 0xff, 0xff}: Purple,
	{0x8e, 0x64, 0x93, 0xff}: Purple,
	{0x5e, 0x3a, 0x5e, 0xff}: Purple,
	{0x44, 0x0e, 0x79, 0xff}: Purple,
	{0x48, 0x36, 0x78, 0xff}: Purple,
	{0x4e, 0x38, 0x80, 0xff}: Purple,
	{0x3b, 0x0e, 0x79, 0xff}: Purple,
	{0x3F, 0x51, 0xB5, 0xff}: Blue,
	{0xc5, 0xca, 0xe9, 0xff}: Blue,
	{0x5c, 0x6b, 0xc0, 0xff}: Blue,
	{0x39, 0x49, 0xab, 0xff}: Blue,
	{0x30, 0x3f, 0x9f, 0xff}: Blue,
	{0x28, 0x35, 0x93, 0xff}: Blue,
	{0x1a, 0x23, 0x7e, 0xff}: Blue,
	{0x53, 0x6d, 0xfe, 0xff}: Blue,
	{0x3d, 0x5a, 0xfe, 0xff}: Blue,
	{0x30, 0x4f, 0xfe, 0xff}: Blue,
	{0x21, 0x96, 0xF3, 0xff}: Blue,
	{0xbb, 0xde, 0xfb, 0xff}: Blue,
	{0x90, 0xca, 0xf9, 0xff}: Blue,
	{0x64, 0xb5, 0xf6, 0xff}: Blue,
	{0x42, 0xa5, 0xf5, 0xff}: Blue,
	{0x1e, 0x88, 0xe5, 0xff}: Blue,
	{0x19, 0x76, 0xd2, 0xff}: Blue,
	{0x15, 0x65, 0xc0, 0xff}: Blue,
	{0x0d, 0x47, 0xa1, 0xff}: Blue,
	{0x82, 0xb1, 0xff, 0xff}: Blue,
	{0x44, 0x8a, 0xff, 0xff}: Blue,
	{0x29, 0x79, 0xff, 0xff}: Blue,
	{0x29, 0x62, 0xff, 0xff}: Blue,
	{0x03, 0xa9, 0xf6, 0xff}: Blue,
	{0xb3, 0xe5, 0xfc, 0xff}: Blue,
	{0x81, 0xd4, 0xfa, 0xff}: Blue,
	{0x4f, 0xc3, 0xf7, 0xff}: Blue,
	{0x29, 0xb6, 0xf6, 0xff}: Blue,
	{0x03, 0x9b, 0xe5, 0xff}: Blue,
	{0x02, 0x88, 0xd1, 0xff}: Blue,
	{0x02, 0x77, 0xbd, 0xff}: Blue,
	{0x01, 0x57, 0x9b, 0xff}: Blue,
	{0x80, 0xd8, 0xff, 0xff}: Blue,
	{0x40, 0xc4, 0xff, 0xff}: Blue,
	{0x00, 0xb0, 0xff, 0xff}: Blue,
	{0x00, 0x91, 0xea, 0xff}: Blue,
	{0x60, 0x7d, 0x8b, 0xff}: Blue,
	{0x78, 0x90, 0x9c, 0xff}: Blue,
	{0x54, 0x6e, 0x7a, 0xff}: Blue,
	{0x37, 0x47, 0x4f, 0xff}: Blue,
	{0xe4, 0xeb, 0xfd, 0xff}: Blue,
	{0x7d, 0xd3, 0xea, 0xff}: Blue,
	{0x07, 0x63, 0x99, 0xff}: Blue,
	{0x28, 0x44, 0x6b, 0xff}: Blue,
	{0x4a, 0xc8, 0xf5, 0xff}: Blue,
	{0x08, 0x00, 0xf4, 0xff}: Blue,
	{0x01, 0x2d, 0x5f, 0xff}: Blue,
	{0xb2, 0xeb, 0xf2, 0xff}: Cyan,
	{0x80, 0xde, 0xea, 0xff}: Cyan,
	{0x4d, 0xd0, 0xe1, 0xff}: Cyan,
	{0x26, 0xc6, 0xda, 0xff}: Cyan,
	{0x00, 0xb8, 0xd4, 0xff}: Cyan,
	{0x00, 0xBC, 0xD4, 0xff}: Cyan,
	{0x00, 0xac, 0xc1, 0xff}: Cyan,
	{0x00, 0x97, 0xa7, 0xff}: Cyan,
	{0x00, 0x83, 0x8f, 0xff}: Cyan,
	{0x00, 0x60, 0x64, 0xff}: Cyan,
	{0x84, 0xff, 0xff, 0xff}: Cyan,
	{0x18, 0xff, 0xff, 0xff}: Cyan,
	{0x00, 0xe5, 0xff, 0xff}: Cyan,
	{0x00, 0x96, 0x88, 0xff}: Teal,
	{0x00, 0x89, 0x7b, 0xff}: Teal,
	{0x00, 0x79, 0x6b, 0xff}: Teal,
	{0x00, 0x69, 0x5c, 0xff}: Teal,
	{0x04, 0x5d, 0x5c, 0xff}: Teal,
	{0x24, 0x5a, 0x5f, 0xff}: Teal,
	{0x03, 0x45, 0x4f, 0xff}: Teal,
	{0x2c, 0x54, 0x5e, 0xff}: Teal,
	{0x17, 0x47, 0x41, 0xff}: Teal,
	{0xe8, 0xf5, 0xe9, 0xff}: Green,
	{0xc8, 0xe6, 0xc9, 0xff}: Green,
	{0xab, 0xc7, 0xb0, 0xff}: Green,
	{0xa5, 0xd6, 0xa7, 0xff}: Green,
	{0x81, 0xc7, 0x84, 0xff}: Green,
	{0x66, 0xbb, 0x6a, 0xff}: Green,
	{0x4C, 0xAF, 0x50, 0xff}: Green,
	{0x43, 0xa0, 0x47, 0xff}: Green,
	{0x38, 0x8e, 0x3c, 0xff}: Green,
	{0x2e, 0x7d, 0x32, 0xff}: Green,
	{0x1b, 0x5e, 0x20, 0xff}: Green,
	{0xf1, 0xf8, 0xe9, 0xff}: Green,
	{0xdc, 0xed, 0xc8, 0xff}: Green,
	{0xc5, 0xe1, 0xa5, 0xff}: Green,
	{0xae, 0xd5, 0x81, 0xff}: Green,
	{0x8b, 0xc3, 0x4a, 0xff}: Green,
	{0x9c, 0xcc, 0x65, 0xff}: Green,
	{0x7c, 0xb3, 0x42, 0xff}: Green,
	{0x68, 0x9f, 0x38, 0xff}: Green,
	{0x55, 0x8b, 0x2f, 0xff}: Green,
	{0x33, 0x69, 0x1e, 0xff}: Green,
	{0xb9, 0xf6, 0xca, 0xff}: Green,
	{0x69, 0xf0, 0xae, 0xff}: Green,
	{0x00, 0xc8, 0x53, 0xff}: Green,
	{0x00, 0xe6, 0x76, 0xff}: Green,
	{0xcc, 0xff, 0x90, 0xff}: Green,
	{0xb2, 0xff, 0x59, 0xff}: Green,
	{0x76, 0xff, 0x03, 0xff}: Green,
	{0x64, 0xdd, 0x17, 0xff}: Green,
	{0xdd, 0xd5, 0x79, 0xff}: Green,
	{0xee, 0xec, 0xa2, 0xff}: Green,
	{0x24, 0x4e, 0x3b, 0xff}: Green,
	{0x9a, 0x9d, 0x47, 0xff}: Green,
	{0xbe, 0xbd, 0x76, 0xff}: Green,
	{0x5c, 0x5a, 0x30, 0xff}: Green,
	{0xb3, 0xc1, 0x6c, 0xff}: Green,
	{0xac, 0xa7, 0x83, 0xff}: Green,
	{0x47, 0x4c, 0x25, 0xff}: Green,
	{0xcd, 0xd0, 0x87, 0xff}: Green,
	{0x79, 0x6d, 0x41, 0xff}: Green,
	{0xf0, 0xf4, 0xc3, 0xff}: Lime,
	{0xe6, 0xee, 0x9c, 0xff}: Lime,
	{0xdc, 0xe7, 0x75, 0xff}: Lime,
	{0xd4, 0xe1, 0x57, 0xff}: Lime,
	{0xCD, 0xDC, 0x39, 0xff}: Lime,
	{0xc0, 0xca, 0x33, 0xff}: Lime,
	{0xaf, 0xb4, 0x2b, 0xff}: Lime,
	{0xee, 0xff, 0x41, 0xff}: Lime,
	{0xc6, 0xff, 0x00, 0xff}: Lime,
	{0xae, 0xea, 0x00, 0xff}: Lime,
	{0xff, 0xf9, 0xc4, 0xff}: Yellow,
	{0xff, 0xf5, 0x9d, 0xff}: Yellow,
	{0xff, 0xf1, 0x76, 0xff}: Yellow,
	{0xff, 0xee, 0x58, 0xff}: Yellow,
	{0xff, 0xff, 0x8d, 0xff}: Yellow,
	{0xff, 0xff, 0x00, 0xff}: Yellow,
	{0xff, 0xd5, 0x4f, 0xff}: Yellow,
	{0xff, 0xca, 0x28, 0xff}: Yellow,
	{0xe3, 0xce, 0x81, 0xff}: Yellow,
	{0xd1, 0xaf, 0x52, 0xff}: Yellow,
	{0xee, 0xbb, 0x2b, 0xff}: Yellow,
	{0xd3, 0xa8, 0x3a, 0xff}: Yellow,
	{0xc5, 0xa7, 0x02, 0xff}: Yellow,
	{0x9f, 0x82, 0x01, 0xff}: Yellow,
	{0xe8, 0xce, 0x03, 0xff}: Yellow,
	{0xf9, 0xa8, 0x25, 0xff}: Orange,
	{0xFF, 0x98, 0x00, 0xff}: Orange,
	{0xff, 0xa7, 0x26, 0xff}: Orange,
	{0xfb, 0x8c, 0x00, 0xff}: Orange,
	{0xf5, 0x7c, 0x00, 0xff}: Orange,
	{0xef, 0x6c, 0x00, 0xff}: Orange,
	{0xff, 0x91, 0x00, 0xff}: Orange,
	{0xff, 0x6d, 0x00, 0xff}: Orange,
	{0xfd, 0x9a, 0x31, 0xff}: Orange,
	{0x7d, 0x27, 0x04, 0xff}: Orange,
	{0xfd, 0x57, 0x1f, 0xff}: Orange,
	{0xf8, 0x67, 0x04, 0xff}: Orange,
	{0xfd, 0x9a, 0x00, 0xff}: Orange,
	{0xfe, 0x8a, 0x00, 0xff}: Orange,
	{0xf1, 0x96, 0x52, 0xff}: Orange,
	{0xe5, 0x83, 0x47, 0xff}: Orange,
	{0xc9, 0x4c, 0x30, 0xff}: Orange,
	{0x9f, 0x56, 0x01, 0xff}: Orange,
	{0xfa, 0x68, 0x01, 0xff}: Orange,
	{0xbb, 0x72, 0x3d, 0xff}: Orange,
	{0xff, 0x52, 0x52, 0xff}: Red,
	{0xf4, 0x43, 0x36, 0xff}: Red,
	{0xef, 0x53, 0x50, 0xff}: Red,
	{0xe5, 0x39, 0x35, 0xff}: Red,
	{0xf6, 0x29, 0x2e, 0xff}: Red,
	{0xfc, 0x25, 0x2d, 0xff}: Red,
	{0xd3, 0x2f, 0x2f, 0xff}: Red,
	{0xc6, 0x28, 0x28, 0xff}: Red,
	{0xba, 0x28, 0x30, 0xff}: Red,
	{0xb7, 0x1c, 0x1c, 0xff}: Red,
	{0xd5, 0x00, 0x00, 0xff}: Red,
	{0xdb, 0x08, 0x06, 0xff}: Red,
	{0xcf, 0x09, 0x04, 0xff}: Red,
	{0xd8, 0x1a, 0x14, 0xff}: Red,
	{0xcc, 0x17, 0x08, 0xff}: Red,
	{0xd8, 0x0a, 0x07, 0xff}: Red,
	{0xde, 0x26, 0x16, 0xff}: Red,
	{0xee, 0x24, 0x0f, 0xff}: Red,
	{0xa1, 0x21, 0x1f, 0xff}: Red,
	{0x70, 0x12, 0x19, 0xff}: Red,
	{0x51, 0x12, 0x18, 0xff}: Red,
	{0x49, 0x11, 0x14, 0xff}: Red,
	{0xfc, 0xe4, 0xec, 0xff}: Pink,
	{0xfd, 0xc8, 0xeb, 0xff}: Pink,
	{0xe7, 0x9f, 0xa6, 0xff}: Pink,
	{0xf8, 0xbb, 0xd0, 0xff}: Pink,
	{0xf4, 0x8f, 0xb1, 0xff}: Pink,
	{0xff, 0x80, 0xab, 0xff}: Pink,
	{0xff, 0x40, 0x81, 0xff}: Pink,
	{0xf5, 0x00, 0x57, 0xff}: Pink,
	{0xf0, 0x62, 0x92, 0xff}: Pink,
	{0xec, 0x40, 0x7a, 0xff}: Pink,
	{0xe9, 0x1e, 0x63, 0xff}: Pink,
	{0xd8, 0x1b, 0x60, 0xff}: Pink,
	{0xc2, 0x18, 0x5b, 0xff}: Pink,
	{0xff, 0x00, 0xff, 0xff}: Magenta,
	{0xe5, 0x00, 0xe5, 0xff}: Magenta,
	{0xf0, 0x00, 0xb5, 0xff}: Magenta,
	{0xce, 0x00, 0x9b, 0xff}: Magenta,
	{0xc0, 0x05, 0x5b, 0xff}: Magenta,
	{0xb0, 0x00, 0x85, 0xff}: Magenta,
	{0xa8, 0x28, 0x63, 0xff}: Magenta,
	{0x5b, 0x00, 0x2f, 0xff}: Magenta,
	{0x4b, 0x01, 0x21, 0xff}: Magenta,
	{0x86, 0x02, 0x25, 0xff}: Magenta,
	{0xcb, 0x02, 0x3d, 0xff}: Magenta,
	{0x64, 0x07, 0x1a, 0xff}: Magenta,
	{0x9e, 0x00, 0x47, 0xff}: Magenta,
	{0xdc, 0x7a, 0xcf, 0xff}: Magenta,
	{0xed, 0xde, 0xac, 0xff}: Gold,
	{0xe8, 0xb4, 0x51, 0xff}: Gold,
	{0xc0, 0x8a, 0x3e, 0xff}: Gold,
	{0xa2, 0x7d, 0x4b, 0xff}: Gold,
	{0x75, 0x55, 0x31, 0xff}: Gold,
	{0xd1, 0x93, 0x27, 0xff}: Gold,
	{0xde, 0xa2, 0x53, 0xff}: Gold,
	{0xd5, 0xaa, 0x6f, 0xff}: Gold,
	{0xf5, 0xea, 0xd4, 0xff}: Gold,
}

func Colorful(actualColor colorful.Color) (result Color) {
	var distance = 1.0

	for rgba, i := range ColorMap {
		colorColorful, _ := colorful.MakeColor(rgba)
		currentDistance := colorColorful.DistanceLab(actualColor)

		if distance >= currentDistance {
			distance = currentDistance
			result = i
		}
	}

	return result
}
