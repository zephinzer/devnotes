# Node Issues
## ENOSPC on attempted `--watch` flag
This issue arises when there are not enough system watchers. To increase the watcher limit:

```bash
echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p
```