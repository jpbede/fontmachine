// Package machinery contains all necessary functions for generating a SDF PBF
package machinery

import (
	"github.com/go-courier/fontnik"
	"github.com/golang/freetype/truetype"
	"github.com/golang/protobuf/proto"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FontMachinery struct {
	FontPath string
	fontSize float64

	foundFonts map[string]string
}

// NewFontMachinery creates a new FontMachinery
func NewFontMachinery(opts ...func(*FontMachinery)) *FontMachinery {
	machinery := &FontMachinery{}
	for _, optFunc := range opts {
		optFunc(machinery)
	}

	machinery.ScanFontDirectory()

	return machinery
}

// GetAvailableFonts returns a list of available fonts
func (fm *FontMachinery) GetAvailableFonts() {

}

// ScanFontDirectory scans the given font directory to find all fonts
func (fm *FontMachinery) ScanFontDirectory() {
	if fm.foundFonts == nil {
		fm.foundFonts = make(map[string]string)
	}

	fileList := []string{}
	filepath.Walk(fm.FontPath, func(path string, f os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext == ".ttf" || ext == ".otf" {
			fileList = append(fileList, path)
		}
		return nil
	})

	for _, file := range fileList {
		if font, err := readFont(file); err == nil {
			fontName := font.Name(truetype.NameIDFontFullName)
			// the regular fonts doesn't have the "Regular" part in their name
			// but this is needed so add it
			if strings.Contains(file, "Regular") {
				fontName += " Regular"
			}
			fm.foundFonts[fontName] = file
		}
	}
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
		fontPath := fm.foundFonts[fontName]
		font, readError := readFont(fontPath)
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

	var opts fontnik.SDFBuilderOpt
	if fm.fontSize > 0 {
		opts = fontnik.SDFBuilderOpt{
			FontSize: fm.fontSize,
		}
	}

	for _, font := range fonts {
		builder := fontnik.NewSDFBuilder(font, opts)
		glyphs := builder.Glyphs(min, max)
		result.Stacks = append(result.Stacks, glyphs.Stacks...)
	}

	composedPBF, marshalErr := proto.Marshal(&result)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return composedPBF, nil
}
