package index

import (
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

func (p MainPresenter) Bind(view MainView) error {
	return view.RegisterSettingsNavigationHandler(func() {
		p.navigator.Navigate("settings")
	})
}
