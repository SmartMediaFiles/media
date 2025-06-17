package media

import (
	"testing"

	media_image "github.com/smartmediafiles/media.image"
	"github.com/smartmediafiles/media/media/types"
	"github.com/stretchr/testify/assert"
)

// TestMediaRegistry conducts a suite of tests on the global Medias registry
// to ensure its integrity and the correctness of its helper functions.
func TestMediaRegistry(t *testing.T) {

	// TestMediaTypesLoaded ensures that all primary media types are present in the registry.
	t.Run("MediaTypesLoaded", func(t *testing.T) {
		expectedMediaTypes := []string{
			"raw",
			"image",
			"video",
			"vector",
			"sidecar",
			"", // for Unknown
		}
		actualMediaTypes := Medias.GetMediaTypes()

		for _, mt := range expectedMediaTypes {
			assert.Contains(t, actualMediaTypes, types.MediaType(mt), "Expected media type %s not found", mt)
		}
		assert.GreaterOrEqual(t, len(actualMediaTypes), len(expectedMediaTypes), "There should be at least all expected media types")
	})

	// TestGetMediaForImage checks that retrieving the 'image' media type returns a valid and non-empty map.
	t.Run("GetMediaForImage", func(t *testing.T) {
		images := Medias.GetFileTypeExtensions(Image)
		assert.NotNil(t, images, "Image media map should not be nil")
		assert.NotEmpty(t, images, "Image media map should not be empty")

		// Check for a known file type within images
		jpegExtensions := images.GetExtensions(media_image.ImageJpeg)
		assert.NotNil(t, jpegExtensions, "JPEG extensions should not be nil")
		assert.NotEmpty(t, jpegExtensions, "JPEG extensions should not be empty")
		assert.Contains(t, jpegExtensions, types.FileExtension(".jpg"), "Should contain .jpg extension")
	})

	// GetAllExtensionsNotEmpty verifies that the image media type in the registry contains at least one extension.
	t.Run("GetAllExtensionsNotEmpty", func(t *testing.T) {
		allImageExtensions := Medias.GetFileTypeExtensions(Image).GetAllExtensions()
		assert.NotEmpty(t, allImageExtensions, "Expected at least one extension in the image media type")
	})

	// TestGetMediaTypeByExtension verifies that the function correctly identifies the media type for various extensions.
	t.Run("GetMediaTypeByExtension", func(t *testing.T) {
		tests := []struct {
			name         string
			extension    string
			expectedType types.MediaType
		}{
			{"JPG Extension", ".jpg", Image},
			{"Case-insensitive JPG", ".JPG", Image},
			{"PNG Extension", ".png", Image},
			{"MP4 Extension", ".mp4", Video},
			{"MOV Extension", ".mov", Video},
			{"CR2 Raw Extension", ".cr2", Raw},
			{"SVG Vector Extension", ".svg", Vector},
			{"XMP Sidecar Extension", ".xmp", Sidecar},
			{"Unknown Extension", ".txt", Sidecar},
			{"Extension without dot", "jpeg", Image},
			{"Empty Extension", "", Unknown},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Add a dot if missing, for extensions like "jpeg"
				extWithDot := tt.extension
				if len(extWithDot) > 1 && extWithDot[0] != '.' {
					extWithDot = "." + extWithDot
				}
				mediaType := GetMediaTypeByExtension(types.FileExtension(extWithDot))
				assert.Equal(t, tt.expectedType, mediaType, "Mismatch in media type for extension %s", tt.extension)
			})
		}
	})
}

// Test_Media performs basic checks on the Medias registry to ensure it is not empty
// and that its core methods for retrieving media types, file types, and extensions return data.
func Test_Media(t *testing.T) {

	// Retrieve all registered media types.
	mediaTypes := Medias.GetMediaTypes()

	// Ensure that at least one media type is registered.
	if len(mediaTypes) == 0 {
		t.Errorf("Expected mediaTypes to have at least one element")
	}

	// Retrieve the file types for the first registered media type.
	fileTypes := Medias.GetFileTypes(mediaTypes[0])

	// Ensure that the media type has at least one file type associated with it.
	if len(fileTypes) == 0 {
		t.Errorf("Expected fileTypes to have at least one element")
	}

	// Retrieve extensions for the first file type of the first media type.
	exts := Medias.GetExtensions(mediaTypes[0], fileTypes[0])

	// Ensure that the file type has at least one extension.
	if len(exts) == 0 {
		t.Errorf("Expected exts to have at least one element")
	}
}

// Test_MediaImage focuses on the Image media type, ensuring that it is correctly loaded
// in the Medias registry and that file types and extensions can be retrieved for it.
func Test_MediaImage(t *testing.T) {

	// Retrieve the collection of file types for the Image media type.
	images := Medias.GetFileTypeExtensions(Image)

	// Ensure the Image media type is registered and not empty.
	if len(images) == 0 {
		t.Errorf("Expected images to have at least one element")
	}

	// Retrieve all file types associated with the Image media type.
	fileTypes := images.GetFileTypes()

	// Ensure that at least one file type is defined for images.
	if len(fileTypes) == 0 {
		t.Errorf("Expected fileTypes to have at least one element")
	}

	// Retrieve extensions for the first file type.
	imgExts := images.GetExtensions(fileTypes[0])

	// Ensure the file type has at least one extension.
	if len(imgExts) == 0 {
		t.Errorf("Expected imgExts to have at least one element")
	}

	// Retrieve all extensions for all file types under the Image media type.
	allExts := images.GetAllExtensions()

	// Ensure there is at least one image extension overall.
	if len(allExts) == 0 {
		t.Errorf("Expected allExts to have at least one element")
	}
}

// Test_GetMediaTypeByExtension tests the GetMediaTypeByExtension function with a map of extensions to expected media types.
func Test_GetMediaTypeByExtension(t *testing.T) {
	// Define a test map of extensions to their expected media types.
	extensionToMediaType := map[types.FileExtension]types.MediaType{
		media_image.ExtensionJpg: Image,
		media_image.ExtensionPng: Image,
		// media_video.ExtensionMp4: Video,
	}

	for ext, expectedType := range extensionToMediaType {
		// Act: get the media type for the current extension.
		mediaType := GetMediaTypeByExtension(ext)

		// Assert: check if the returned media type matches the expected one.
		if mediaType != expectedType {
			t.Errorf("Expected media type for extension %s to be %s, but got %s", ext, expectedType, mediaType)
		}
	}
}
