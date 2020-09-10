package framework

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type ViewModel interface {
	Bind(builder *gtk.Builder) error
}

type Navigator struct {
	window     *gtk.Window
	viewModels map[string]ViewModel
}

func NewNavigator(window *gtk.Window) Navigator {
	return Navigator{
		window:     window,
		viewModels: make(map[string]ViewModel),
	}
}

// Having to set later due to circular instantiation problem
func (n *Navigator) SetViewModels(viewModels map[string]ViewModel) {
	n.viewModels = viewModels
}

func (n *Navigator) Navigate(viewName string) {
	var err error

	viewModel, hasViewModel := n.viewModels[viewName]

	if !hasViewModel {
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

	err = viewModel.Bind(builder)

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
