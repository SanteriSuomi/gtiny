package utils

// IfValueElse is a utility function that returns one of two strings based on the
// boolean value of the first argument. If the value is true, it returns the
// second argument (ifTrue). If the value is false, it returns the third argument
// (ifFalse).
//
// Parameters:
// - value: A boolean value indicating the condition to check.
// - ifTrue: A string to return if the condition is true.
// - ifFalse: A string to return if the condition is false.
//
// Returns:
//   - A string that is either ifTrue or ifFalse, depending on the value of the
//     condition.
func IfValueElse(value bool, ifTrue string, ifFalse string) string {
	if value {
		return ifTrue
	}
	return ifFalse
}
