package gui

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gdk"
)

type Config struct {
	X, Y, Width, Height int
	Cssfile, Title      string
}

// default config
const (
	DEFAULT_X      = 500
	DEFAULT_Y      = 500
	DEFAULT_WIDTH  = 800
	DEFAULT_HEIGHT = 600
	DEFAULT_CSS    = "default.css"
	DEFAULT_TITLE  = "meermenu"
)

// LoadConfig from string out of conf file
func LoadConfig(conf string) (Config, error) {
	//Format name:value
	//search for width
	var config Config
	//get default screen
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		return config, err
	}

	//set default values
	config.X = DEFAULT_X
	config.Y = DEFAULT_Y
	config.Width = DEFAULT_WIDTH
	config.Height = DEFAULT_HEIGHT
	config.Cssfile = DEFAULT_CSS
	config.Title = DEFAULT_TITLE

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
				return config, err
			}
			//get screen width
			config.Width =  * width / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			config.Width, err = strconv.Atoi(strmatch[1])
			if err != nil {
				return config, err
			}
		}
	}

	//search for height
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
				return config, err
			}
			//get screen height
			config.Height = gdk.ScreenHeight() * height / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			config.Height, err = strconv.Atoi(strmatch[1])
			if err != nil {
				return config, err
			}
		}
	}

	//search for x position
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
				return config, err
			}
			//get screen width
			config.X = gdk.ScreenWidth() * x / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			config.X, err = strconv.Atoi(strmatch[1])
			if err != nil {
				return config, err
			}
		}
	}

	//search for y position
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
				return config, err
			}
			//get screen height
			config.Y = gdk.ScreenHeight() * y / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			config.Y, err = strconv.Atoi(strmatch[1])
			if err != nil {
				return config, err
			}
		}
	}

	//search for css file
	match = regexp.MustCompile(`css:(\S+)`)
	strmatch = match.FindStringSubmatch(conf)
	if len(strmatch) > 1 {
		config.Cssfile = strmatch[1]
	}

	//search for title
	match = regexp.MustCompile(`title:(\S+)`)
	strmatch = match.FindStringSubmatch(conf)
	if len(strmatch) > 1 {
		config.Title = strmatch[1]
	}

	return config, nil
}
