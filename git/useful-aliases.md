# Useful Aliases for Git
Add the following to your `~/.profile` file.

```bash
alias ga='git add';
alias gcm='git commit -m';
alias gcmae='git commit --allow-empty -m';
alias gco='git checkout';
alias gcob='git cehckout -b';
alias gf='git fetch';
alias gi='git init';
alias gl='git log';
alias gln='git log -n';
alias gplrecb='git pull --rebase origin $(git branch | grep '*' | cut -f 2 -d ' ')';
alias gplcb="git pull origin $(git branch | grep '*' | cut -f 2 -d ' ')";
alias gpscb="git push origin $(git branch | grep '*' | cut -f 2 -d ' ')";
alias gr='git rebase';
alias gri='git rebase -i HEAD~';
alias gs='git status';
```