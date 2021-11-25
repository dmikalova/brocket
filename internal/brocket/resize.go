package brocket

func Resize(c Conf) {
	wm := NewWM(c)
	wm.resize()
}
