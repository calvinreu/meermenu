package gui

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/mattn/go-gtk/gdk"
)

type Config struct {
	x, y, width, height int
}

// LoadConfig from string out of conf file
func LoadConfig(conf string) (Config, error) {
	//Format name:value
	//search for width
	var config Config
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
			config.width = gdk.ScreenWidth() * width / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			config.width, err = strconv.Atoi(strmatch[1])
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
			config.height = gdk.ScreenHeight() * height / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			height, err = strconv.Atoi(strmatch[1])
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
			config.x = gdk.ScreenWidth() * x / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			config.x, err = strconv.Atoi(strmatch[1])
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
			config.y = gdk.ScreenHeight() * y / 100
		} else {
			var err error
			//remove px
			strmatch[1] = strings.Replace(strmatch[1], "px", "", -1)
			//convert to int
			config.y, err = strconv.Atoi(strmatch[1])
			if err != nil {
				return config, err
			}
		}
	}

	return config, nil
}
