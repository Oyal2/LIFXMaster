package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LIFXProducts []struct {
	Vid      int       `json:"vid"`
	Name     string    `json:"name"`
	Defaults Defaults  `json:"defaults"`
	Products []Product `json:"products"`
}

// Default values for the given products
type Defaults struct {
	Hev               bool  `json:"hev"`
	Color             bool  `json:"color"`
	Chain             bool  `json:"chain"`
	Matrix            bool  `json:"matrix"`
	Relays            bool  `json:"relays"`
	Buttons           bool  `json:"buttons"`
	Infrared          bool  `json:"infrared"`
	Multizone         bool  `json:"multizone"`
	TemperatureRange  []int `json:"temperature_range"`
	ExtendedMultizone bool  `json:"extended_multizone"`
}

// Product represents a single product with its properties.
type Product struct {
	PID      int       `json:"pid"`
	Name     string    `json:"name"`
	Features Features  `json:"features"`
	Upgrades []Upgrade `json:"upgrades"`
}

// Features represents the features of a product.
type Features struct {
	Hev                  bool  `json:"hev,omitempty"`
	Color                bool  `json:"color,omitempty"`
	Chain                bool  `json:"chain,omitempty"`
	Matrix               bool  `json:"matrix,omitempty"`
	Relays               bool  `json:"relays,omitempty"`
	Buttons              bool  `json:"buttons,omitempty"`
	Infrared             bool  `json:"infrared,omitempty"`
	Multizone            bool  `json:"multizone"`
	TemperatureRange     []int `json:"temperature_range,omitempty"`
	ExtendedMultizone    bool  `json:"extended_multizone,omitempty"`
	MinExtMZFirmware     *int  `json:"min_ext_mz_firmware,omitempty"`            // Optional field
	MinExtMZFirmwareComp []int `json:"min_ext_mz_firmware_components,omitempty"` // Optional field
}

// Upgrade represents an upgrade available for a product.
type Upgrade struct {
	Major    int      `json:"major"`
	Minor    int      `json:"minor"`
	Features Features `json:"features"`
}

// FetchProducts makes an HTTP GET request to the specified URL and decodes the JSON response.
func FetchProducts() (map[int]*Product, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/LIFX/products/master/products.json")
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var contentType = resp.Header.Get("Content-Type")
	if contentType != "application/json" && contentType != "text/plain; charset=utf-8" {
		return nil, fmt.Errorf("unexpected content type: %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var lifxProducts LIFXProducts
	err = json.Unmarshal(body, &lifxProducts)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	products := make(map[int]*Product)
	for _, lifx := range lifxProducts {
		for _, product := range lifx.Products {
			p := product
			products[product.PID] = &p
		}
	}

	return products, nil
}
