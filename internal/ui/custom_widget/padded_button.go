package customWidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type PaddedButton struct {
	widget.Button
	Padding fyne.Size
}

func NewPaddedRunButton(tapped func(), padding fyne.Size) fyne.CanvasObject {
	button := &PaddedButton{
		Padding: padding,
	}
	button.ExtendBaseWidget(button)
	button.OnTapped = tapped
	button.Text = "Run!"
	return button
}

func (b *PaddedButton) MinSize() fyne.Size {
	size := b.Button.MinSize()
	size.Width += b.Padding.Width
	size.Height += b.Padding.Height
	return size
}
