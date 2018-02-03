# Retrieving the current `kubectl` context

> Requires `jq` tool to be installed (on OS X: `brew install jq`)

```bash
$ kubectl config view --output=json | jq '."current-context"'
```