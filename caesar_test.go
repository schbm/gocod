package gocod

import "testing"

func TestCaesarEncode(t *testing.T) {
	message := "Hallo zämme i biz; .- bla 1234 0-9 0-10 üöqwevyx"
	t.Log(message)
	result := CaesarEncode(message, 1)
	t.Log(result)
	reversed := CaesarDecode(result, 1)
	t.Log(reversed)

	if message != reversed {
		t.Error("strings do not match")
	}
}
