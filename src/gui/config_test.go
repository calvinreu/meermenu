package gui

import "testing"

// test if config is loaded correctly
func TestLoadConfig(t *testing.T) {
	//test config
	config, err := LoadConfig("monitor: 1\nwidth: 100px\nheight: 100\nx: 100\ny: 100\ncssfile: test.css\ntitle: test")
	if err != nil {
		t.Error(err)
	}
	if config.Monitor != 1 {
		t.Error("Monitor not loaded correctly")
	}
	if config.Width != 100 {
		t.Error("Width not loaded correctly")
	}
	if config.Height != 100 {
		t.Error("Height not loaded correctly")
	}
	if config.X != 100 {
		t.Error("X not loaded correctly")
	}
	if config.Y != 100 {
		t.Error("Y not loaded correctly")
	}
	if config.Cssfile != "test.css" {
		t.Error("Cssfile not loaded correctly")
	}
	if config.Title != "test" {
		t.Error("Title not loaded correctly")
	}
}
