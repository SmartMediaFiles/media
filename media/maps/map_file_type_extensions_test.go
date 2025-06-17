package maps

import (
	"testing"

	"github.com/smartmediafiles/media/media/types"
	"github.com/stretchr/testify/assert"
)

func TestMapFileTypeExtensions(t *testing.T) {
	m := MapFileTypeExtensions{
		fileTypeImage: {fileExtensionJpg, fileExtensionPng},
		fileTypeVideo: {fileExtensionMp4, fileExtensionAvi},
	}

	// TestGetFileTypes ensures the method returns all configured file types.
	t.Run("GetFileTypes", func(t *testing.T) {
		expected := []types.FileType{fileTypeImage, fileTypeVideo}
		result := m.GetFileTypes()
		assert.ElementsMatch(t, expected, result, "GetFileTypes() should return all file types")
	})

	// TestGetFileTypeAndExtension verifies that the correct file type and extension
	// are extracted from a given file name or path.
	t.Run("GetFileTypeAndExtension", func(t *testing.T) {
		tests := []struct {
			name          string
			fileName      string
			wantFileType  types.FileType
			wantExtension types.FileExtension
		}{
			{"JPG file", "photo.jpg", fileTypeImage, fileExtensionJpg},
			{"PNG file with path", "/path/to/image.png", fileTypeImage, fileExtensionPng},
			{"AVI file", "movie.avi", fileTypeVideo, fileExtensionAvi},
			{"Unknown extension", "document.txt", types.FileType(""), types.FileExtension("")},
			{"No extension", "file", types.FileType(""), types.FileExtension("")},
			{"Empty filename", "", types.FileType(""), types.FileExtension("")},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotFileType, gotExtension := m.GetFileTypeAndExtension(tt.fileName)
				assert.Equal(t, tt.wantFileType, gotFileType, "File type does not match")
				assert.Equal(t, tt.wantExtension, gotExtension, "File extension does not match")
			})
		}
	})

	// TestGetFileType checks that the correct file type is returned for a given extension.
	t.Run("GetFileType", func(t *testing.T) {
		tests := []struct {
			name          string
			fileExtension types.FileExtension
			want          types.FileType
		}{
			{"JPG extension", fileExtensionJpg, fileTypeImage},
			{"PNG extension", fileExtensionPng, fileTypeImage},
			{"AVI extension", fileExtensionAvi, fileTypeVideo},
			{"Unknown extension", ".txt", ""},
			{"Empty extension", "", ""},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := m.GetFileType(tt.fileExtension)
				assert.Equal(t, tt.want, got, "GetFileType() did not return expected file type")
			})
		}
	})

	// TestGetExtensions verifies that the method returns the correct slice of extensions for a specific file type.
	t.Run("GetExtensions", func(t *testing.T) {
		tests := []struct {
			name     string
			fileType types.FileType
			want     []types.FileExtension
		}{
			{"Image extensions", fileTypeImage, []types.FileExtension{fileExtensionJpg, fileExtensionPng}},
			{"Video extensions", fileTypeVideo, []types.FileExtension{fileExtensionMp4, fileExtensionAvi}},
			{"Non-existent type", fileTypeFake, nil},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := m.GetExtensions(tt.fileType)
				// Check for deep equality for slices.
				if !assert.ObjectsAreEqual(tt.want, got) {
					t.Errorf("GetExtensions() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// TestGetExtensionsAsString ensures that extensions are correctly returned as a slice of strings.
	t.Run("GetExtensionsAsString", func(t *testing.T) {
		tests := []struct {
			name     string
			fileType types.FileType
			want     []string
		}{
			{"Image extensions", fileTypeImage, []string{".jpg", ".png"}},
			{"Video extensions", fileTypeVideo, []string{".mp4", ".avi"}},
			{"Non-existent type", fileTypeFake, []string{}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := m.GetExtensionsAsString(tt.fileType)
				assert.ElementsMatch(t, tt.want, got, "GetExtensionsAsString() did not return expected extensions")
			})
		}
	})

	// TestGetAllExtensions checks that the method returns a flattened slice of all extensions from all file types.
	t.Run("GetAllExtensions", func(t *testing.T) {
		expected := []types.FileExtension{fileExtensionJpg, fileExtensionPng, fileExtensionMp4, fileExtensionAvi}
		result := m.GetAllExtensions()
		assert.ElementsMatch(t, expected, result, "GetAllExtensions() should return all extensions from all types")
	})

	// TestGetAllExtensionsAsString verifies that all extensions from all types are correctly returned as a flattened string slice.
	t.Run("GetAllExtensionsAsString", func(t *testing.T) {
		expected := []string{".jpg", ".png", ".mp4", ".avi"}
		result := m.GetAllExtensionsAsString()
		assert.ElementsMatch(t, expected, result, "GetAllExtensionsAsString() should return all extensions as strings")
	})
}
