# Useful Aliases
The following aliases are concatenations of the following shorthands:

```
a      -> apply
d      -> describe
cm     -> configmap
del    -> delete
exec   -> exec
g      -> get
i      -> ingress
kc     -> kubectl
l      -> logs
n      -> nodes
p      -> pod
s      -> service
sec    -> secret
```

The lengths and abbreviations were decided based on how frequently they seemed to be used by me and they seem reasonable. Higher impact verbs such as `delete` are given a slightly longer abbreviation - `del` - so that one doesn't accidentally trigger them.

Include the following inside your `~/.profile` or specific `.[?]shrc` file:

```sh
alias kc='kubectl';
alias kca='kubectl apply -f';
alias kccv='kubectl config view';
alias kcd='kubectl describe';
alias kcdp='kubectl describe pod';
alias kcdn='kubectl describe node';
alias kcdd='kubectl describe deployment';
alias kcds='kubectl describe service';
alias kcdsec='kubectl describe sercret';
alias kcdi='kubectl describe ingress';
alias kcdcm='kubectl describe configmap';
alias kcdel='kubectl delete';
alias kcdelcm='kubectl delete configmap';
alias kcdeld='kubectl delete deployment';
alias kcdelp='kubectl delete pod';
alias kcdels='kubectl delete service';
alias kcdeli='kubectl delete ingress';
alias kcg='kubectl get';
alias kcgp='kubectl get pods -o wide';
alias kcgn='kubectl get nodes -o wide';
alias kcgep='kubectl get ep -o wide';
alias kcgd='kubectl get deploy -o wide';
alias kcgs='kubectl get svc -o wide';
alias kcgsec='kubectl get secrets -o wide';
alias kcgi='kubectl get ing -o wide';
alias kcgcm='kubectl get configmap -o wide';
alias kcl='kubectl logs';
alias kclf='kubectl logs -f';
alias kcexec='kubectl exec';
```
