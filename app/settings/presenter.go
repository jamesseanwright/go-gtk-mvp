package settings

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/framework"
)

type SettingsPresenter struct {
	navigator *framework.Navigator
}

func NewPresenter(navigator *framework.Navigator) SettingsPresenter {
	return SettingsPresenter{
		navigator,
	}
}

func (p SettingsPresenter) Bind(builder *gtk.Builder) error {
	var err error

	backButtonObj, err := builder.GetObject("back-button")
	backButton, isButtonValid := backButtonObj.(*gtk.Button)

	if !isButtonValid {
		return errors.New("Unable to operate upon back button")
	}

	backButton.Connect("clicked", func() {
		p.navigator.Navigate("index")
	})

	return err
}
