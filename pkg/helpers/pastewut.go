package helpers

import (
  "math/rand"
  "time"
)


func GenerateCode() string {
  var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
  const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := make([]byte, 10)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(code)
}
