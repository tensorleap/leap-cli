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

You'll get a multi-select prompt for what you changed — engine no longer
auto-detects, so you must pick what's true:

```
? What did you change in your code?
  [Use arrows to move, space to select, type to filter]
  [ ] Added or edited a metadata           — triggers full re-evaluation
  [ ] Added or edited a metric             — triggers full re-evaluation
  [ ] Edited metric direction or insight-config — cheap update
  [ ] Added or edited a visualization      — regenerate all visualizations
  [ ] Reinforce insights                   — regenerate insights from scratch
```

Then a plan summary:

```
This will:
  • Regenerate visualizations
```

If you pick metadata or metric, the plan collapses to:

```
This will:
  • Re-evaluate (full)
```

## Skip the prompt — name the artifacts explicitly

`--update` (`-u`) tells `leap push` exactly what changed in the code.
It **implies `--eval`** — you no longer need both flags.

```bash
leap push -o my-model -u viz                          # regenerate visualizations only
leap push -o my-model -u metric_direction             # cheap metric-direction patch
leap push -o my-model -u metric                       # triggers a full re-evaluation
leap push -o my-model -u metadata                     # triggers a full re-evaluation
leap push -o my-model -u visualization -u insights    # refresh both
```

Accepted values (case-insensitive):

| short                    | full (still accepted)     |
| ------------------------ | ------------------------- |
| `metadata`               | `update_metadata`         |
| `metric`                 | `update_metric`           |
| `metric_direction` / `direction` | `update_metric_direction` |
| `insights`               | `update_insights`         |
| `visualization` / `viz`  | `update_visualization`    |

`metadata` and `metric` always trigger a fresh evaluation —
update_evaluate doesn't support them today.

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

The UI's update-evaluate dialog walks the same five-action prompt as
the CLI's `--eval` flow.

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
| Overwrite + cheap metric-direction patch          | `leap push -o <name> -u metric_direction`     |
| Overwrite + interactive prompt for what changed   | `leap push -o <name> -e`                      |
| Overwrite + force full re-eval                    | `leap push -o <name> -u metric`               |
