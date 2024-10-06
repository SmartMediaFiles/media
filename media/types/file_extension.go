package types

// FileExtension is a string representing the file extension.
type FileExtension string

// String returns the string representation of the file extension.
func (e FileExtension) String() string {
	return string(e)
}

// Equals checks if the file extension is equal to the given file extension.
func (e FileExtension) Equals(s string) bool {
	return e.String() == s
}

// NotEquals checks if the file extension is not equal to the given file extension.
func (e FileExtension) NotEquals(s string) bool {
	return !e.Equals(s)
}
