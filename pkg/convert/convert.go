package convert

// offsetTranslate takes an input string and list of offset range definitions
// to translate from one set of characters to another. Offset definitions are
// made of a starting character, ending character, and replacement for the
// starting character.
func offsetTranslate(text string, offsets [][]rune) string {
	out := ""
	for _, r := range text {
		found := false

		for _, definition := range offsets {
			start := definition[0]
			end := definition[1]
			offset := definition[2] - start

			if r >= start && r <= end {
				out += string(r + offset)
				found = true
				break
			}
		}

		if !found {
			out += string(r)
		}
	}

	return out
}


// Cursive attempts to convert alphanumeric characters to their cursive
// equivalents. Characters without a cursive equivalent are left untouched.
func Cursive(text string) string {
	// There is no single group of these unfortunately, so we need a separate
	// set for each non-contiguous grouping.
	offsets := [][]rune{
		{'A', 'A', '𝒜'},
		{'B', 'B', 'ℬ'},
		{'C', 'D', '𝒞'},
		{'E', 'F', 'ℰ'},
		{'G', 'G', '𝒢'},
		{'H', 'H', 'ℋ'},
		{'I', 'I', 'ℐ'},
		{'J', 'K', '𝒥'},
		{'L', 'L', 'ℒ'},
		{'M', 'M', 'ℳ'},
		{'N', 'Q', '𝒩'},
		{'R', 'R', 'ℛ'},
		{'S', 'Z', '𝒮'},
		{'a', 'd', '𝒶'},
		{'e', 'e', 'ℯ'},
		{'f', 'f', '𝒻'},
		{'g', 'g', 'ℊ'},
		{'h', 'n', '𝒽'},
		{'o', 'o', 'ℴ'},
		{'p', 'z', '𝓅'},
	}

	return offsetTranslate(text, offsets)
}
