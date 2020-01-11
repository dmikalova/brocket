GPL License:

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    <http://www.gnu.org/licenses/>

Usage:
    brocket [-hHLv] [-m|n][-a|p|R][-c] COMMAND [TITLE]

    brocket attempts to see if an application is running, and if so, bring the
     application's window to the front, otherwise launch an instance.
     In essence, a single window mode.

Arguments:
    COMMAND   The shell command to search the window classes with.
    TITLE     The window title to search.(case insensitive substring)

              A title search is optional and the command will be used instead.
              A command search should be sufficient to find a general program.
              Multiple results will be cycled through.

Modes:
    -a   Move to the desktop with the window.
    -p   Only activate windows on the current desktop. (default)
    -R   Bring the window from any desktop.

    -c   Ignore the window class and only do a window title search.
    -m   Move the active window to the desktop on the left.
    -n   Move the active window to the desktop on the right.

Help:
    -h   Print help.
    -H   Print more help.
    -L   GPL License.
    -v   Be verbose.

Examples:
    brocket kate
    brocket Dolphin dolphin
    brocket "LibreOffice Writer" libreoffice
    brocket -R amarok
    brocket -Rv konsole
    brocket -ac gwenview
    brocket -m

Compatability:
    - brocket is based on wmctrl, and should work with the same WMs as wmctrl.
    - wmctrl is compatible with 'EWMH/NetWM compatible X Window Managers'
       Explicitly: Enlightenment, IceWM, KWin, Sawfish, and Xfce.
       Implicitly: Awesome, Fluxbox, Compiz, Openbox, and Metacity,
                   (and XMonad if you turn on EWHM), and probably others.
    - Reportedly, Awesome, Gnome Shell, Fluxbox, and XMonad have run-or-raise.
    - With all that being said, brocket has only been developed for/on Kwin.

Peculiarities:
    - The trumping order of modes is [-m|n][-a|p|R]
    - The oldest window is the first result.

Suggestions:
    - Try making a global shortcut, i.e. map 'Super+v' to 'brocket vlc'.
    - Try aliases in ~/.bashrc. This will only affect your terminals, and not
       krunner or KDE4's global shortcuts. Run multiple instances by scaping.

Future Plans:
    - Make sure application is focused after launch.
    - Center window function.
    - Include ~/.local/share/applications/ for .desktop files.
    - Check to make sure brocket isn't already running.
    - Use xdotool instead of wmctrl or in conjunction, may be faster.
    - Rewrite in Python, may be faster.

    - Launch onto specific desktops. (ie always launch Amarok on desktop 2)
    - Raise to specific screens.
    - -aR from specific desktops. (like a -p 1 option)

    - Raise from tray (may not be possible with wmctrl)
    - Close window after inactivity time, which is not possible in wmctrl.

    - Use lockrun to really ensure only one instance on startup
    - Faster!

Information:
    brocket v2.2 by dmikalova
    https://gitlab.com/dmikalova/brocket