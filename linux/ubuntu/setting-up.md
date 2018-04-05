# Applications
## Chromium
Because screw Chrome.

```bash
sudo apt install chromium-browser;
```

## Chrome Gnome Shell
For installing extensions

```bash
sudo apt-get install chrome-gnome-shell
```

## DConf Editor
```bash
sudo apt install dconf-editor;
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

## Emoji Support

```bash
sudo apt-add-repository ppa:eosrei/fonts;
sudo apt-get update;
sudo apt-get install fonts-twemoji-svginot;
```

## File System Support

### EX-FAT

```bash
sudo apt-get install exfat-utils exfat-fuse;
```

## Kubernetes (`kubectl`)
```bash
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl;

chmod +x ./kubectl;

sudo mv ./kubectl /usr/local/bin/kubectl
```

## Laptop Mode Tools
```bash
sudo apt install laptop-mode-tools;
```

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

## Nvidia Drivers
> Note that Nvidia drivers are proprietary

```bash
sudo add-apt-repository ppa:graphics-drivers/ppa;
sudo apt update;
```

## Telegram
https://telegram.org

## Tilix
```bash
sudo apt install tilix;
```

## Unity Tweak Tool
This replaces the GNOME tweak tool which doesn't work on Ubuntu.

```bash
sudo apt install unity-tweak-tool;
```

## Visual Studio Code
https://code.visualstudio.com/docs/setup/linux#_visual-studio-code-is-unable-to-watch-for-file-changes-in-this-large-workspace-error-enospc

## Yarn
```bash
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -;
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list;
```

# Other Useful Reading
- [PowerManagement/ReducedPower](https://help.ubuntu.com/community/PowerManagement/ReducedPower)
