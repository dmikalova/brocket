package wmctrl

import (
	"fmt"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
)

func Wmctrl() {
	// c := cmd.NewCommand(`xprop -root | grep "_NET_ACTIVE_WINDOW\(WINDOW\)" | grep -o '0x[0-9a-f]+'`)
	// c := cmd.NewCommand(`printf '0x%08x\n' "$(xprop -root | grep "_NET_ACTIVE_WINDOW\(WINDOW\)" | grep -o '0x[0-9a-f]+')"`)
	// c := cmd.NewCommand(`printf '0x%08x\n' "$(xprop -root | grep "_NET_ACTIVE_WINDOW\(WINDOW\)" | grep -o '0x[0-9a-f]+')"`)
	// c := cmd.NewCommand(`printf '0x%08x\n' "$(xprop -root | grep "_NET_ACTIVE_WINDOW\(WINDOW\)" | grep -o '0x[0-9a-f]+')"`)
	// c := cmd.NewCommand(`printf '0x%08x\n' "$(xprop -root | grep "_NET_ACTIVE_WINDOW\(WINDOW\)" | grep -o '0x[0-9a-f]+')"`)

	// err := c.Execute()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(c.Stdout())
	// fmt.Println(c.Stderr())

	X, err := xgbutil.NewConn()
	w, err := ewmh.ActiveWindowGet(X)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(w)
	// get this in hex:
	// ~/.bin/brocket.sh -v

	l, err := ewmh.ClientListStackingGet(X)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(l)
}

// brocket.sh -c firefox Firefox; wmctrl -r :ACTIVE: -b remove,maximized_vert; wmctrl -r :ACTIVE: -b remove,maximized_horz; wmctrl -r :ACTIVE: -e 0,0,20,800,600

// https://github.com/BurntSushi/xgbutil/tree/master/ewmh
