package maps

import "github.com/smartmediafiles/media/media/types"

// MapMediaTypeFiles is a map of media types to their file types and extensions.
type MapMediaTypeFiles map[types.MediaType]MapFileTypeExtensions

// GetMedia returns the media type for the given file type.
func (m MapMediaTypeFiles) GetMedia(mediaType types.MediaType) MapFileTypeExtensions {
	return m[mediaType]
}

// GetMediaTypes returns a list of all media types.
func (m MapMediaTypeFiles) GetMediaTypes() []types.MediaType {
	var mediaTypes []types.MediaType
	for mediaType := range m {
		mediaTypes = append(mediaTypes, mediaType)
	}
	return mediaTypes
}

// GetFileTypes returns a list of all file types for the given media type.
func (m MapMediaTypeFiles) GetFileTypes(mediaType types.MediaType) []types.FileType {
	var fileTypes []types.FileType
	for fileType := range m[mediaType] {
		fileTypes = append(fileTypes, fileType)
	}
	return fileTypes
}

// GetExtensions returns a list of all file extensions for the given media type and file type.
func (m MapMediaTypeFiles) GetExtensions(mediaType types.MediaType, fileType types.FileType) []types.FileExtension {
	var extensions []types.FileExtension
	extensions = append(extensions, m[mediaType][fileType]...)
	return extensions
}
