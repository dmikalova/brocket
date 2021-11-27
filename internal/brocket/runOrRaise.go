package brocket

// RunOrRaise an application as defined in conf
func RunOrRaise(c Conf) {
	// connect to X server
	wm := newWM(c)
	if wm.conf.List == true {
		wm.list()
		return
	}
	wm.runOrRaise()
}
