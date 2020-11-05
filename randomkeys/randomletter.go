package randomletter

import (
	"math/rand"
	"time"
)

func main() {

}

// Generate Random letter with n letters
func Split(n int) []string {

	var anzahlkeys = 5

	var keys []string

	var wort string = RandStringRunes(n * anzahlkeys)

	if anzahlkeys*n <= len(wort) {

		for i := 0; i < anzahlkeys; i++ {

			keys = append(keys, wort[n*i:n*(i+1)])

		}

	}

	return keys

}

func RandStringRunes(n int) string {

	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)

	for i := range b {

		b[i] = letterRunes[rand.Intn(len(letterRunes))]

	}

	return string(b)

}
