package gui

// Gtk3 Renderer
// http://developer.gnome.org/gtk3/stable/GtkCellRenderer.html

import (
	//go-gtk
	"encoding/json"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

type Output struct {
	//gtk window
	window *gtk.Window
	//renderer
	renderer *gtk.CellRendererAccel
}

type OutputConfig struct {
	//window title
	Title string
	//window width
	Width int
	//window height
	Height int
}

func InitOutput(conf string) (*Output, error) {
	output := &Output{
		window:   gtk.NewWindow(gtk.WINDOW_POPUP),
		renderer: gtk.NewCellRendererAccel(),
	}
	//load conf from json
	//load file
	file, err := os.Open(conf)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//decode json
	decoder := json.NewDecoder(file)
	var config OutputConfig
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	output.window.SetTitle()
	output.window.SetSizeRequest(800, 600)
	output.window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	output.window.ShowAll()
	return output
}
