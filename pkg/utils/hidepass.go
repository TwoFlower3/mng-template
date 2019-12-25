package utils

const (
	hideSymbol rune = '*'
)

// HidePass replace symbols
func HidePass(pass string) string {
	if len(pass) == 1 {
		return "*"
	}

	r := []rune(pass)
	for i := 0; i < len(r); i++ {
		if k := i % 5; k != 0 {
			r[i] = hideSymbol
		}
	}
	return string(r)
}
