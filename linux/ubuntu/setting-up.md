# Applications

## AutoKey

```sh
sudo apt install autokey-gtk;
```

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

## Compiz Settings Manager

```bash
sudo apt install compiz-settings-manager compiz-plugins compiz-extra-plugins
```

## DConf Editor
```bash
sudo apt install dconf-editor;
```

## Docker
### Setup
```bash
sudo apt-get -y install apt-transport-https ca-certificates curl software-properties-common;
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -;
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable";
sudo apt-get update;
sudo apt-get -y install docker-ce python python-pip;
sudo systemctl enable docker;
sudo pip install docker-compose;
```

> Source: https://docs.docker.com/install/linux/docker-ce/ubuntu/#set-up-the-repository

### Running as non-root
Run:

```bash
sudo groupadd docker;
sudo usermod -aG docker $USER;
```

> You need to logout after this and log back in

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

## Golang

```sh
sudo add-apt-repository ppa:gophers/archive;
sudo apt update;
sudo apt-get install golang-1.10-go;
```

## Kubernetes

### `kubectl`
#### Via Download
```bash
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl;

chmod +x ./kubectl;

sudo mv ./kubectl /usr/local/bin/kubectl;
```

#### Via Package Manager
```
apt-get update && apt-get install -y apt-transport-https;
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -;
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF;
apt-get update;
apt-get install -y kubectl;
```

### `minikube`

```sh
MINIKUBE_VERSION="$(curl -sSL -X GET https://api.github.com/repos/kubernetes/minikube/releases/latest | jq -r '.tag_name')";
curl -Lo minikube "https://storage.googleapis.com/minikube/releases/${MINIKUBE_VERSION}/minikube-linux-amd64";
chmod +x minikube;
sudo mv minikube "/opt/minikube-${MINIKUBE_VERSION}";
sudo ln -s /opt/minikube-${MINIKUBE_VERSION} /usr/local/bin/minikube;
```

## Laptop Mode Tools
```bash
sudo apt install laptop-mode-tools;
```

To start it, run:

```sh
gksu lmt-config-gui;
```

## MySQL
### Client Only
```sh
sudo apt install mysql-client-core-5.7;
```
### Server
```
sudo apt install mysql-server;
sudo mysql_secure_installation;
```

### MySQL WorkBench
```bash
https://dev.mysql.com/downloads/workbench/
```

## No Notifications
For a *Do Not Disturb* mode.

```bash
sudo add-apt-repository ppa:vlijm/nonotifs;
sudo apt-get update && sudo apt-get install nonotifs;
```

## Node
```bash
sudo apt-get install -y build-essential;
NE=6
curl -sL https://deb.nodesource.com/setup_${NE}.x | sudo -E bash -
```
> Change the `NODE_VERSION` variable value to 6/7/8/9 *et cetera*.

Also add the path `${PWD}/node_modules/.bin` to your exported `PATH` environment variable so that you can run binaries from local repositories.

## Node Version Manager
```bash
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash
```

## Nvidia Drivers
> Note that Nvidia drivers are proprietary

```bash
sudo add-apt-repository ppa:graphics-drivers/ppa;
sudo apt update;
```

## Physical Hard-Disk Access

Go to VirtualBox GUI and create a new image without a hard drive.

Run `sudo fdisk -l` to retrieve the Windows partition.

Run `VBoxManage internalcommands createrawvmdk -filename "</path/to/file>.vmdk" -rawdisk /dev/sda` to create the VMDK.

## Postman

```bash
wget https://dl.pstmn.io/download/latest/linux64 -O postman.tar.gz;
sudo tar -xzf postman.tar.gz -C /opt;
rm postman.tar.gz;
sudo ln -s /opt/Postman/Postman /usr/bin/postman;
```

## PowerTop

```sh
sudo apt install powertop fancontrol;
```

## Selenium

```bash
rm ~/selenium-server-standalone-*.jar;
rm ~/chromedriver_linux64.zip;
sudo rm /usr/local/bin/chromedriver;
sudo rm /usr/local/bin/selenium-server-standalone.jar;
sudo apt-get update;
sudo apt-get install -y unzip openjdk-8-jre-headless xvfb libxi6 libgconf-2-4;
wget -N http://chromedriver.storage.googleapis.com/$(curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE)/chromedriver_linux64.zip -P ~/;
unzip ~/chromedriver_linux64.zip -d ~/;
rm ~/chromedriver_linux64.zip;
sudo mv -f ~/chromedriver /usr/local/bin/chromedriver;
sudo chown root:root /usr/local/bin/chromedriver;
sudo chmod 0755 /usr/local/bin/chromedriver;
wget -N http://selenium-release.storage.googleapis.com/3.4/selenium-server-standalone-3.4.0.jar -P ~/;
sudo mv -f ~/selenium-server-standalone-3.4.0.jar /usr/local/bin/selenium-server-standalone.jar;
sudo chown root:root /usr/local/bin/selenium-server-standalone.jar;
sudo chmod 0755 /usr/local/bin/selenium-server-standalone.jar;
```

## Redis

### Client Only

```sh
sudo apt install redis-tools
```

## Telegram
https://telegram.org

## Tilix

```bash
sudo add-apt-repository ppa:webupd8team/terminix;
sudo apt-get update;
sudo apt install tilix;
```

## Unity Tweak Tool
This replaces the GNOME tweak tool which doesn't work on Ubuntu.

```bash
sudo apt install unity-tweak-tool;
```

## Vault (HashiCorp)
https://www.vaultproject.io/downloads.html

## Visual Studio Code
https://code.visualstudio.com/docs/setup/linux#_visual-studio-code-is-unable-to-watch-for-file-changes-in-this-large-workspace-error-enospc

## Yarn
```bash
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -;
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list;
sudo apt update;
sudo apt install yarn;
```

# Others

## Remove Amazon Integration
sudo rm /usr/share/applications/ubuntu-amazon-default.desktop;
sudo rm /usr/share/unity-webapps/userscripts/unity-webapps-amazon;

## Keyboard Backlight
This setting is located at `/sys/class/leds/tpacpi::kbd_backlight/brightness`. Or you could use `Fn+Space`.

# Other Useful Reading
- [PowerManagement/ReducedPower](https://help.ubuntu.com/community/PowerManagement/ReducedPower)
