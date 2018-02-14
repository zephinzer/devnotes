# Useful Aliases for Git
Add the following to your `~/.profile` file.

```bash
alias ga='git add';
alias gcb='$(git branch | grep '*' | cut -f 2 -d ' ')';
alias gcm='git commit -m';
alias gcmae='git commit --allow-empty -m';
alias gco='git checkout';
alias gcob='gco -b';
alias gf='git fetch';
alias gi='git init';
alias gl='git log';
alias gln='git log -n';
alias gplrecb='git pull --rebase origin $(gcb)';
alias gplcb="git pull origin $(gcb)";
alias gpscb="git push origin $(gcb)";
alias gr='git rebase';
alias gri='git rebase -i HEAD~';
alias gs='git status';
```