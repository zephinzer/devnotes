# Git
## Useful Commands
### Search for a string in the commit history across all time

```
SEARCH_STRING='?';
git rev-list --all | (
  while read revision; do
    git grep -F "${SEARCH_STRING}" $revision;
  done;
);
```

> Credits: https://stackoverflow.com/a/4468394

### Search for a string in the code base across all time

```
SEARCH_STRING='?';
git rev-list --all | GIT_PAGER=cat xargs git grep '${SEARCH_STRING}'
```

> Credits: https://stackoverflow.com/a/16899941