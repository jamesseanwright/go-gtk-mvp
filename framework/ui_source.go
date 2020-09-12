package framework

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

type UiSource interface {
	RegisterEvent(controlName, eventName string, handler func()) error
	SetLabelText(controlName, text string) error
}

type GtkUiSource struct {
	builder *gtk.Builder
}

func NewGtkUiSource(builder *gtk.Builder) GtkUiSource {
	return GtkUiSource{builder}
}

func (s GtkUiSource) SetLabelText(controlName, text string) error {
	var err error

	obj, err := s.builder.GetObject(controlName)

	label, isLabel := obj.(*gtk.Label)

	if !isLabel {
		err = fmt.Errorf("%s is not a GTK label", controlName)
	}

	label.SetLabel(text)

	return err
}

func (s GtkUiSource) RegisterEvent(controlName, eventName string, handler func()) error {
	var err error

	obj, err := s.builder.GetObject(controlName)

	// TODO: make generic for all connectable widgets
	button, isButton := obj.(*gtk.Button)

	if !isButton {
		err = fmt.Errorf("%s is not a GTK button", controlName)
	}

	button.Connect(eventName, handler)

	return err
}
