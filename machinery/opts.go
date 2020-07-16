package machinery

func WithFontPath(path string) func(*FontMachinery) {
	return func(machinery *FontMachinery) {
		machinery.FontPath = path
	}
}

func WithFontSize(fontSize float64) func(*FontMachinery) {
	return func(machinery *FontMachinery) {
		machinery.fontSize = fontSize
	}
}
