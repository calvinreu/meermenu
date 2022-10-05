package gui

import "github.com/mattn/go-gtk/gtk"

type Entry struct {
	Info    string
	Texture *gtk.Image
}

func NewEntry(info string, texture string) *Entry {
	return &Entry{
		Info:    info,
		Texture: gtk.NewImageFromFile(texture),
	}
}

func (entry *Entry) RefreshEntry(info string) {
	entry.Info = info
}
