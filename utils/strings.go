package utils

func Pop(x string)(string, string) {
	return string(x[0]), x[1:]
}