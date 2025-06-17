package types

// FileExtension represents a file extension, including the leading dot (e.g., ".jpg").
// It is the primary means of identifying a file's format.
type FileExtension string

// String returns the file extension as a string.
func (e FileExtension) String() string {
	return string(e)
}

// Equal checks if the file extension matches a given string.
// This provides a case-sensitive comparison.
func (e FileExtension) Equal(s string) bool {
	return s == e.String()
}

// NotEqual checks if the file extension is different from a given string.
func (e FileExtension) NotEqual(s string) bool {
	return !e.Equal(s)
}
