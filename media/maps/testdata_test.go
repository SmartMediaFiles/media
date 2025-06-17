package maps

import "github.com/smartmediafiles/media/media/types"

// #############################################################################
//
//	Test Data for maps package
//
// #############################################################################
//
//	This file centralizes all the variables used across the different test
//	files in the 'maps' package. This approach avoids redeclaration issues
//	and keeps test data consistent.
//
// #############################################################################

// -----------------------------------------------------------------------------
// Media Types
// -----------------------------------------------------------------------------
var mediaTypeImage = types.MediaType("image")
var mediaTypeVideo = types.MediaType("video")
var mediaTypeFake = types.MediaType("fake")

// -----------------------------------------------------------------------------
// File Types
// -----------------------------------------------------------------------------
// These are used in TestMapFileTypeExtensions and represent generic file categories.
var fileTypeImage = types.FileType("image")
var fileTypeVideo = types.FileType("video")
var fileTypeFake = types.FileType("fake")

// These are used in TestMapMediaTypeFiles and represent specific file formats.
var fileTypeJpeg = types.FileType("jpeg")
var fileTypePng = types.FileType("png")
var fileTypeMp4 = types.FileType("mp4")

// -----------------------------------------------------------------------------
// File Extensions
// -----------------------------------------------------------------------------
var fileExtensionJpg = types.FileExtension(".jpg")
var fileExtensionJpeg = types.FileExtension(".jpeg")
var fileExtensionPng = types.FileExtension(".png")
var fileExtensionMp4 = types.FileExtension(".mp4")
var fileExtensionAvi = types.FileExtension(".avi")
