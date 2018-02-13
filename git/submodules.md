# Git Submodules

## Removing a Git Submodule

```bash
#!/bin/sh
SUBMODULE=$1;
git rm "${SUBMODULE}";
rm -rf ".git/modules/${SUBMODULE}";
git config -f .git/config --remove-section "submodule.${SUBMODULE}";
```

> Credits: https://stackoverflow.com/a/21211232