package machinery

import (
	"github.com/golang/freetype/truetype"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewFontMachinery(t *testing.T) {
	expected := &FontMachinery{}
	actual := NewFontMachinery()

	if !cmp.Equal(expected, actual) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}

func TestFontMachinery_ComposeByFontNames(t *testing.T) {
}

func TestFontMachinery_ComposeByFonts(t *testing.T) {
	var fonts []*truetype.Font
	machinery := NewFontMachinery()
	font, rErr := readFont("../fonts/Roboto-Bold.ttf")
	if rErr != nil {
		t.Errorf("Got error while reading font: %s", rErr.Error())
		t.FailNow()
	}
	fonts = append(fonts, font)

	_, err := machinery.ComposeByFonts(fonts, 0, 255)
	if err != nil {
		t.Errorf("Got error while composing: %s", err.Error())
	}
}

func TestFontMachinery_ComposeFontstack(t *testing.T) {

}
