package maps

import "github.com/smartmediafiles/media/media/types"

// MapFileTypeExtensions is a map of file types to file extensions.
type MapFileTypeExtensions map[types.FileType][]types.FileExtension

// GetFileTypes returns a list of file types.
func (m MapFileTypeExtensions) GetFileTypes() []types.FileType {
	var fileTypes []types.FileType
	for fileType := range m {
		fileTypes = append(fileTypes, fileType)
	}
	return fileTypes
}

// GetExtensions returns a list of file extensions for the given file type.
func (m MapFileTypeExtensions) GetExtensions(fileType types.FileType) []types.FileExtension {
	return m[fileType]
}

// GetAllExtensions returns a list of all file extensions.
func (m MapFileTypeExtensions) GetAllExtensions() []types.FileExtension {
	var extensions []types.FileExtension
	for _, exts := range m {
		extensions = append(extensions, exts...)
	}
	return extensions
}
