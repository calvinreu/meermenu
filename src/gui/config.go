package gui

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

type Config struct {
	x, y, width, height int
}

// LoadConfig from string out of conf file
func LoadConfig(conf string) (Config, error) {
	//Format name:value
	//search for width
	var config Config

	var width int
	match := regexp.MustCompile(`width:(\d+)`)
	strmatch := match.FindStringSubmatch(conf)
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
	strmatch = match.FindStringSubmatch(conf)
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
	strmatch = match.FindStringSubmatch(conf)
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
	strmatch = match.FindStringSubmatch(conf)
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
