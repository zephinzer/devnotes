# Troubleshooting Ubuntu

## `tput` requires `$TERM`
I had this issue in 16.04.4 after installing my own custom `~/.profile` which printed debug information on software versions and loaded paths and then ran a `tput reset` command to clear the terminal.

The solution is to use `tput -Txterm reset` instead. The `-T` flag specifies the terminal when the `$TERM` is not defined yet on Unity starting.

## Missing Unity
I had this issue in 16.0.4.4 after installing my own custom `~/.profile` which added my custom paths. The issue was forgetting to include certain paths. This resulted in the Ubuntu desktop showing up but no launcher/Unity being loaded.

When defining your own set of paths, always start with the system defined ones.

A bad example:

```sh
...
export PATH="/=/bin:/usr/bin/...";
...
```

A working example:

```sh
...
export PATH="${PATH}:/=/bin:/usr/bin/...";
...
```

### Related reading
- https://itsfoss.com/how-to-fix-no-unity-no-launcher-no-dash-in-ubuntu-12-10-quick-tip/
- https://askubuntu.com/questions/17381/unity-doesnt-load-no-launcher-no-dash-appears

## Nothing loads except the main Ubuntu
This is related to the Missing Unity issue. To debug this, switch to another TTY by using `ctrl+alt+f[1-6]`. To open unity manually, run:

```sh
DISPLAY=:0 unity &
```