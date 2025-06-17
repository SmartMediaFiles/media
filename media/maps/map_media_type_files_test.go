package maps

import (
	"testing"

	"github.com/smartmediafiles/media/media/types"
	"github.com/stretchr/testify/assert"
)

var fileTypeJpeg = types.FileType("jpeg")
var fileTypePng = types.FileType("png")
var fileTypeMp4 = types.FileType("mp4")

var fileExtensionJpg = types.FileExtension("jpg")
var fileExtensionJpeg = types.FileExtension("jpeg")
var fileExtensionPng = types.FileExtension("png")
var fileExtensionMp4 = types.FileExtension("mp4")

var mediaTypeImage = types.MediaType("image")
var mediaTypeVideo = types.MediaType("video")

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

	t.Run("GetMedia", func(t *testing.T) {
		result := testMap.GetMedia(mediaTypeImage)
		assert.Equal(t, 2, len(result), "Should have 2 file types for MediaTypeImage")
		assert.Contains(t, result, fileTypeJpeg, "Should contain FileTypeJPEG")
		assert.Contains(t, result, fileTypePng, "Should contain FileTypePNG")
	})

	t.Run("GetMediaTypes", func(t *testing.T) {
		result := testMap.GetMediaTypes()
		assert.Equal(t, 2, len(result), "Should return 2 media types")
		assert.Contains(t, result, mediaTypeImage, "Should contain MediaTypeImage")
		assert.Contains(t, result, mediaTypeVideo, "Should contain MediaTypeVideo")
	})

	t.Run("GetFileTypes", func(t *testing.T) {
		result := testMap.GetFileTypes(mediaTypeImage)
		assert.Equal(t, 2, len(result), "Should return 2 file types for MediaTypeImage")
		assert.Contains(t, result, fileTypeJpeg, "Should contain FileTypeJPEG")
		assert.Contains(t, result, fileTypePng, "Should contain FileTypePNG")
	})

	t.Run("GetFileExtensions", func(t *testing.T) {
		result := testMap.GetExtensions(mediaTypeImage, fileTypeJpeg)
		assert.Equal(t, 2, len(result), "Should return 2 extensions for JPEG")
		assert.Contains(t, result, fileExtensionJpg, "Should contain FileExtensionJPG")
		assert.Contains(t, result, fileExtensionJpeg, "Should contain FileExtensionJPEG")
	})
}
