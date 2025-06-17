// Package maps defines the structure of the data mappings used to associate
// media types, file types, and file extensions within the SmartMediaFiles ecosystem.
package maps

import (
	"github.com/smartmediafiles/media/media/types"
)

// MapMediaTypeFiles is a map where keys are `MediaType` and values are `MapFileTypeExtensions`.
// It serves as the top-level structure for the central media registry.
//
// Example:
//
//	registry := MapMediaTypeFiles{
//	    "image": MapFileTypeExtensions{...},
//	    "video": MapFileTypeExtensions{...},
//	}
type MapMediaTypeFiles map[types.MediaType]MapFileTypeExtensions

// GetMedia returns the complete map of file types to extensions for a given media type.
func (m MapMediaTypeFiles) GetMedia(mediaType types.MediaType) MapFileTypeExtensions {
	return m[mediaType]
}

// GetMediaTypes returns a slice containing all the `MediaType` keys from the map.
// The order of the returned types is not guaranteed.
func (m MapMediaTypeFiles) GetMediaTypes() []types.MediaType {
	mediaTypes := make([]types.MediaType, 0, len(m))
	for mediaType := range m {
		mediaTypes = append(mediaTypes, mediaType)
	}
	return mediaTypes
}

// GetFileTypes returns a slice of all `FileType` keys for a given `MediaType`.
// The order of the returned types is not guaranteed.
func (m MapMediaTypeFiles) GetFileTypes(mediaType types.MediaType) []types.FileType {
	mediaMap := m[mediaType]
	fileTypes := make([]types.FileType, 0, len(mediaMap))
	for fileType := range mediaMap {
		fileTypes = append(fileTypes, fileType)
	}
	return fileTypes
}

// GetExtensions returns a slice of all `FileExtension` values for a given `MediaType` and `FileType`.
func (m MapMediaTypeFiles) GetExtensions(mediaType types.MediaType, fileType types.FileType) []types.FileExtension {
	if media, ok := m[mediaType]; ok {
		if extensions, ok := media[fileType]; ok {
			// Return a copy to prevent modification of the original slice.
			result := make([]types.FileExtension, len(extensions))
			copy(result, extensions)
			return result
		}
	}
	return nil
}
