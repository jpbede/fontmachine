package machinery

import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"os"
)

func readFont(path string) (*truetype.Font, error) {
	font, oErr := os.Open(path)
	if oErr != nil {
		return nil, oErr
	}
	fontByte, rErr := ioutil.ReadAll(font)
	if rErr != nil {
		return nil, rErr
	}
	parsedFont, pErr := truetype.Parse(fontByte)
	if pErr != nil {
		return nil, pErr
	}
	return parsedFont, nil
}
