package settings

import (
	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/framework"
)

type SettingsView struct {
	presenter SettingsPresenter
	uiSource  framework.UiSource
}

func NewView(presenter SettingsPresenter) SettingsView {
	return SettingsView{
		presenter: presenter,
		uiSource:  framework.NewGtkUiSource(nil),
	}
}

// TODO: move to base view struct
func (v SettingsView) Bind(builder *gtk.Builder) error {
	v.uiSource = framework.NewGtkUiSource(builder)
	return v.presenter.Bind(v)
}

/* TODO: implement backwards navigation
 * with a navigation stack instead */
func (v SettingsView) RegisterBackNavigationHandler(handler func()) error {
	return v.uiSource.RegisterEvent("back-button", "clicked", handler)
}
