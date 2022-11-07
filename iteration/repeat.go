package iteration

func Repeat(val string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += val
	}

	return repeated
}
