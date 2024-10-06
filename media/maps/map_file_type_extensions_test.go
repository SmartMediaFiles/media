package maps

import (
	"github.com/smartmediafiles/media/media/types"
	"reflect"
	"testing"
)

var fileTypeImage = types.FileType("image")
var fileTypeVideo = types.FileType("video")
var fileTypeFake = types.FileType("fake")

func TestMapFileTypeExtensions_GetFileTypes(t *testing.T) {
	m := MapFileTypeExtensions{
		fileTypeImage: {".jpg", ".png"},
		fileTypeVideo: {".mp4", ".avi"},
	}

	// map is not an ordered collection, so we can't predict the order of the elements
	expectedList1 := []types.FileType{fileTypeImage, fileTypeVideo}
	expectedList2 := []types.FileType{fileTypeVideo, fileTypeImage}
	result := m.GetFileTypes()

	if !reflect.DeepEqual(result, expectedList1) && !reflect.DeepEqual(result, expectedList2) {
		t.Errorf("GetFileTypes() = %v, want %v or %v", result, expectedList1, expectedList2)
	}
}

func TestMapFileTypeExtensions_GetExtensions(t *testing.T) {
	m := MapFileTypeExtensions{
		fileTypeImage: {".jpg", ".png"},
		fileTypeVideo: {".mp4", ".avi"},
	}

	tests := []struct {
		name     string
		fileType types.FileType
		want     []types.FileExtension
	}{
		{"Image extensions", fileTypeImage, []types.FileExtension{".jpg", ".png"}},
		{"Video extensions", fileTypeVideo, []types.FileExtension{".mp4", ".avi"}},
		{"Non-existent type", fileTypeFake, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.GetExtensions(tt.fileType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExtensions(%v) = %v, want %v", tt.fileType, got, tt.want)
			}
		})
	}
}

func TestMapFileTypeExtensions_GetAllExtensions(t *testing.T) {
	m := MapFileTypeExtensions{
		fileTypeImage: {".jpg", ".png"},
		fileTypeVideo: {".mp4", ".avi"},
	}

	expected := []types.FileExtension{".jpg", ".png", ".mp4", ".avi"}
	result := m.GetAllExtensions()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetAllExtensions() = %v, want %v", result, expected)
	}
}
