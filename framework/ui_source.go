package framework

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

type UiSource interface {
	RegisterEvent(controlName, eventName string, handler func()) error
	SetImageProperty(controlName, prop string, value interface{}) error
	SetLabelProperty(controlName, prop string, value interface{}) error
}

type GtkUiSource struct {
	builder *gtk.Builder
}

func NewGtkUiSource(builder *gtk.Builder) GtkUiSource {
	return GtkUiSource{builder}
}

/* gotk3 seemingly doesn't support common
 * property operations on *gtk.Widget, hence we
 * need generic setter functions per type :(
 * TODO: share code between setters */
func (s GtkUiSource) SetImageProperty(controlName, prop string, value interface{}) error {
	var err error

	obj, err := s.builder.GetObject(controlName)

	widget, isWidget := obj.(*gtk.Image)

	if !isWidget {
		err = fmt.Errorf("%s is not a GTK image", controlName)
	}

	widget.SetProperty(prop, value)

	return err
}

func (s GtkUiSource) SetLabelProperty(controlName, prop string, value interface{}) error {
	var err error

	obj, err := s.builder.GetObject(controlName)

	widget, isWidget := obj.(*gtk.Label)

	if !isWidget {
		err = fmt.Errorf("%s is not a GTK label", controlName)
	}

	widget.SetProperty(prop, value)

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
