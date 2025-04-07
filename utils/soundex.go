package utils

import "strings"

func Soundex(s string) string {
	s = strings.ToUpper(s)

	if len(s) == 0 {
		return ""
	}

	soundexMap := map[rune]string{
		'B': "1", 'F': "1", 'P': "1", 'V': "1",
		'C': "2", 'G': "2", 'J': "2", 'K': "2", 'Q': "2", 'S': "2", 'X': "2", 'Z': "2",
		'D': "3", 'T': "3",
		'L': "4",
		'M': "5", 'N': "5",
		'R': "6",
	}
	var result strings.Builder
	firstLetter := rune(s[0])
	result.WriteRune(firstLetter)

	prevCode := ""

	for _, r := range s[1:] {
		code, ok := soundexMap[r]
		if !ok {
			code = ""
		}
		if code != "" && code != prevCode {
			result.WriteString(code)
		}
		if code != "" {
			prevCode = code
		}
	}

	for result.Len() < 4 {
		result.WriteString("0")
	}

	return result.String()[:4]
}
