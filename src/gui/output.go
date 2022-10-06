package gui

// Gtk3 Renderer
// http://developer.gnome.org/gtk3/stable/GtkCellRenderer.html

import (
	//go-gtk

	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Output struct {
	//gtk window
	window *gtk.Window
	//renderer
	renderer *gtk.CellRendererAccel
}

// NewOutput creates a new output window
func NewOutput(conffile string) *Output {
	conf, err := LoadConfig(conf)
	if err != nil {
		panic(err)
	}
	//create window
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("")