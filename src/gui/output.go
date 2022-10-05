package gui

// Gtk3 Renderer
// http://developer.gnome.org/gtk3/stable/GtkCellRenderer.html

import (
	//go-gtk

	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

type Output struct {
	//gtk window
	window *gtk.Window
	//renderer
	renderer *gtk.CellRendererAccel
}

func LoadConfig(conf string) (Config, error) {
	buffer := make([]byte, 1024)
	output := Output{
		window:   gtk.NewWindow(gtk.WINDOW_POPUP),
		renderer: gtk.NewCellRendererAccel(),
	}

	//remove window border
	output.window.SetDecorated(false)

	//load file
	file, err := os.Open(conf)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//load config
	//Format
	// name:value

	length, err := file.Read(buffer)

	if err != nil {
		return output, err
	}

	//parse config
	//search for width
	var width int
	match := regexp.MustCompile(`width:(\d+)`)
	strmatch := match.FindStringSubmatch(string(buffer[:length]))
	if len(strmatch) > 1 {
		//check if width is in % or px
		if strings.Contains(strmatch[1], "%") {
			//format width
			strmatch[1] = strings.Replace(strmatch[1], "%", "", -1)
			//convert to int
			width, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
			//get screen width
			width = gdk.ScreenWidth() * width / 100
		} else {
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			width, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
		}
	}

	//search for height
	var height int
	match = regexp.MustCompile(`height:(\d+)`)
	strmatch = match.FindStringSubmatch(string(buffer[:length]))
	if len(strmatch) > 1 {
		//check if height is in % or px
		if strings.Contains(strmatch[1], "%") {
			//format height
			strmatch[1] = strings.Replace(strmatch[1], "%", "", -1)
			//convert to int
			height, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
			//get screen height
			height = gdk.ScreenHeight() * height / 100
		} else {
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			height, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
		}
	}

	//search for x position
	var x int
	match = regexp.MustCompile(`x:(\d+)`)
	strmatch = match.FindStringSubmatch(string(buffer[:length]))
	if len(strmatch) > 1 {
		//check if x is in % or px
		if strings.Contains(strmatch[1], "%") {
			//format x
			strmatch[1] = strings.Replace(strmatch[1], "%", "", -1)
			//convert to int
			x, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
			//get screen width
			x = gdk.ScreenWidth() * x / 100
		} else {
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			x, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
		}
	}

	//search for y position
	var y int
	match = regexp.MustCompile(`y:(\d+)`)
	strmatch = match.FindStringSubmatch(string(buffer[:length]))
	if len(strmatch) > 1 {
		//check if y is in % or px
		if strings.Contains(strmatch[1], "%") {
			//format y
			strmatch[1] = strings.Replace(strmatch[1], "%", "", -1)
			//convert to int
			y, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
			//get screen height
			y = gdk.ScreenHeight() * y / 100
		} else {
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			y, err := strconv.Atoi(strmatch[1])
			if err != nil {
				return output, err
			}
		}
	}

	output.window.SetTitle("menu")

	output.window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	output.window.ShowAll()
	return output, nil
}
