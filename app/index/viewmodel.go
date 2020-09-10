package index

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/framework"
)

type MainViewModel struct {
	navigator *framework.Navigator
}

func NewViewModel(navigator *framework.Navigator) MainViewModel {
	return MainViewModel{
		navigator,
	}
}

func (vm MainViewModel) Bind(builder *gtk.Builder) error {
	var err error

	settingsButtonObj, err := builder.GetObject("settings-button")
	settingsButton, isButtonValid := settingsButtonObj.(*gtk.Button)

	if !isButtonValid {
		return errors.New("Unable to operate upon settings button")
	}

	settingsButton.Connect("clicked", func() {
		vm.navigator.Navigate("settings")
	})

	return err
}
