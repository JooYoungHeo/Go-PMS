package archavon

var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false

	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}

	return result
}

func ModifyHangulPost(s string) string {
	post := "는"

	if HasConsonantSuffix(s) {
		post = "은"
	}

	return post
}
