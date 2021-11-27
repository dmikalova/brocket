package brocket

// Resize a window as defined in Conf
func Resize(c Conf) {
	wm := newWM(c)
	wm.resize()
}
