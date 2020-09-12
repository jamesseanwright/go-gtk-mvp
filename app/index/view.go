package index

import (
	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/framework"
)

type MainView struct {
	presenter MainPresenter
	uiSource  framework.UiSource
}

func NewView(presenter MainPresenter) MainView {
	return MainView{
		presenter: presenter,
		uiSource:  framework.NewGtkUiSource(nil),
	}
}

// TODO: move to base view struct
func (v MainView) Bind(builder *gtk.Builder) error {
	v.uiSource = framework.NewGtkUiSource(builder)
	return v.presenter.Bind(v)
}

func (v MainView) RegisterSettingsNavigationHandler(handler func()) error {
	return v.uiSource.RegisterEvent("settings-button", "clicked", handler)
}

func (v MainView) SetQuoteText(quote string) {
	v.uiSource.SetLabelText("quote", quote)
}
