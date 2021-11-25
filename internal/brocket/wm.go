package brocket

import (
	"fmt"
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
	// Frame  bool
	List   bool
	Height int
	Width  int
	X      int
	Y      int
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

func (wm *wm) desktop() uint {
	d, err := ewmh.CurrentDesktopGet(wm.x)
	if err != nil {
		panic(err.Error())
	}
	return d
}

type workarea struct {
	x int
	y int
	h int
	w int
}

func (wm *wm) workarea() workarea {
	ws, err := ewmh.WorkareaGet(wm.x)
	if err != nil {
		panic(err.Error())
	}
	w := ws[wm.desktop()]
	return workarea{
		x: int(w.X),
		y: int(w.Y),
		h: int(w.Height),
		w: int(w.Width),
	}
}

type frame struct {
	left int
	top  int
}

// frame multiplier for gtk windows
func (wm *wm) gtkclass(w xproto.Window) bool {
	gtk := []string{"firefox", "lens"}
	for _, c := range gtk {
		if c == wm.class(w) {
			return true
		}
	}
	return false
}

func (wm *wm) frame(a xproto.Window) *ewmh.FrameExtents {
	if wm.gtkclass(a) {
		fmt.Println("reframing")
		f, err := ewmh.FrameExtentsGet(wm.x, a)
		if err != nil {
			panic(err.Error())
		}
		return &ewmh.FrameExtents{
			Left:   f.Left / 2,
			Right:  f.Right/2 - 1,
			Top:    f.Top,
			Bottom: f.Bottom/2 - 1,
		}
	}

	return &ewmh.FrameExtents{
		Left:   0,
		Right:  0,
		Top:    0,
		Bottom: 0,
	}
}

func (wm *wm) resize() {
	a := wm.active()
	f := wm.frame(a)
	wa := wm.workarea()
	x := (wa.w*wm.conf.X)/100 + wa.x - f.Left
	y := (wa.h*wm.conf.Y)/100 + wa.y - f.Top
	w := wa.w*wm.conf.Width/100 - f.Right
	h := wa.h*wm.conf.Height/100 - f.Bottom

	fmt.Println("xywh", x, y, w, h)
	// fmt.Println("workarea", wa)
	// fmt.Println("a", a)

	ewmh.WmStateReq(wm.x, a, ewmh.StateRemove, "_NET_WM_STATE_MAXIMIZED_VERT")
	ewmh.WmStateReq(wm.x, a, ewmh.StateRemove, "_NET_WM_STATE_MAXIMIZED_HORZ")

	err := ewmh.MoveresizeWindow(wm.x, a, x, y, w, h)
	if err != nil {
		panic(err.Error())
	}

	// hints, err := icccm.WmNormalHintsGet(wm.x, a)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("hints", hints)

	// fmt.Println(wm.conf)
	// fmt.Println(w)
}

// brocket.sh -c firefox Firefox; wmctrl -r :ACTIVE: -b remove,maximized_vert; wmctrl -r :ACTIVE: -b remove,maximized_horz; wmctrl -r :ACTIVE: -e 0,0,20,800,600

// brocket.sh -c code; wmctrl -r :ACTIVE: -b remove,maximized_vert; wmctrl -r :ACTIVE: -b remove,maximized_horz; wmctrl -r :ACTIVE: -e 0,0,20,800,600

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
	running := false
	for _, w := range wm.stack() {
		if ws[w].class == wm.conf.Class {
			running = true
			// skip active window
			if w == a {
				continue
			}
			ewmh.ActiveWindowReq(wm.x, w)
			return
		}
	}
	if running {
		return
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

func (wm *wm) class(w xproto.Window) string {
	c, err := icccm.WmClassGet(wm.x, w)
	if err != nil {
		panic(err.Error())
	}
	return strings.ToLower(c.Class)
}

func (wm *wm) windows() map[xproto.Window]window {
	ws := map[xproto.Window]window{}
	for _, w := range wm.stack() {
		c := wm.class(w)
		ws[w] = window{
			class: c,
			id:    w,
		}
	}
	return ws
}
