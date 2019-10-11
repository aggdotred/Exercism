// Package accumulate takes a collection and operation and returns the result
package accumulate

// Accumulate takes an operation name and a string collection and returns the result
func Accumulate(collection []string, operation string) []string {
	result := []string{"place", "holder"}

	if operation == "echo" {
		result = collection
	}
	if operation == "capitalize" {
		result = firstCaps(collection)
	}
	if operation == "strings.ToUpper" {
		result = uppercase(collection)
	}
	return result
}

func firstCaps(collection []string) []string {
	return []string{"place", "holder"}
}

func uppercase(collection []string) []string {
	return []string{"place", "holder"}
}
