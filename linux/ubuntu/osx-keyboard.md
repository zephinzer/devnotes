# Keybindings for Mac OS X Users

## Global Settings
### AutoKey Phrases
- CursorToBottom - `<ctrl>+<down>` = `<ctrl>+<end>`
- CursorToEnd - `<ctrl>+<right>` = `<end>`
- CursorToNext - `<super>+<right>` = `<ctrl>+<right>`
- CursorToPrevious - `<super>+<left>` = `<ctrl>+<left>`
- CursorToStart - `<ctrl>+<left>` = `<home>`
- CursorToTop - `<ctrl>+<up>` = `<ctrl>+<home>`
- SelectTillBottom - `<ctrl>+<shift>+<down>` = `<ctrl>+<shift>+<end>`
- SelectTillEnd - `<ctrl>+<shift>+<right>` = `<shift>+<end>`
- SelectTillNext - `<shift>+<super>+<right>` = `<ctrl>+<shift>+<right>`
- SelectTillPrevious - `<shift>+<super>+<left>` = `<ctrl>+<shift>+<left>`
- SelectTillStart - `<ctrl>+<shift>+<left>` = `<shift>+<home>`
- SelectTillTop - `<ctrl>+<shift>+<up>` = `<ctrl>+<shift>+<home>`
- TabNavNext - `<ctrl>+<shift>+]` = `<ctrl>+<page_down>`
- TabNavPrevious - `<ctrl>+<shift>+[` = `<ctrl>+<page_up>`

### Disable Shortcut Hints Overlay
#### Install Compiz Configuration Settings Manager
Install it using:

```bash
sudo apt install compizconfig-settings-manager;
```
#### Disable Shortcut Hints
Open Compiz Configuration Settings Manager and go to the **Desktop** section. Click on **Ubuntu Unity Plugin** and go to the **General** navigation tab. Unselect the option *Enable Shortcut Hints Overlay*.

### Swap the Alt/Ctrl Keys
#### Using XModMap
Create a file at `~/.Xmodmap` with the following contents:

```
clear control
clear mod1
keycode 37 = Alt_L Meta_L
keycode 64 = Control_L
add control = Control_L Control_R
add mod1 = Alt_L Meta_L
```

Next run the command to let the settings take effect:

```bash
xmodmap ~./.Xmodmap;
```

### Configure Application Switcher
By default, Ubuntu uses `Alt+Tab` to switch applications. We configure this to use `Ctrl+Tab` so that together with the `Xmodmap` configuration, the application switcher works as expected.

#### Using Compiz Configuration Settings Manager
Install it using:

```bash
sudo apt install compizconfig-settings-manager;
```

#### Setting Application Switcher Keybinding
Open the Compiz Configuration Settings Manager and go to Ubuntu Unity Plugin, scroll to **Desktops** and click on the **Switcher** top navigation tab. Set the keybinding for **Key to start the Switcher** to `<Control><Tab>`. This allo  ws us to use the re-mapped `Alt` hardware key to switch through tabs, like how OS X does it.

#### Setting Window Switcher Keybinding
> This engages a known bug in Unity Ubuntu Plugin where even though the setting shows *Disabled*, it's actually enabled. The solution is just to override the keybinding to something that doesn't get in the way.

Open the Compiz Configuration Settings Manager and go to Ubuntu Unity Plugin, scroll to **Desktops** and click on the **Switcher** top navigation tab. Set the keybinding for the **Key to flip through windows in the switcher** and **Key to flip through windows in the switcher backwards** to `<Control><Alt>grave` and `<Control><Shift><Alt>asciitilde` respectively. (`grave`) is the backquote, (`asciitilde`) is your tilda.

###### References
- Error in Unity Ubuntu Plugin: https://askubuntu.com/questions/168265/how-do-i-disable-or-change-the-alt-shortcut-to-switch-between-multiple-win

### Enable the Static Application Switcher

#### Install Compiz Configuration Settings Manager
Install it using:

```bash
sudo apt install compizconfig-settings-manager;
```

#### Enable the Static Application Switcher
Open the Compiz Configuration Settings Manager, scroll to **Windows Management** and enable **Static Application Switcher**. You should encounter some conflicts, for most of the cases, disable the keybindings from the old **Application Switcher**

## Application Settings
### Terminal
Open the Terminal and go to the preferences from the menu bar (**Terminal > Preferences**).

Click on the **Shortcuts** tab menu item and:

Scroll to the **File** section. Change:
- **Close Terminal** -> `Ctrl+W`
- **Close All Terminals** -> `Ctrl+Q`

Scroll down to the **Edit** section. Change:
- **Copy** -> `Ctrl+C`
- **Paste** -> `Ctrl+V`

> Note that this automatically re-assigns your `tty` interrupt to `Ctrl+Alt+C` but since we use that much less than copy and paste (:

### Tilix
Tilix is a nice replacement for iTerm2. However, some shortcuts are different. Get Tilix with:

```bash
sudo apt install tilix;
```

Click on the menu bar on the Tilix icon and click on **Preferences**.

Go to the **Shortcuts** menu bar item and set the following shortcuts so that the keyboard bindings are closer to iTerm2:
- **Application/Open a new session** - `Ctrl+T`
- **Application/Open preferences** - `Ctrl+,`
- **Session/Add terminal down** - `Ctrl+Shift+D`
- **Session/Add terminal right** - `Ctrl+D`
- **Session/Edit the session name** - `Ctrl+Shift+Return`
- **Terminal/Close terminal** - `Ctrl+W`
- **Terminal/Copy** - `Ctrl+C`
- **Terminal/Paste** - `Ctrl+V`

### Visual Studio Code
#### Install Sublime Keybindings
Search the themes for sublime and install the Sublime keybindings from the extension **Sublime Text Keymap**.

#### Cursor Keybindings
Open the keyboard shortcuts and make the following changes:

- `editor.action.copyLinesDownAction` -> `Shift+Windows+DownArrow`
- `editor.action.insertCursorAbove` -> `Ctrl+Shift+Windows+UpArrow` *missing G from OS X*
- `editor.action.insertCursorBelow` -> `Ctrl+Shift+Windows+DownArrow` *missing G from OS X*
- `editor.action.copyLinesUpAction` -> `Shift+Windows+UpArrow`
- `editor.action.moveLinesDownAction` -> `Windows+DownArrow`
- `editor.action.moveLinesUpAction` -> `Windows+UpArrow`
- `renameFile` -> `Shift+Enter` *Enter on OS X*
- `deleteFile` -> `Shift+Backspace` *Backspace on OS X*
- `workbench.action.nextEditor` -> `Ctrl+PageDown`
- `workbench.action.previousEditor` -> `Ctrl+PageUp`
- `workbench.action.terminal.toggleTerminal` -> `Alt+\``
- `workbench.action.terminal.copySelection` -> `Ctrl+C`
- `workbench.action.terminal.paste` -> `Ctrl+V`

> Key mappings with a * notation indicate where the keymapping is not exactly according to OS X