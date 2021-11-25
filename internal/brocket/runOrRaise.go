package brocket

func RunOrRaise(c Conf) {
	// connect to X server
	wm := NewWM(c)
	if wm.conf.List == true {
		wm.list()
		return
	}
	wm.runOrRaise()
}
