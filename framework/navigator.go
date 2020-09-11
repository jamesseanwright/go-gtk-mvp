package framework

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type Presenter interface {
	Bind(builder *gtk.Builder) error
}

type Navigator struct {
	window     *gtk.Window
	presenters map[string]Presenter
}

func NewNavigator(window *gtk.Window) Navigator {
	return Navigator{
		window:     window,
		presenters: make(map[string]Presenter),
	}
}

// Having to set later due to circular instantiation problem
func (n *Navigator) SetPresenters(presenters map[string]Presenter) {
	n.presenters = presenters
}

func (n *Navigator) Navigate(viewName string) {
	var err error

	presenter, hasPresenter := n.presenters[viewName]

	if !hasPresenter {
		// TODO: return errors instead!
		log.Fatalf("Unable to locate view model %s", viewName)
		return
	}

	currentRoot, err := n.window.GetChild()
	builder, err := gtk.BuilderNewFromFile(fmt.Sprintf("app/%s/view.ui", viewName))
	rootObject, err := builder.GetObject("root")
	rootWidget, isRootValid := rootObject.(gtk.IWidget)

	if !isRootValid {
		// TODO: return errors instead!
		log.Fatalf("Unable to load root object for view %s", viewName)
		return
	}

	err = presenter.Bind(builder)

	if err != nil {
		log.Fatalf("Unable to navigate to view %s: %s", viewName, err)
		return
	}

	if currentRoot != nil {
		n.window.Remove(currentRoot)
	}

	n.window.Add(rootWidget)
	n.window.SetDefaultSize(640, 480)
	n.window.ShowAll()
}
