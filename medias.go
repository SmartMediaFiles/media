// Package media is the central point of the SmartMediaFiles ecosystem.
//
// It acts as an aggregator that gathers definitions from all specialized
// media type libraries (images, videos, RAW, etc.) into a single registry.
// This package provides a unified overview of all supported file formats,
// facilitating their identification and handling.
//
// The main components of this package are:
//   - Constant definitions for each major media type (Raw, Image, Video, etc.).
//   - The `Medias` variable, a comprehensive map that associates each media type
//     with its corresponding file types and extensions.
//
// Using this package is the first step for any application looking to work
// with multimedia files within the SmartMediaFiles ecosystem.
package media

import (
	media_image "github.com/smartmediafiles/media.image"
	media_raw "github.com/smartmediafiles/media.raw"
	media_sidecar "github.com/smartmediafiles/media.sidecar"
	media_unknown "github.com/smartmediafiles/media.unknown"
	media_vector "github.com/smartmediafiles/media.vector"
	media_video "github.com/smartmediafiles/media.video"
	"github.com/smartmediafiles/media/media/maps"
	"github.com/smartmediafiles/media/media/types"
)

// List of the main media types supported by the ecosystem.
// These constants serve as keys to access specific definitions
// in the `Medias` map.
const (
	// Raw represents raw image files from digital cameras.
	// These files contain minimally processed data from the image sensor.
	// Examples: .CR2, .NEF, .ARW.
	Raw types.MediaType = "raw"

	// Image represents standard raster image files.
	// These formats are widely used for display and on the web.
	// Examples: .JPEG, .PNG, .GIF.
	Image types.MediaType = "image"

	// Video represents video files.
	// These formats contain both video and audio tracks.
	// Examples: .MP4, .MOV, .AVI.
	Video types.MediaType = "video"

	// Vector represents vector graphic files.
	// Unlike raster images, they are based on mathematical equations.
	// Examples: .SVG, .AI.
	Vector types.MediaType = "vector"

	// Sidecar represents files that store metadata about other files.
	// They often accompany RAW or video files.
	// Examples: .XMP, .MIE.
	Sidecar types.MediaType = "sidecar"

	// Unknown represents an unidentified or unsupported media type.
	// It is the default value for files with unrecognized extensions.
	Unknown types.MediaType = ""
)

// Medias is the central registry of all supported media and file types.
//
// This variable is a map that associates each `MediaType` (defined above)
// with a `MapFileTypeExtensions`, which in turn maps specific `FileType`
// to their lists of `FileExtension`.
//
// # Example Usage:
//
// To get all file extensions for JPEG images:
//
//	jpegExtensions := media.Medias[media.Image][media_image.ImageJpeg]
//	fmt.Println(jpegExtensions) // Outputs: [.jpg .jpeg .jpe .jif .jfif .jfi]
//
// To get all file types for videos:
//
//	videoFileTypes := media.Medias.GetFileTypes(media.Video)
//	fmt.Println(videoFileTypes)
var Medias = maps.MapMediaTypeFiles{
	Raw:     media_raw.RawFileTypesExtensions,
	Image:   media_image.ImageFileTypesExtensions,
	Video:   media_video.VideoFileTypesExtensions,
	Vector:  media_vector.VectorFileTypesExtensions,
	Sidecar: media_sidecar.SidecarFileTypesExtensions,
	Unknown: media_unknown.UnknownFileTypesExtensions,
}
