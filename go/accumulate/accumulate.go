// Package accumulate takes a collection and operation and returns the result
package accumulate

// Accumulate takes an operation name and a string collection and returns the result
func Accumulate(collection []string, operation string) string {
	var result []string
	if operation == "echo" {
		result = collection
	}
	if operation == "capitalize" {
		result = capitalize(collection)
	}
	if operation == "strings.ToUpper" {
		result = uppercase(collection)
	}
	return result
}

func capitalize(collection []string) []string {
	return "placeholder"
}

func uppercase(collection []string) []string {
	return "placeholder"
}
