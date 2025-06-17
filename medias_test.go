package media

import (
	"testing"
)

func Test_Media(t *testing.T) {

	// Get all media types
	mediaTypes := Medias.GetMediaTypes()

	// check if mediaTypes is empty
	if len(mediaTypes) == 0 {
		t.Errorf("Expected mediaTypes to have at least one element")
	}

	// GetFile type of the first media type
	fileTypes := Medias.GetFileTypes(mediaTypes[0])

	// check if fileTypes is empty
	if len(fileTypes) == 0 {
		t.Errorf("Expected fileTypes to have at least one element")
	}

	// Get file type extensions of the first file type
	exts := Medias.GetExtensions(mediaTypes[0], fileTypes[0])

	// check if exts is empty
	if len(exts) == 0 {
		t.Errorf("Expected exts to have at least one element")
	}
}

func Test_MediaImage(t *testing.T) {

	// test media of type image
	images := Medias.GetMedia(Image)

	// check if images is empty
	if len(images) == 0 {
		t.Errorf("Expected images to have at least one element")
	}

	// Get all file types
	fileTypes := images.GetFileTypes()

	// check if fileTypes is empty
	if len(fileTypes) == 0 {
		t.Errorf("Expected fileTypes to have at least one element")
	}

	// Get extensions of the first file type
	imgExts := images.GetFileExtensions(fileTypes[0])

	// Check if imgExts is empty
	if len(imgExts) == 0 {
		t.Errorf("Expected imgExts to have at least one element")
	}

	// get all extensions
	allExts := images.GetAllExtensions()

	// check if allExts is empty
	if len(allExts) == 0 {
		t.Errorf("Expected allExts to have at least one element")
	}
}

func Test_GetMediaTypeByExtension(t *testing.T) {
	// Define a map of extensions to expected media types
	extensionToMediaType := map[string]string{
		".jpg": "Image",
		".png": "Image",
		".mp3": "Audio",
		".mp4": "Video",
	}

	for ext, expectedType := range extensionToMediaType {
		// Get the media type for the given extension
		mediaType := Medias.GetMediaTypeByExtension(ext)

		// Check if the media type matches the expected type
		if mediaType != expectedType {
			t.Errorf("Expected media type for extension %s to be %s, but got %s", ext, expectedType, mediaType)
		}
	}
}
