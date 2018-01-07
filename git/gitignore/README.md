# .gitignore

## Blacklist by extension

> File:[`.gitignore.by.blacklist`](.gitignore.by.blacklist)

Ignores all files given a set of extensions. In this context, it ignores all markdown (`*.md`) files.

```
*.md
```

## Whitelist by extension

> File:[`.gitignore.by.whitelist`](.gitignore.by.whitelist)

Ignores all files except what is defined. In this context, it ignores all file extensions except markdown (`*.md`) files.

```
*
!/**/
!*.md
```