on run argv
    tell application "iTerm2"
      tell current session of current window
        set background image to item 1 of argv
      end tell
    end tell
end run
