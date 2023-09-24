package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// My Theme
// Always dark and scaled
type scaledTheme struct {
	fyne.Theme
	scale float32
}

func (st scaledTheme) Size(tsn fyne.ThemeSizeName) float32 {
	return st.Theme.Size(tsn) * st.scale
}

func newScaledTheme(scale float32) scaledTheme {
	return scaledTheme{
		Theme: theme.DefaultTheme(),
		scale: scale, //just a little nicer
	}
}

func applyTheme(scale float32) {
	myTheme := newScaledTheme(scale)
	fyne.CurrentApp().Settings().SetTheme(myTheme)
}

func platformDefaultScale[T float32 | float64]() T {
	if fyne.CurrentDevice().IsMobile() {
		return 2.5
	} else {
		return 1.2
	}
}

// My layout
// Works like a VBox, but
// Bottom up instead of top down
// Width is stretched
func NewBottomUpWidthStretched(objects ...fyne.CanvasObject) *fyne.Container {
	return &fyne.Container{Layout: NewBottomUpWidthStretchedLayout(), Objects: objects}
}

type bottomUpWidthStretchedLayout struct {
}

func NewBottomUpWidthStretchedLayout() fyne.Layout {
	return &bottomUpWidthStretchedLayout{}
}

func (c *bottomUpWidthStretchedLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	var offset float32
	for i := len(objects) - 1; i >= 0; i-- {
		child := objects[i]
		childMin := child.MinSize()
		child.Resize(fyne.NewSize(size.Width, childMin.Height))
		child.Move(fyne.NewPos(0, size.Height-childMin.Height-offset))
		offset += childMin.Height
	}
}

func (c *bottomUpWidthStretchedLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		minSize = minSize.Max(child.MinSize())
	}

	return minSize
}
