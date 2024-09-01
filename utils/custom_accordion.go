package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CompactAccordionItem struct {
	widget.AccordionItem
	content fyne.CanvasObject
}

func NewCompactAccordionItem(title string, content fyne.CanvasObject) *CompactAccordionItem {
	item := &CompactAccordionItem{content: content}
	item.AccordionItem = widget.AccordionItem{Title: title, Detail: content}
	return item
}

func (cai *CompactAccordionItem) MinSize() fyne.Size {
	return cai.content.MinSize()
}

func (cai *CompactAccordionItem) Layout(size fyne.Size) {
	cai.content.Resize(size)
}
