package machinery

func WithFontPath(path string) func(*FontMachinery) {
	return func(machinery *FontMachinery) {
		machinery.FontPath = path
	}
}
