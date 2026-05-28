# `leap push` — quick examples

A tour of the new `leap push` ergonomics. The legacy spellings still
work (back-compat); these are the shorter / nicer paths.

## Push a new version

```bash
leap push                       # interactive: asks name, then model file
leap push -n my-model -m model.h5  # non-interactive
leap push -e                    # push, then evaluate
leap push -e --batch 32         # eval with a specific batch size
```

The interactive flow now asks for the **version name first**, then the
model file — matches the way you think about a push: "what am I
calling it → which file is it".

## Overwrite an existing version

```bash
leap push -o 6a16a0cf            # by id
leap push -o my-model            # by name (picker opens if ambiguous)
```

The CLI echoes the resolved target before doing anything destructive:

```
Overwriting my-model (id 6a16a0cf, matched "my-model")
```

If the name matches multiple versions, an interactive picker shows
each with its id + status:

```
? Multiple versions named "model" — pick one
  ▸ #02  model  6a16a0cf  evaluated 2m ago
    #01  model  6a16c86c  evaluated 3d ago
```

In a non-TTY context (CI) the picker is replaced by an error that
lists the candidate ids so you can re-run with the exact one.

## Overwrite + run an evaluation

```bash
leap push -o my-model -e
```

You'll get a multi-select prompt for what you changed (only edits the
engine can't auto-detect):

```
? What did you change in your code? (auto-detected: added meta/viz, metric-direction, insight-config)
  [Use arrows to move, space to select, type to filter]
  [ ] Edited an existing metadata
  [ ] Edited an existing visualization
  [ ] Added or edited a metric
  [ ] Force re-run insights
```

Then a plan summary:

```
This will:
  • Update visualizations
  • Auto-detect added meta/viz, metric-direction, insight-config
```

## Skip the prompt — name the artifacts explicitly

`--update` (`-u`) tells `leap push` exactly which artifacts to refresh.
It **implies `--eval`** — you no longer need both flags.

```bash
leap push -o my-model -u viz                       # regenerate visualizations only
leap push -o my-model -u metadata -u insights      # refresh both
leap push -o my-model -u metric                    # full re-evaluation triggered by metric change
```

Accepted values (case-insensitive):

| short          | full (still accepted) |
| -------------- | --------------------- |
| `metadata`     | `update_metadata`     |
| `metric`       | `update_metric`       |
| `insights`     | `update_insights`     |
| `visualization` / `viz` | `update_visualization` |

## Overwrite without evaluating

```bash
leap push -o my-model            # no --eval, no --update
```

When you overwrite without `--eval`, the CLI reminds you that
evaluation isn't automatic:

```
NOTE: Overwriting replaces the version. Use the update-evaluate dialog
      in the UI to re-evaluate (or re-run with --eval).
```

The UI's update-evaluate dialog lets you drive the same diff-based
refresh interactively, post-push.

## Deprecated flag still works

```bash
leap push --overwrite-version 6a16a0cf -e
```

→ prints `Flag --overwrite-version has been deprecated, use --overwrite (-o) instead`
and proceeds normally. Update your scripts at your leisure.

## Cheat sheet

| goal                                             | command                                       |
| ------------------------------------------------ | --------------------------------------------- |
| Push a new version                               | `leap push`                                   |
| Push + eval                                      | `leap push -e`                                |
| Overwrite by id                                  | `leap push -o <id>`                           |
| Overwrite by name                                | `leap push -o <name>`                         |
| Overwrite + force-refresh viz                    | `leap push -o <name> -u viz`                  |
| Overwrite + refresh metadata and insights         | `leap push -o <name> -u metadata -u insights` |
| Overwrite + full re-eval (auto-detect)            | `leap push -o <name> -e`                      |
