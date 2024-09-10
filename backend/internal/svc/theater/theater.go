package theater

import (
	"math"

	"github.com/EdlinOrg/prominentcolor"
	"github.com/hisamafahri/coco"
	"github.com/kbinani/screenshot"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

func GenerateDominantColorPalette(k, screen int) ([]structure.HSBK, error) {
	bounds := screenshot.GetDisplayBounds(screen)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}

	// Use the KmeansWithAll function from the prominentcolor package
	colorItems, err := prominentcolor.KmeansWithAll(
		k,
		img,
		prominentcolor.ArgumentAverageMean,
		600,
		[]prominentcolor.ColorBackgroundMask{},
	)

	if err != nil {
		return nil, err
	}

	// Convert ColorItems to HSB
	hsbColors := make([]structure.HSBK, len(colorItems))
	for itemIdx, item := range colorItems {
		output := coco.Rgb2Hsl(float64(item.Color.R), float64(item.Color.G), float64(item.Color.B))
		hsbColors[itemIdx] = structure.HSBK{
			Hue:        uint16(math.Round(65535.0 * output[0] / 360.0)),
			Saturation: uint16(math.Round(0xFFFF * float64(output[1]) / 100)),
			Brightness: uint16(math.Round(0xFFFF * float64(output[2]) / 100)),
			Kelvin:     3500,
		}
	}

	return hsbColors, nil
}

// RGBToHSBK converts RGB values (uint32) to HSBK
func RGBToHSBK(rr, gg, bb uint32) structure.HSBK {
	// Convert uint32 to float64 and normalize to 0-1 range
	r := float64(rr) / 65535.0
	g := float64(gg) / 65535.0
	b := float64(bb) / 65535.0

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	delta := max - min

	var h, s float64
	// v := max

	if delta == 0 {
		h = 0
		s = 0
	} else {
		s = delta / max
		switch max {
		case r:
			h = (g - b) / delta
			if g < b {
				h += 6
			}
		case g:
			h = 2 + (b-r)/delta
		case b:
			h = 4 + (r-g)/delta
		}
		h *= 60
	}

	// Convert H, S, B to uint16 range (0-65535)
	hue := uint16(math.Round(65535.0 * h / 360.0))
	saturation := uint16(math.Round(65535.0 * s))
	brightness := uint16(math.Round(65535.0 * 1))

	// For Kelvin, we'll use a default value of 3500K
	// You may want to adjust this based on your specific requirements
	kelvin := uint16(3500)

	return structure.HSBK{
		Hue:        hue,
		Saturation: saturation,
		Brightness: brightness,
		Kelvin:     kelvin,
	}
}
