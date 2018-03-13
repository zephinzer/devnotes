

# Applications
## Chromium
```bash
sudo apt install chromium-browser;
```

## DConf Editor
```bash
sudo apt install dconf-editor;
```

## Gnome Keyring
https://wiki.archlinux.org/index.php/GNOME/Keyring

## MySQL WorkBench
```bash
https://dev.mysql.com/downloads/workbench/
```

## Node
```bash
NE=6
curl -sL https://deb.nodesource.com/setup_${NE}.x | sudo -E bash -
```
> Change the `NODE_VERSION` variable value to 6/7/8/9 *et cetera*.

## Node Version Manager
```bash
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.8/install.sh | bash
```

## Telegram
https://telegram.org

## Visual Studio Code
https://code.visualstudio.com/docs/setup/linux#_visual-studio-code-is-unable-to-watch-for-file-changes-in-this-large-workspace-error-enospc

## Yarn
```bash
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -;
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list;
```

## Docker
https://docs.docker.com/install/linux/docker-ce/ubuntu/#set-up-the-repository

### Running as non-root
Run:

```bash
sudo groupadd docker;
sudo usermod -aG docker $USER;
```

Source: https://docs.docker.com/install/linux/linux-postinstall/

## Gnome Network Manager
- this helps with connecting to a VPN using a GUI

```bash
sudo apt install network-manager-openvpn-gnome;
```

## Gnome Tweak Tool
```bash
sudo apt install gnome-tweak-tool;
```

## Laptop Mode Tools
```bash
sudo apt install laptop-mode-tools;
```

# Other Useful Reading
- [PowerManagement/ReducedPower](https://help.ubuntu.com/community/PowerManagement/ReducedPower)
