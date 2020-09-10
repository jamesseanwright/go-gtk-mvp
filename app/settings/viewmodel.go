package settings

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/framework"
)

type SettingsViewModel struct {
	navigator *framework.Navigator
}

func NewViewModel(navigator *framework.Navigator) SettingsViewModel {
	return SettingsViewModel{
		navigator,
	}
}

func (vm SettingsViewModel) Bind(builder *gtk.Builder) error {
	var err error

	backButtonObj, err := builder.GetObject("back-button")
	backButton, isButtonValid := backButtonObj.(*gtk.Button)

	if !isButtonValid {
		return errors.New("Unable to operate upon back button")
	}

	backButton.Connect("clicked", func() {
		vm.navigator.Navigate("index")
	})

	return err
}
