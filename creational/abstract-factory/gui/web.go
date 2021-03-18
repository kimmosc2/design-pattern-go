package gui

type webUI struct {
}

func (w webUI) createInput() IInput {
	return &webInput{}
}

func (w webUI) createButton() IButton {
	return &webButton{}
}
