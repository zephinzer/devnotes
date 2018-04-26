# Useful Aliases for Shell Scripts

```bash
alias sshwp='ssh -vv -o PreferredAuthentications=password -o PubkeyAuthentication=no';
alias ll='ls -lA';
command -v xclip &>/dev/null && alias pbcopy='xclip -selection clipboard';
``` 