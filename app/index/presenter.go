package index

import (
	"james.engineering/hello-go-gtk/framework"
	"james.engineering/hello-go-gtk/services"
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
	quoteChan := services.FetchSwansonQuote()

	go func() {
		quote := <-quoteChan
		view.DisplayQuote(quote)
	}()

	return view.RegisterSettingsNavigationHandler(func() {
		p.navigator.Navigate("settings")
	})
}
