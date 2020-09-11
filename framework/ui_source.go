package framework

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
)

type UiSource interface {
	RegisterEvent(controlName, eventName string, handler func()) error
}

type GtkUiSource struct {
	builder *gtk.Builder
}

func NewGtkUiSource(builder *gtk.Builder) GtkUiSource {
	return GtkUiSource{builder}
}

func (s GtkUiSource) RegisterEvent(controlName, eventName string, handler func()) error {
	var err error

	obj, err := s.builder.GetObject(controlName)

	// TODO: make generic for all connectable widgets
	button, isButtonValid := obj.(*gtk.Button)

	if !isButtonValid {
		err = errors.New("Unable to operate upon settings button")
	}

	button.Connect(eventName, handler)

	return err
}
