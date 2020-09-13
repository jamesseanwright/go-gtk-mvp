# Go GTK MVP

A GTK app, written in Go, that follows the model-view-presenter pattern. I'm building this as part of a spike on native cross-platform desktop apps.

## Project Layout

```
.
├── app => the views, presenters, and models organised by feature
│   ├── index => the main view
│   │   ├── presenter.go => glue between the view and the model
│   │   ├── view.go => provides methods (inc. getters and setters) for reading from and writing to the view
│   │   └── view.ui => XML UI definition for the view (GtkBuilder)
│   └── settings => similar to the above sibling
│       ├── presenter.go
│       ├── view.go
│       └── view.ui
├── framework => model-view-presenter framework
│   ├── navigator.go => updates the window with the requested view
│   └── ui_source.go => provides methods for querying and modifying GTK views in an abstract and isolated manner
├── services => integrations with external dependencies
├── go.mod
├── go.sum
├── main.go => program entry point
├── README.md
```

## Testing

TODO

## Running Locally

Once GTK 3 is present on your machine:

```sh
$ make build
$ cd dist
$ ./main
```
