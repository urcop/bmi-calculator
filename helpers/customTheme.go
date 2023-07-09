package helpers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type CustomTheme struct {
	fyne.Theme
}

func (m CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 246, G: 247, B: 249, A: 255}
	case theme.ColorNameInputBorder:
		return color.Black
	case theme.ColorNameInputBackground:
		return color.White
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{84, 92, 159, 255}
	case theme.ColorNameButton:
		return color.NRGBA{149, 154, 212, 255}
	case theme.ColorNameForeground:
		return color.NRGBA{84, 92, 159, 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
