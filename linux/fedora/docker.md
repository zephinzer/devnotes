# Docker Issues
## Docker as Root
To enable Docker to run on Fedora:

```
sudo groupadd docker && sudo gpasswd -a ${USER} docker && sudo systemctl restart docker && newgrp docker
```

**References**
- https://developer.fedoraproject.org/tools/docker/docker-installation.html

## No permissions on mounted volume
To get permissions on a mounted volume:

```
docker run ... -v "$(pwd):/app/...:Z"
```

> Note the use of `Z` to mount.

- https://docs.docker.com/storage/bind-mounts/#configure-the-selinux-label
**References**