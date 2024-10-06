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

// List of supported media types.
const (
	Raw     types.MediaType = "raw"
	Image   types.MediaType = "image"
	Video   types.MediaType = "video"
	Vector  types.MediaType = "vector"
	Sidecar types.MediaType = "sidecar"
	Unknown types.MediaType = ""
)

// Medias is a map of media types to their file types and extensions.
var Medias = maps.MapMediaTypeFiles{
	Raw:     media_raw.RawFileTypesExtensions,
	Image:   media_image.ImageFileTypesExtensions,
	Video:   media_video.VideoFileTypesExtensions,
	Vector:  media_vector.VectorFileTypesExtensions,
	Sidecar: media_sidecar.SidecarFileTypesExtensions,
	Unknown: media_unknown.UnknownFileTypesExtensions,
}
