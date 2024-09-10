package structure

type HSBK struct {
	Hue        uint16 `json:"hue"`        // Hue component of the color
	Saturation uint16 `json:"saturation"` // Saturation component of the color
	Brightness uint16 `json:"brightness"` // Brightness component of the color
	Kelvin     uint16 `json:"kelvin"`     // Kelvin component of the color temperature
}
