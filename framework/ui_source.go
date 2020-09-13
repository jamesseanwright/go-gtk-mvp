package framework

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

type UiSource interface {
	RegisterEvent(controlName, eventName string, handler func()) error
	SetProperty(controlName, prop string, value interface{}) error
}

type GtkUiSource struct {
	builder *gtk.Builder
}

func NewGtkUiSource(builder *gtk.Builder) GtkUiSource {
	return GtkUiSource{builder}
}

func (s GtkUiSource) SetProperty(controlName, prop string, value interface{}) error {
	var err error

	obj, err := s.builder.GetObject(controlName)

	widget, isWidget := obj.(gtk.IWidget)

	if !isWidget {
		err = fmt.Errorf("%s is not a GTK widget", controlName)
	}

	widget.ToWidget().SetProperty(prop, value)

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
