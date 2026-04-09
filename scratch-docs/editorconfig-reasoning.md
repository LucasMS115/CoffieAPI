# .editorconfig reasoning

## What was done

A repository-level `.editorconfig` file was added at the root of `CoffieAPI`.

Current rules:

```editorconfig
root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true

[*.md]
trim_trailing_whitespace = false
```

## Why this was necessary

After adding `.gitattributes`, Git correctly knew that tracked source and config files should use `LF` line endings.

However, one warning still appeared for `tests/feature/user/integration/user_register_test.go`:

```text
warning: in the working copy of 'tests/feature/user/integration/user_register_test.go', CRLF will be replaced by LF the next time Git touches it
```

That warning showed that the repository policy and the editor behavior were still out of sync.

Git was reporting:

- index: `LF`
- working tree: `CRLF`
- repo attribute: `text eol=lf`

So the Git side was already correct, but the file was still being written to disk with `CRLF` by the editor.

## Why `.gitattributes` was not enough on its own

`.gitattributes` controls how Git interprets and normalizes files.
It does not guarantee how an editor will save a file before Git sees it.

That means a repository can still get repetitive line-ending warnings if:

- Git expects `LF`
- the editor keeps saving the working copy as `CRLF`

In that situation, Git is doing the right thing, but it still has to warn because the working tree does not match the repo policy.

## Why `.editorconfig` is the right complement

`.editorconfig` gives the editor a repository-level formatting policy.

In this case it tells the editor to:

- save files with `LF`
- use UTF-8
- ensure a final newline exists
- trim trailing whitespace by default
- avoid trimming Markdown trailing whitespace automatically

This complements `.gitattributes` well:

- `.gitattributes` defines what the repository expects
- `.editorconfig` defines how editors should save files so they already match that expectation

Together they reduce the chance of Git and the editor fighting over line endings.

## Why `LF` is still the right choice here

The same reasoning used for `.gitattributes` applies here too:

- Go source files are normally kept with `LF`
- Docker, YAML, SQL, and shell-related files are more predictable with `LF`
- cross-platform repositories are easier to keep clean when source files are normalized the same way everywhere

This is not about preventing Windows usage. It is about making the editor save files in the format the repository already declared as correct.

## Practical outcome

Adding `.editorconfig` should reduce the remaining line-ending warnings for files edited in tools that respect EditorConfig.

It also improves consistency beyond line endings by making other basic formatting behavior explicit.

## Important follow-up note

Adding `.editorconfig` does not instantly rewrite files that are already open or already saved with `CRLF`.

A file like `tests/feature/user/integration/user_register_test.go` may still need one fresh save after the editor picks up the new settings.

So the intended steady state is:

- Git expects `LF`
- the editor saves `LF`
- the warning stops appearing for newly edited files
