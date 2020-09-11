package settings

import (
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

// TODO: implement stack-based navigation
func (p SettingsPresenter) Bind(view SettingsView) error {
	return view.RegisterBackNavigationHandler(func() {
		p.navigator.Navigate("index")
	})
}
