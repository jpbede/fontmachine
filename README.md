# Fontmachine
[![](https://godoc.org/github.com/jpbede/fontmachine?status.svg)](http://godoc.org/github.com/jpbede/fontmachine)
[![](https://img.shields.io/docker/cloud/build/jpbe/fontmachine)](https://hub.docker.com/r/jpbe/fontmachine)

A small library and server to generate SDF font glyphs, written in Golang.

The http server implements the Mapbox GL Glyphs API (`/{fontstack}/{range}.pbf`)

## Using Fontmachine
There are serval ways to use Fontmachine.

### Docker
When you want to run Fontmachine with docker, use the following command:

`docker run -p 8080:8080 -v {font path}:/fonts jpbe/fontmachine`

Place all font files (ttf) in a folder and replace `{font path}` in the above command with the path.

### Hosted
I provide a hosted version of Fontmachine, which is reachable at `use.fontmachine.io`

If you want to use this just replace your glyph path with `https://use.fontmachine.io/{fontstack}/{range}.pbf`

#### Available fonts
- All Roboto fonts
- All Open Sans fonts
- All Noto fonts

#### Usage policy
This kind of service is being provided with best effort. Be kind and do not stress this service.
I will ratelimit the requests if needed. 

Please contact me before using this service in a large manner.

I would be happy if you would sponsor to keep this service running.