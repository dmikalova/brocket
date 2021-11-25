package wm

import (
	"errors"
	"log"
	"os/exec"
	"strings"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
)

type Conf struct {
	Cmd   string
	Class string
	List  bool
}

type window struct {
	class string
	id    xproto.Window
}

type wm struct {
	conf Conf
	x    *xgbutil.XUtil
}

func NewWM(c Conf) wm {
	if c.Class == "" {
		panic(errors.New("No class found for searching"))
	}

	wm := wm{
		conf: c,
	}
	wm.connect()
	return wm
}

func (wm *wm) active() xproto.Window {
	a, err := ewmh.ActiveWindowGet(wm.x)
	if err != nil {
		panic(err.Error())
	}
	return a
}

func (wm *wm) connect() {
	x, err := xgbutil.NewConn()
	if err != nil {
		panic(err.Error())
	}
	wm.x = x
}

func (wm *wm) list() {
	ws := wm.windows()
	log.Println("id\t\tclass")
	for _, w := range ws {
		log.Printf("%d\t%s\n", w.id, w.class)
	}
}

func (wm *wm) run() {
	runCmd := wm.conf.Cmd
	if runCmd == "" {
		runCmd = wm.conf.Class
	}
	cmd := exec.Command(runCmd)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func (wm *wm) runOrRaise() {
	a := wm.active()
	ws := wm.windows()
	for _, w := range wm.stack() {
		if ws[w].class == wm.conf.Class {
			// skip active window
			if w == a {
				continue
			}
			ewmh.ActiveWindowReq(wm.x, w)
			return
		}
	}
	wm.run()
}

func (wm *wm) stack() (s []xproto.Window) {
	s, err := ewmh.ClientListStackingGet(wm.x)
	if err != nil {
		panic(err.Error())
	}

	// Reverse the stack list
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	return s
}

func (wm *wm) windows() map[xproto.Window]window {
	ws := map[xproto.Window]window{}
	for _, id := range wm.stack() {
		class, err := icccm.WmClassGet(wm.x, id)
		if err != nil {
			panic(err.Error())
		}
		ws[id] = window{
			class: strings.ToLower(class.Class),
			id:    id,
		}
	}
	return ws
}

func RunOrRaise(c Conf) {
	// connect to X server
	wm := NewWM(c)
	if wm.conf.List == true {
		wm.list()
		return
	}
	wm.runOrRaise()
}

// brocket.sh -c firefox Firefox; wmctrl -r :ACTIVE: -b remove,maximized_vert; wmctrl -r :ACTIVE: -b remove,maximized_horz; wmctrl -r :ACTIVE: -e 0,0,20,800,600

// brocket.sh -c code; wmctrl -r :ACTIVE: -b remove,maximized_vert; wmctrl -r :ACTIVE: -b remove,maximized_horz; wmctrl -r :ACTIVE: -e 0,0,20,800,600
