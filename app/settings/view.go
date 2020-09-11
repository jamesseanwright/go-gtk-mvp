package settings

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
)

type SettingsView struct {
	builder   *gtk.Builder
	presenter SettingsPresenter
}

func NewView(presenter SettingsPresenter) SettingsView {
	return SettingsView{
		builder:   nil,
		presenter: presenter,
	}
}

// TODO: move to base view struct
func (v SettingsView) Bind(builder *gtk.Builder) error {
	v.builder = builder
	return v.presenter.Bind(v)
}

/* TODO: implement backwards navigation
 * with a navigation stack instead */
func (v SettingsView) RegisterBackNavigationHandler(handler func()) error {
	var err error

	backButtonObj, err := v.builder.GetObject("back-button")
	backButton, isButtonValid := backButtonObj.(*gtk.Button)

	if !isButtonValid {
		return errors.New("Unable to operate upon back button")
	}

	backButton.Connect("clicked", handler)

	return err
}
