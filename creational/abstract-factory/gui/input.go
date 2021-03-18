package gui

type input struct {
	name           string
	text           string
	onChangeHandle func()
}

func (i *input) SetName(name string) {
	i.name = name
}

func (i *input) SetText(text string) {
	i.text = text
}

func (i *input) OnChange(f func()) {
	i.onChangeHandle = f
}

type IInput interface {
	SetName(name string)
	SetText(text string)
	OnChange(func())
}
