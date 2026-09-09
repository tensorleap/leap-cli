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

You'll get a multi-select prompt for what to update — engine no longer
auto-detects, so you pick what to refresh:

```
? What do you want to update?
  [Use arrows to move, space to select, type to filter]
  [ ] Metadata — add new metadata
  [ ] Metric — full re-eval
  [ ] Metric config — metric direction + insights
  [ ] Visualizations
  [ ] Samples — evaluate newly added samples only
```

Then a plan summary:

```
This will:
  • Regenerate visualizations
```

If you pick metric, the plan collapses to:

```
This will:
  • Re-evaluate (full)
```

## Skip the prompt — name the artifacts explicitly

`--update` (`-u`) tells `leap push` exactly what changed in the code.
It **implies `--eval`** — you no longer need both flags.

```bash
leap push -o my-model -u viz                          # regenerate visualizations only
leap push -o my-model -u metric_config                # metric direction + insights
leap push -o my-model -u metric                       # triggers a full re-evaluation
leap push -o my-model -u metadata                     # add new metadata
leap push -o my-model -u samples                      # evaluate only the newly added samples
leap push -o my-model -u samples -u metadata          # both, in one continue-evaluate
```

Accepted values (case-insensitive):

| short                    | full (still accepted)     |
| ------------------------ | ------------------------- |
| `metadata`               | `update_metadata`         |
| `metric`                 | `update_metric`           |
| `metric_config` / `metric-config` | `update_metric_config` |
| `visualization` / `viz`  | `update_visualization`    |
| `samples`                | `update_samples`          |

`metric` always triggers a fresh evaluation. `samples` runs a
continue-evaluate: only sample ids that are new since the last evaluation
are evaluated and appended to the existing results (any dataset state —
existing ids must stay stable; removing samples is rejected).

## Overwrite without evaluating

```bash
leap push -o my-model            # no --eval, no --update
```

You still get the multi-select prompt — the answer is recorded on the
version so the next time someone (UI, another CLI run) runs an
evaluation against this version they see your declared intent.
Picking **Samples** is the exception: it always runs the continue-evaluate
(there is nothing to record without evaluating the new samples), so the
"Run update-evaluate after push?" question is skipped.

```
NOTE: Overwriting replaces the version. Use the update-evaluate dialog
      in the UI to re-evaluate (or re-run with --eval).
Recorded update actions on version (no evaluation dispatched).
```

The UI's update-evaluate dialog uses the same five-action prompt as
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
| Overwrite + metric config patch                   | `leap push -o <name> -u metric_config`        |
| Overwrite + evaluate only newly added samples     | `leap push -o <name> -u samples`              |
| Overwrite + interactive prompt for what changed   | `leap push -o <name> -e`                      |
| Overwrite + force full re-eval                    | `leap push -o <name> -u metric`               |
