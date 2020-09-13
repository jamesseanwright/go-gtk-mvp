package main

import (
	"log"

	"github.com/gotk3/gotk3/gio"
	"github.com/gotk3/gotk3/gtk"
	"james.engineering/hello-go-gtk/app/index"
	"james.engineering/hello-go-gtk/app/settings"
	"james.engineering/hello-go-gtk/framework"
)

func registerResources() error {
	gresource, err := gio.LoadGResource("resources")

	if err != nil {
		return err
	}

	gio.RegisterGResource(gresource)

	return err
}

func main() {
	gtk.Init(nil)

	var err error

	err = registerResources()

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Hello GTK")

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetDefaultSize(640, 480)

	cssProvider, err := gtk.CssProviderNew()
	err = cssProvider.LoadFromPath("app/style.css")
	screen := win.GetScreen()

	gtk.AddProviderForScreen(
		screen,
		cssProvider,
		gtk.STYLE_PROVIDER_PRIORITY_APPLICATION,
	)

	navigator := framework.NewNavigator(win)

	views := map[string]framework.View{
		"index":    index.NewView(index.NewPresenter(&navigator)),
		"settings": settings.NewView(settings.NewPresenter(&navigator)),
	}

	navigator.SetViews(views)

	navigator.Navigate("index")

	if err != nil {
		log.Fatal(err)
		gtk.MainQuit()
	}

	gtk.Main()
}
