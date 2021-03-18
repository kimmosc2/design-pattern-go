package arm

type ak47 struct {
	gun
}

func NewAK47() Arm {
	return &ak47{}
}
