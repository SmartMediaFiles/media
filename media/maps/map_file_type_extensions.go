package maps

import (
	"path/filepath"

	"github.com/smartmediafiles/media/media/types"
)

// MapFileTypeExtensions is a map where keys are `FileType` and values are slices of `FileExtension`.
// This structure is used to group all possible file extensions under a specific file format.
//
// Example:
//
//	jpegMap := MapFileTypeExtensions{
//	    "jpeg": {".jpg", ".jpeg", ".jfif"},
//	}
type MapFileTypeExtensions map[types.FileType][]types.FileExtension

// GetFileTypes returns a list of file types.
func (m MapFileTypeExtensions) GetFileTypes() []types.FileType {
	var fileTypes []types.FileType
	for fileType := range m {
		fileTypes = append(fileTypes, fileType)
	}
	return fileTypes
}

// GetFileTypeAndExtension returns the file type and extension for the given file name.
func (m MapFileTypeExtensions) GetFileTypeAndExtension(fileName string) (types.FileType, types.FileExtension) {
	fileExt := types.FileExtension(filepath.Ext(filepath.Clean(fileName)))
	for fileType, extensions := range m {
		for _, ext := range extensions {
			if ext == fileExt {
				return fileType, ext
			}
		}
	}
	return "", ""
}

// GetFileType returns the file type for the given file extension.
func (m MapFileTypeExtensions) GetFileType(fileExtension types.FileExtension) types.FileType {
	for fileType, extensions := range m {
		for _, ext := range extensions {
			if ext == fileExtension {
				return fileType
			}
		}
	}
	return ""
}

// GetExtensions returns a slice of all `FileExtension` values for a given `FileType`.
// It returns a copy of the slice to prevent mutation of the original data.
func (m MapFileTypeExtensions) GetExtensions(fileType types.FileType) []types.FileExtension {
	if extensions, ok := m[fileType]; ok {
		// Return a copy to prevent modification of the original slice.
		result := make([]types.FileExtension, len(extensions))
		copy(result, extensions)
		return result
	}
	return nil
}

// GetExtensionsAsString returns a slice of all file extensions for a given `FileType`,
// with each extension represented as a raw string.
func (m MapFileTypeExtensions) GetExtensionsAsString(fileType types.FileType) []string {
	extensions := m.GetExtensions(fileType)
	stringExtensions := make([]string, 0, len(extensions))
	for _, ext := range extensions {
		stringExtensions = append(stringExtensions, ext.String())
	}
	return stringExtensions
}

// GetAllExtensions returns a flattened slice containing all file extensions from all file types in the map.
// The order is not guaranteed, and there may be duplicates if extensions are shared.
func (m MapFileTypeExtensions) GetAllExtensions() []types.FileExtension {
	var allExtensions []types.FileExtension
	for _, extensions := range m {
		allExtensions = append(allExtensions, extensions...)
	}
	return allExtensions
}

// GetAllExtensionsAsString returns a flattened slice of all file extensions as raw strings.
func (m MapFileTypeExtensions) GetAllExtensionsAsString() []string {
	extensions := m.GetAllExtensions()
	stringExtensions := make([]string, 0, len(extensions))
	for _, ext := range extensions {
		stringExtensions = append(stringExtensions, ext.String())
	}
	return stringExtensions
}
