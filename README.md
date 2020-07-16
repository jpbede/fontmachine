# Fontmachine
[![](https://godoc.org/github.com/jpbede/fontmachine?status.svg)](http://godoc.org/github.com/jpbede/fontmachine)
![Docker Build Status](https://img.shields.io/docker/build/jpbe/fontmachine)

A small library and server to generate SDF font glyphs, written in Golang.

The http server implements the Mapbox GL Glyphs API (`/{fontstack}/{range}.pbf`)

## Using Fontmachine
There are serval ways to use Fontmachine.

### Docker
When you want to run Fontmachine with docker, use the following command:

`docker run -p 8080:8080 -v {font path}:/fonts jpbe/fontmachine`

Place all font files (ttf) in a folder and replace `{font path}` in the above command with the path.

### Hosted
_Coming soon_