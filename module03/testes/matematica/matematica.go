package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsavel por calcular o que você já sabe ... :)
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}

	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
