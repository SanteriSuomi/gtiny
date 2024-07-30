package utils

func IfElse(value bool, ifTrue string, ifFalse string) string {
	if value {
		return ifTrue
	}
	return ifFalse
}
