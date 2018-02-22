# Useful Aliases for Kubernetes
The following aliases are concatenations of the following shorthands:

```
a      -> apply
d      -> describe
cm     -> configmap
del    -> delete
ep     -> endpoints
exec   -> exec
g      -> get
i      -> ingress
kc     -> kubectl
l      -> logs
n      -> nodes
ns     -> namespace
p      -> pod
pv     -> persistentvolume
pvc    -> persistentvolumeclaim
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
alias kcdcm='kubectl describe configmap';
alias kcdd='kubectl describe deployment';
alias kcdep='kubectl describe endpoints';
alias kcdi='kubectl describe ingress';
alias kcdns='kubectl describe namespaces';
alias kcdn='kubectl describe node';
alias kcdp='kubectl describe pod';
alias kcdpv='kubectl describe persistentvolume';
alias kcdpvc='kubectl describe persistentvolumeclaim';
alias kcdsec='kubectl describe secret';
alias kcds='kubectl describe service';
alias kcdel='kubectl delete';
alias kcdelcm='kubectl delete configmap';
alias kcdeld='kubectl delete deployment';
alias kcdelep='kubectl delete endpoints';
alias kcdeli='kubectl delete ingress';
alias kcdelp='kubectl delete pod';
alias kcdelpv='kubectl delete persistentvolume';
alias kcdelpvc='kubectl delete persistentvolumeclaim';
alias kcdelsec='kubectl delete secret';
alias kcdels='kubectl delete service';
alias kcg='kubectl get';
alias kcgcm='kubectl get configmap -o wide';
alias kcgd='kubectl get deployment -o wide';
alias kcgep='kubectl get endpoints -o wide';
alias kcgi='kubectl get ingress -o wide';
alias kcgns='kubectl get namespaces -o wide';
alias kcgn='kubectl get nodes -o wide';
alias kcgp='kubectl get pods -o wide';
alias kcgpv='kubectl get persistentvolumes -o wide';
alias kcgpvc='kubectl get persistentvolumeclaims -o wide';
alias kcgs='kubectl get services -o wide';
alias kcgsec='kubectl get secrets -o wide';
alias kcl='kubectl logs';
alias kclf='kubectl logs -f';
alias kcexec='kubectl exec';
```
