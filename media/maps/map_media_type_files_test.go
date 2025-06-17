package maps

import (
	"testing"

	"github.com/smartmediafiles/media/media/types"
	"github.com/stretchr/testify/assert"
)

func TestMapMediaTypeFiles(t *testing.T) {
	// Création d'un MapMediaTypeFiles de test
	testMap := MapMediaTypeFiles{
		mediaTypeImage: {
			fileTypeJpeg: {fileExtensionJpg, fileExtensionJpeg},
			fileTypePng:  {fileExtensionPng},
		},
		mediaTypeVideo: {
			fileTypeMp4: {fileExtensionMp4},
		},
	}

	// TestGetMedia verifies that the method returns the correct MapFileTypeExtensions for a given media type,
	// and returns nil for a non-existent media type.
	t.Run("GetMedia", func(t *testing.T) {
		expected := MapFileTypeExtensions{
			fileTypeJpeg: {fileExtensionJpg, fileExtensionJpeg},
			fileTypePng:  {fileExtensionPng},
		}
		result := testMap.GetFileTypeExtensions(mediaTypeImage)
		assert.Equal(t, expected, result, "Should return correct map for media type")

		resultNil := testMap.GetFileTypeExtensions(mediaTypeFake)
		assert.Nil(t, resultNil, "Should return nil for non-existent media type")
	})

	// TestGetMediaTypes ensures that all media type keys are correctly returned.
	t.Run("GetMediaTypes", func(t *testing.T) {
		expected := []types.MediaType{mediaTypeImage, mediaTypeVideo}
		result := testMap.GetMediaTypes()
		assert.ElementsMatch(t, expected, result, "Should return all media types")
	})

	// TestGetFileTypes checks that the method returns the correct file types for a given media type.
	t.Run("GetFileTypes", func(t *testing.T) {
		tests := []struct {
			name      string
			mediaType types.MediaType
			want      []types.FileType
		}{
			{
				name:      "Image file types",
				mediaType: mediaTypeImage,
				want:      []types.FileType{fileTypeJpeg, fileTypePng},
			},
			{
				name:      "Video file types",
				mediaType: mediaTypeVideo,
				want:      []types.FileType{fileTypeMp4},
			},
			{
				name:      "Fake media type",
				mediaType: mediaTypeFake,
				want:      []types.FileType{},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := testMap.GetFileTypes(tt.mediaType)
				assert.ElementsMatch(t, tt.want, result, "Should return correct file types for media type")
			})
		}
	})

	// TestGetExtensions verifies that the correct extensions are returned for a specific media and file type combination.
	t.Run("GetExtensions", func(t *testing.T) {
		tests := []struct {
			name      string
			mediaType types.MediaType
			fileType  types.FileType
			want      []types.FileExtension
		}{
			{
				name:      "Image JPEG extensions",
				mediaType: mediaTypeImage,
				fileType:  fileTypeJpeg,
				want:      []types.FileExtension{fileExtensionJpg, fileExtensionJpeg},
			},
			{
				name:      "Image PNG extensions",
				mediaType: mediaTypeImage,
				fileType:  fileTypePng,
				want:      []types.FileExtension{fileExtensionPng},
			},
			{
				name:      "Video MP4 extensions",
				mediaType: mediaTypeVideo,
				fileType:  fileTypeMp4,
				want:      []types.FileExtension{fileExtensionMp4},
			},
			{
				name:      "Fake media type",
				mediaType: mediaTypeFake,
				fileType:  fileTypeJpeg,
				want:      nil,
			},
			{
				name:      "Image with fake file type",
				mediaType: mediaTypeImage,
				fileType:  types.FileType("fake"),
				want:      nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := testMap.GetExtensions(tt.mediaType, tt.fileType)
				assert.Equal(t, tt.want, result, "Should return correct extensions")
			})
		}
	})

	// TestGetMediaTypeByExtension verifies that GetMediaTypeByExtension correctly finds the media type for an extension.
	t.Run("GetMediaTypeByExtension", func(t *testing.T) {
		tests := []struct {
			name      string
			extension types.FileExtension
			want      types.MediaType
		}{
			{
				name:      "find jpg",
				extension: fileExtensionJpg,
				want:      mediaTypeImage,
			},
			{
				name:      "find jpeg",
				extension: fileExtensionJpeg,
				want:      mediaTypeImage,
			},
			{
				name:      "find png",
				extension: fileExtensionPng,
				want:      mediaTypeImage,
			},
			{
				name:      "find mp4",
				extension: fileExtensionMp4,
				want:      mediaTypeVideo,
			},
			{
				name:      "find case-insensitive JPG",
				extension: types.FileExtension(".JPG"),
				want:      mediaTypeImage,
			},
			{
				name:      "find case-insensitive JPEG",
				extension: types.FileExtension(".jPeG"),
				want:      mediaTypeImage,
			},
			{
				name:      "not found",
				extension: types.FileExtension(".gif"),
				want:      "",
			},
			{
				name:      "empty extension",
				extension: types.FileExtension(""),
				want:      "",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := testMap.GetMediaTypeByExtension(tt.extension)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
