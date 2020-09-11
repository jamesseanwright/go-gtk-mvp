package index

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
)

type MainView struct {
	builder   *gtk.Builder
	presenter MainPresenter
}

func NewView(presenter MainPresenter) MainView {
	return MainView{
		builder:   nil,
		presenter: presenter,
	}
}

// TODO: move to base view struct
func (v MainView) Bind(builder *gtk.Builder) error {
	v.builder = builder
	return v.presenter.Bind(v)
}

func (v MainView) RegisterSettingsNavigationHandler(handler func()) error {
	var err error

	settingsButtonObj, err := v.builder.GetObject("settings-button")
	settingsButton, isButtonValid := settingsButtonObj.(*gtk.Button)

	if !isButtonValid {
		return errors.New("Unable to operate upon settings button")
	}

	settingsButton.Connect("clicked", handler)

	return err
}
