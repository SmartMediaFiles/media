package types

// FileType represents a specific file format within a MediaType, like "jpeg" or "png".
// It provides a more granular classification than MediaType.
type FileType string

// String returns the file type as a string.
func (t FileType) String() string {
	return string(t)
}

// Equal checks if the file type matches a given string.
// This provides a case-sensitive comparison.
func (t FileType) Equal(s string) bool {
	return s == t.String()
}

// NotEqual checks if the file type is different from a given string.
func (t FileType) NotEqual(s string) bool {
	return !t.Equal(s)
}
