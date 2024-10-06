package types

// FileType represents a media file type.
type FileType string

// String returns the file type as string.
func (t FileType) String() string {
	return string(t)
}

// Equal checks if the file type matches.
func (t FileType) Equal(s string) bool {
	return t.String() == s
}

// NotEqual checks if the file type is different.
func (t FileType) NotEqual(s string) bool {
	return !t.Equal(s)
}
