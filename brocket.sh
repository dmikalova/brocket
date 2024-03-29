#!/bin/sh
time1="$(date +%s.%N)"
license="GPL License:

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    <http://www.gnu.org/licenses/>
    "

# Help
help="Usage:
    $(basename "${0}") [-hHLv] [-m|n][-a|p|R][-c] COMMAND [TITLE]

    Brocket attempts to see if an application is running, and if so, bring the
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
    $(basename "${0}") kate
    $(basename "${0}") Dolphin dolphin
    $(basename "${0}") \"LibreOffice Writer\" libreoffice
    $(basename "${0}") -R amarok
    $(basename "${0}") -Rv konsole
    $(basename "${0}") -ac gwenview
    $(basename "${0}") -m
    "

longHelp="Compatability:
    - Brocket is based on wmctrl, and should work with the same WMs as wmctrl.
    - wmctrl is compatible with 'EWMH/NetWM compatible X Window Managers'
       Explicitly: Enlightenment, IceWM, KWin, Sawfish, and Xfce.
       Implicitly: Awesome, Fluxbox, Compiz, Openbox, and Metacity,
                   (and XMonad if you turn on EWHM), and probably others.
    - Reportedly, Awesome, Gnome Shell, Fluxbox, and XMonad have run-or-raise.
    - With all that being said, Brocket has only been developed for/on Kwin.

Peculiarities:
    - The trumping order of modes is [-m|n][-a|p|R]
    - The oldest window is the first result.

Suggestions:
    - Try making a global shortcut, i.e. map 'Super+v' to 'brocket vlc'.
    - Try aliases in ~/.bashrc. This will only affect your terminals, and not
       krunner or KDE4's global shortcuts. Run multiple instances by \escaping.

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
    Brocket by David Mikalova
    https://gitlab.com/dmikalova/brocket
    "

# The default variable values.
argsAction="-R"
moveMode="Current Desktop"
searchMode="Class"
verbose=""

# Get the passed options.
while getopts ":achHLmnpRv" option; do
    case $option in
    a)
        argsAction="-a"
        moveMode="All Desktops"
        ;;

    c) searchMode="Title" ;;

    m)
        argsMove="-1"
        moveMode="Window move"
        ;;

    n)
        argsMove="1"
        moveMode="Window move"
        ;;

    p)
        argsAction="-R"
        moveMode="Current Desktop"
        ;;

    R)
        argsAction="-R"
        moveMode="All Desktops"
        ;;

    h)
        echo "$help"
        exit 0
        ;;
    H)
        printf "%s\n%s" "$help" "$longHelp"
        exit 0
        ;;
    L)
        printf "%s\n%s\n%s" "$license" "$help" "$longHelp"
        exit 0
        ;;
    v) verbose="true" ;;
    *)
        echo "Invalid option $*"
        exit
        ;;
    esac
done

# Remove parsed options.
shift $((OPTIND - 1))

# Set the launch command and title search.
command="$1"     # Window class search term and launch command.
windowTitle="$1" # Title bar search terms.
if [ "$2" ]; then
    # Use window title search if given.
    windowTitle="$2"
fi

# Determine the user's desktop number.
currentDesktop=" $(wmctrl -d | grep "\*" | cut -c 1)"
# Determine the user's active windowID.
currentWindow="$(printf '0x%08x\n' "$(xprop -root |
    grep "_NET_ACTIVE_WINDOW(WINDOW)" | grep -E -o "0x[0-9a-f]+")")"

# Preserve the currentDesktop and command terms.
desktopSearch="$currentDesktop" # For searching all desktops.
classSearch="$command"          # For classless searching.
if [ "$moveMode" = "All Desktops" ]; then
    # Search all desktops
    desktopSearch=".."
fi
if [ "$searchMode" = "Title" ]; then
    # Search only window title, ignore class.
    classSearch=".*"
fi
searchResults="$(wmctrl -lx | grep "0x........ $desktopSearch $classSearch\." |
    grep -i "$windowTitle")"

# Check if search match is the current window, If so use the next match.
windowID="$(echo "$searchResults" | grep -A1 "$currentWindow" | tail -1 |
    cut -c 1-10)"
if [ "$windowID" = "$currentWindow" ]; then
    # If the current window is the last result then use the first result.
    windowID="$(echo "$searchResults" | head -1 | cut -c 1-10)"
elif [ ! "$windowID" ]; then
    # If there is only one window then use that one.
    windowID="$(echo "$searchResults" | head -1 | cut -c 1-10)"
fi

# Raise the window or run command.
if [ "$windowID" ]; then
    # Activate the window that was found.
    wmctrl -i "$argsAction" "$windowID "
    time2="$(date +%s.%N)"
else
    # Check if kioclient exists for launch notifications.
    KDEcheck="$(type -t kioclient)"
    # Look for a .desktop file to use for launch notifications.
    desktopFile="$(find -L /usr/share/applications/ -name "$command".desktop |
        grep -0 "$command")"
    if [ "$desktopFile" != "" ] && [ "$KDEcheck" = "file" ]; then
        # Launch with notifications.
        time2="$(date +%s.%N)"
        kioclient exec "$desktopFile"
    else
        # Launch the program.
        time2="$(date +%s.%N)"
        $command &
    fi
fi

# Check for moving the active window to desktop left or right.
if [ "$moveMode" = "Window move" ]; then
    desktopCount="$(($(wmctrl -d | tail -1 | cut -c 1) + 1))"
    # Move ±1 from the current desktop mod the total desktops.
    finalDesktop="$(((argsMove + currentDesktop + desktopCount) % \
        desktopCount))"
    wmctrl -r :ACTIVE: -t "$finalDesktop"
    time2="$(date +%s.%N)"
fi

# Verbose comments moved down for speed.
if [ "$verbose" ]; then
    echo
    echo "Options unparsed: $*"
    echo "Search: $windowTitle"
    echo "Command: $command"
    echo "Current Desktop: $currentDesktop"
    echo "Move Mode: $moveMode"
    echo
    echo "Window List:"
    wmctrl -lx
    echo
    echo "Search Results:"
    echo "$searchResults"
    echo
    echo "Current Window: $currentWindow"
    echo "Move to Window: $windowID"
    echo
    if [ ! "$windowID" ]; then
        echo ".desktop File: $desktopFile"
        echo "KDE: no\b\b$([ "$KDEcheck" = "file" ] && echo 'yes')"
        echo
    fi
    echo "Time taken: 0$(echo "$time2"-"$time1" | bc) seconds"
fi

exit 0
