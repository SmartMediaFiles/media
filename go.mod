module github.com/smartmediafiles/media

go 1.24

require (
	github.com/smartmediafiles/media.image v0.0.0-20241006190546-b2362cb653ac
	github.com/smartmediafiles/media.raw v0.0.0-20241006190655-0c9c19a154a5
	github.com/smartmediafiles/media.sidecar v0.0.0-20241006190813-8744519017cd
	github.com/smartmediafiles/media.unknown v0.0.0-20241006190919-245cbee872bc
	github.com/smartmediafiles/media.vector v0.0.0-20241006191016-9722d89cbbdb
	github.com/smartmediafiles/media.video v0.0.0-20241006192008-6256cd9e332e
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/smartmediafiles/media.image => ../media.image
replace github.com/smartmediafiles/media.sidecar => ../media.sidecar
replace github.com/smartmediafiles/media.unknown => ../media.unknown
replace github.com/smartmediafiles/media.raw => ../media.raw
replace github.com/smartmediafiles/media.vector => ../media.vector
replace github.com/smartmediafiles/media.video => ../media.video
