package index

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/framework"
)

type MainPresenter struct {
	navigator *framework.Navigator
}

func NewPresenter(navigator *framework.Navigator) MainPresenter {
	return MainPresenter{
		navigator,
	}
}

func (presenter MainPresenter) Bind(builder *gtk.Builder) error {
	var err error

	settingsButtonObj, err := builder.GetObject("settings-button")
	settingsButton, isButtonValid := settingsButtonObj.(*gtk.Button)

	if !isButtonValid {
		return errors.New("Unable to operate upon settings button")
	}

	settingsButton.Connect("clicked", func() {
		presenter.navigator.Navigate("settings")
	})

	return err
}
