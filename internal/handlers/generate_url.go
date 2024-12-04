package handlers

import (
	"math/rand"
	"time"
)

func GerarUrlAleatoria() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 5 + r.Intn(6)

	result := make([]byte, length)
	for i := range result {
		switch r.Intn(3) {
		case 0:
			result[i] = byte(r.Intn(10) + '0')
		case 1:
			result[i] = byte(r.Intn(26) + 'A')
		case 2:
			result[i] = byte(r.Intn(26) + 'a')
		}
	}

	return string(result)
}
