package models

import (
	"log"
	"math"
	"strconv"
)

func (m *Macros) GetMacro(keyword string) float32 {
	value, err := strconv.ParseFloat(m.Request.Form.Get(keyword), m.BitSize)
	if err != nil {
		log.Println(err)
		return 0
	}

	ratio := math.Pow(10, float64(m.Precision))
	value = math.Round(value*ratio) / ratio

	return float32(value)
}
