# .gitattributes reasoning

## What was done

A repository-level `.gitattributes` file was added at the root of `CoffieAPI`.

Current rules:

```gitattributes
*.go text eol=lf
*.mod text eol=lf
*.sum text eol=lf
*.md text eol=lf
*.yml text eol=lf
*.yaml text eol=lf
*.sql text eol=lf
Dockerfile text eol=lf
Dockerfile.* text eol=lf
Makefile text eol=lf
*.sh text eol=lf
```

## Why this was necessary

The repository did not have a `.gitattributes` file, so line-ending behavior was being inherited from the machine-wide Git configuration.

On this machine:

- `core.autocrlf=true`
- the setting came from the system Git config

That meant Git was free to apply Windows-style `CRLF` conversion in the working tree even for source files that are better kept as `LF`.

This is what triggered the warning:

```text
warning: in the working copy of 'tests/feature/user/user_store_test.go', LF will be replaced by CRLF the next time Git touches it
```

The warning was not a code bug, but it was a sign that the repository had no explicit line-ending policy of its own.

## Why a repo-level policy is better

Adding `.gitattributes` makes the repository responsible for line endings instead of leaving the decision to each developer's Git installation.

Benefits:

- consistent behavior across Windows, macOS, and Linux
- fewer noisy diffs caused only by line-ending changes
- fewer confusing warnings during add/commit operations
- better alignment with common Go-project expectations, where source files are typically kept as `LF`

In other words, this change makes line endings part of the project's rules rather than part of each developer's local machine setup.

## Why `LF` was chosen

`LF` is the most portable default for source code and config files in this project.

It is a good fit here because:

- Go tooling is naturally comfortable with `LF`
- Docker-related files are usually cleaner and more predictable with `LF`
- YAML, SQL, Markdown, and shell-oriented files are commonly standardized on `LF`
- it avoids cross-platform churn when contributors work on different operating systems

This does not mean Windows cannot work with the repo. It means Git now knows that these tracked files should stay normalized as `LF`.

## What was verified

After adding `.gitattributes`, Git reported the new policy for representative files such as:

- `README.md`
- `cmd/server/main.go`
- `docker-compose.yml`
- `tests/feature/user/user_store_test.go`

Git now shows these files with `attr/text eol=lf`, which confirms that the repository policy is being applied.

## Practical outcome

The immediate goal was not to change application behavior. The goal was to remove ambiguity and prevent future line-ending churn.

This change reduces the chance of:

- accidental CRLF rewrites in tracked source files
- environment-specific warnings that confuse contributors
- commits polluted by formatting-only line-ending changes

## Follow-up note

If a fully normalized working tree is needed immediately, files can be re-checked out so Git rewrites them under the new policy. The important part is that the repository now has an explicit rule going forward.
