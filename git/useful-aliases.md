# Useful Aliases for Git
Add the following to your `~/.profile` file.

```bash
alias ga='git add';
alias gcm='git commit -m';
alias gf='git fetch';
alias gplrecb='git pull --rebase origin $(git branch | grep '*' | cut -f 2 -d ' ')';
alias gplcb='git pull origin $(git branch | grep '*' | cut -f 2 -d ' ')';
alias gpscb='git push origin $(git branch | grep '*' | cut -f 2 -d ' ')';
alias gr='git rebase';
```