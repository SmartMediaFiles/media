// Package types defines the fundamental data types used across the SmartMediaFiles ecosystem.
// These types provide a consistent and type-safe way to represent the core concepts
// of media classification, such as media categories, file formats, and extensions.
package types

// MediaType represents a general category of media content, such as "image", "video", or "raw".
// It is used as the primary key for classifying and grouping different file formats.
type MediaType string

// String returns the media type as a string.
func (t MediaType) String() string {
	return string(t)
}

// Equal checks if the media type matches a given string.
// This provides a case-sensitive comparison.
func (t MediaType) Equal(s string) bool {
	return s == t.String()
}

// NotEqual checks if the media type is different from a given string.
func (t MediaType) NotEqual(s string) bool {
	return !t.Equal(s)
}
