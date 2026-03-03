---
name: sync-workspace
description: Pull, commit, and push all workspace changes to the git remote with automatic conflict resolution. Use when the user asks to sync, push, or save changes to remote.
---

# Sync Workspace

Sync local changes with remote via device-specific branches and main.

## Branch Model

- Each device works on its own branch (e.g. `colin-mbp15-2018`, `colin-hs-mbp2023`).
- `main` is the shared trunk — all devices converge through it.
- On sync: rebase device branch onto latest `origin/main`, then push to both.

## Workflow

1. **Check state**: `git branch --show-current` and `git status`
2. **If on `main`**: simple pull/push (skip rebase flow)
3. **If local changes exist**: stage and commit with a descriptive message
4. **Fetch remote**: `git fetch origin`
5. **Rebase onto main**: `git rebase origin/main`
6. If rebase conflicts arise, resolve them (see below), then `git rebase --continue`
7. **Push device branch**: `git push --force-with-lease origin <device-branch>`
8. **Push to main (fast-forward)**: `git push origin HEAD:main`
9. Report what happened

## Conflict Resolution During Rebase

When `git rebase origin/main` reports conflicts:

1. List conflicted files: `git diff --name-only --diff-filter=U`
2. Read each file, examine both versions (`<<<<<<<`, `=======`, `>>>>>>>`)
3. **Strategy**: Keep the best content from both sides — prefer more complete/recent content, combine non-overlapping additions
4. Remove conflict markers, stage: `git add <file>`
5. Continue rebase: `git rebase --continue`
6. If unresolvable, abort: `git rebase --abort` and report to user

## Fallback: On `main` Branch

If the current branch is `main` (no device branch):

1. Stage and commit local changes
2. `git pull --rebase origin main`
3. `git push origin main`

## Commit Message Guidelines

- Summarize what changed, not how
- Imperative mood: "Add", "Update", "Fix", "Remove"
- Under 72 characters for subject line
- Multiple unrelated changes: bullet points in body

## Abort Conditions

Do NOT proceed if:
- Working tree is clean AND already up to date with remote
- User explicitly cancels
- Push to main fails with non-fast-forward (re-fetch and retry once, then report)
