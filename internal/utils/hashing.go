package utils

func ToBase62(original string) string {
	base62 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	originalBytes := []byte(original)
	originalLength := len(originalBytes)
	result := make([]byte, 0, originalLength)

	for i := 0; i < originalLength; i++ {
		result = append(result, base62[originalBytes[i]%62])
	}

	return string(result)
}
