package go_koans

func aboutEnumeration() {
	{
		var concatenated string
		var total int

		strings := []string{"hello", " world", "!"}
		for i, v := range strings {
			total += i
			concatenated += v
		}

		assert(concatenated == "hello world!") // for loops have a modern variation
		assert(total == 3)                     // which offers both a value and an index
	}

	{
		var totalLength int

		strings := []string{"hello", " world", "!"}
		for _, v := range strings {
			totalLength += len(v)
		}

		assert(totalLength == len(strings[0])+len(strings[1])+len(strings[2])) // although we may omit either value
	}
}
