package piscine

func Compact(ptr *[]string) int {
	first_slice := *ptr
	args := make([]string, 0)
	for _, q := range first_slice {
		if q != "" {
			args = append(args, q)
		}
	}
	*ptr = args
	return len(args)
}
