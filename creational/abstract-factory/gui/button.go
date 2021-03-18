package gui

type button struct {
	name          string
	onClickHandle func()
	onHoverHandle func()
}

func (b *button) OnClick(f func()) {
	b.onClickHandle = f
}

func (b *button) SetName(name string) {
	b.name = name
}

func (b *button) OnHover(f func()) {
	b.onHoverHandle = f
}

type IButton interface {
	OnClick(func())
	SetName(name string)
	OnHover(func())
}

