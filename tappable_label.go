package main

import (
	"time"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

const (
	buttonTapDuration = 250
)

// TappableLabel widget is a label component with appropriate padding and layout.
type TappableLabel struct {
	*widget.Label
	OnTapped func() `json:"-"`

	tapped bool
}

// NewLabel creates a new label widget with the set text content
func NewTappableLabel(text string, tapped func()) *TappableLabel {
	return NewTappableLabelWithStyle(text, fyne.TextAlignLeading, fyne.TextStyle{}, tapped)
}

// NewLabelWithStyle creates a new label widget with the set text content
func NewTappableLabelWithStyle(text string, alignment fyne.TextAlign, style fyne.TextStyle, tappedf func()) *TappableLabel {
	l := &TappableLabel{
		Label: &widget.Label{
			Text:      text,
			Alignment: alignment,
			TextStyle: style,
		},
		OnTapped: tappedf,
	}

	return l
}

func (b *TappableLabel) Tapped(*fyne.PointEvent) {
	b.tapped = true
	defer func() { // TODO move to a real animation
		time.Sleep(time.Millisecond * buttonTapDuration)
		b.tapped = false
		b.Refresh()
	}()
	b.Refresh()

	if b.OnTapped != nil {
		b.OnTapped()
	}
}

// Refresh checks if the text content should be updated then refreshes the graphical context
func (l *TappableLabel) Refresh() {
	l.Label.Refresh()
	l.Label.BaseWidget.Refresh()
}

// SetText sets the text of the label
func (l *TappableLabel) SetText(text string) {
	l.Label.SetText(text)
}
