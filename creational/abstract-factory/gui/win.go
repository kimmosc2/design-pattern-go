package gui

type winUI struct{}

func (w winUI) createInput() IInput {
	return &winInput{}
}

func (w winUI) createButton() IButton {
	return &winButton{}
}
