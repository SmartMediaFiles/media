package types

// MediaType represents a general media content type.
type MediaType string

// String returns the type as string.
func (t MediaType) String() string {
	return string(t)
}

// Equal checks if the type matches.
func (t MediaType) Equal(s string) bool {
	return s == t.String()
}

// NotEqual checks if the type is different.
func (t MediaType) NotEqual(s string) bool {
	return !t.Equal(s)
}
