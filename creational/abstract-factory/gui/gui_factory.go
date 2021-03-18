package gui

import (
	"errors"
)

type UIFactory interface {
	createInput() IInput
	createButton() IButton
}

func NewUIFactory(platform string) (UIFactory, error) {
	switch platform {
	case "win":
		return winUI{}, nil
	case "web":
		return webUI{}, nil
	default:
		return nil,errors.New("no specified platform")
	}
}
