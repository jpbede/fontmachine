// Package machinery contains all necessary functions for generating a SDF PBF
package machinery

import (
	"fmt"
	"github.com/go-courier/fontnik"
	"github.com/golang/freetype/truetype"
	"github.com/golang/protobuf/proto"
	"strconv"
	"strings"
)

type FontMachinery struct {
	FontPath string
}

// NewFontMachinery creates a new FontMachinery
func NewFontMachinery(opts ...func(*FontMachinery)) *FontMachinery {
	machinery := &FontMachinery{}
	for _, optFunc := range opts {
		optFunc(machinery)
	}
	return machinery
}

// ComposeFontstack loads and generates a PBF of given fontstack
func (fm *FontMachinery) ComposeFontstack(fontStack, fontRange string) ([]byte, error) {
	fonts := strings.Split(fontStack, ",")
	minMax := strings.Split(fontRange, "-")
	min, minErr := strconv.Atoi(minMax[0])
	if minErr != nil {
		return nil, minErr
	}
	max, maxErr := strconv.Atoi(minMax[1])
	if maxErr != nil {
		return nil, maxErr
	}
	return fm.ComposeByFontNames(fonts, min, max)
}

// ComposeByFontNames loads and generates a PBF of given font names
func (fm *FontMachinery) ComposeByFontNames(fontNames []string, min, max int) ([]byte, error) {
	var fonts []*truetype.Font

	// loop all names, read and parse the fonts
	for _, fontName := range fontNames {
		font, readError := readFont(fmt.Sprintf("%s/%s.ttf", fm.FontPath, strings.TrimSpace(fontName)))
		if readError != nil {
			return nil, readError
		}
		fonts = append(fonts, font)
	}

	return fm.ComposeByFonts(fonts, min, max)
}

// ComposeByFonts generates a PBF of given truetype fonts
func (fm *FontMachinery) ComposeByFonts(fonts []*truetype.Font, min, max int) ([]byte, error) {
	result := fontnik.Glyphs{}

	for _, font := range fonts {
		builder := fontnik.NewSDFBuilder(font)
		glyphs := builder.Glyphs(min, max)
		result.Stacks = append(result.Stacks, glyphs.Stacks...)
	}

	composedPBF, marshalErr := proto.Marshal(&result)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return composedPBF, nil
}
