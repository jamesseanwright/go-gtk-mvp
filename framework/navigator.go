package framework

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type View interface {
	Bind(builder *gtk.Builder) error
}

type Presenter interface {
	Bind(view View) error
}

type Navigator struct {
	window *gtk.Window
	views  map[string]View
}

func NewNavigator(window *gtk.Window) Navigator {
	return Navigator{
		window: window,
		views:  make(map[string]View),
	}
}

// Having to set later due to circular instantiation problem
func (n *Navigator) SetViews(views map[string]View) {
	n.views = views
}

func (n *Navigator) Navigate(viewName string) {
	var err error

	view, hasView := n.views[viewName]

	if !hasView {
		// TODO: return errors instead!
		log.Fatalf("Unable to locate view %s", viewName)
		return
	}

	currentRoot, err := n.window.GetChild()

	// TODO: loaders for resource pack and file
	// system for production and dev respectively
	builder, err := gtk.BuilderNewFromResource(fmt.Sprintf("/engineering/james/hello-go-gtk/app/%s/view.ui", viewName))
	rootObject, err := builder.GetObject("root")
	rootWidget, isRootValid := rootObject.(gtk.IWidget)

	if !isRootValid {
		// TODO: return errors instead!
		log.Fatalf("Unable to load root object for view %s", viewName)
		return
	}

	err = view.Bind(builder)

	if err != nil {
		log.Fatalf("Unable to navigate to view %s: %s", viewName, err)
		return
	}

	if currentRoot != nil {
		n.window.Remove(currentRoot)
	}

	n.window.Add(rootWidget)
	n.window.ShowAll()
}
