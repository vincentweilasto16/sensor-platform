package generator

import (
	"math"
)

// generate random DeviceCode (e.g., "A", "B", ...)
func setSensorRandomDeviceCode(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = byte('A' + rnd.Intn(26))
	}
	return string(b)
}

// generate random sensor value based on type
func setRandomSensorValueByType(t string) float64 {
	switch t {
	case "temperature":
		return roundTo(rnd.Float64()*50.0-10.0, 1) // -10.0–40.0 °C
	case "humidity":
		return roundTo(rnd.Float64()*100.0, 1) // 0–100 % RH
	case "pressure":
		return roundTo(950.0+rnd.Float64()*100.0, 1) // 950–1050 hPa
	case "co2":
		return roundTo(400.0+rnd.Float64()*1600.0, 0) // 400–2000 ppm
	case "light":
		return roundTo(rnd.Float64()*1000.0, 0) // 0–1000 lux
	default:
		return roundTo(rnd.Float64()*100.0, 1)
	}
}

// round float64 to n decimals
func roundTo(v float64, decimals int) float64 {
	p := math.Pow(10, float64(decimals))
	return math.Round(v*p) / p
}
