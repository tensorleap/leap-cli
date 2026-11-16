# leap-cli

## Tensorleap skills run this CLI

The skills in [tensorleap/skills](https://github.com/tensorleap/skills) run `leap` and read its output: `cmd/auth`, `cmd/server`, `cmd/run`, `cmd/root_cmd/push.go`, `cmd/mcp`, `cmd/version.go`, `pkg/mcp` and `pkg/version/compat.go`. Before you open a PR that changes code there:

1. List every name your diff removes or renames: commands, flags, JSON fields, the text lines a command prints, the local MCP tools and their fields. A skill breaks when it reads a name that is gone.
2. Search tensorleap/skills for each name, in a local checkout or with `gh search code --repo tensorleap/skills <name>`. Also think about changes a search can't find: a value whose meaning or format changed.
3. If a skill uses any of them, open a skills PR that matches the change, and ship it with this CLI (see below). Skills always match the CLI they ship with, so there is no CLI level to bump.
4. Add a `Skills:` line to the PR description: the skills PR link, or `Skills: n/a, searched for <names>` when no skill uses them.

The Skills check (CI) backs this up: it builds this CLI and fails when a skill calls a `leap` command or flag that no longer exists, or reads a field of `leap run info -o json` or `leap version -o json` that this PR removed, renamed or retyped. It checks the skills PR linked on the `Skills:` line, or skills `main`. A new JSON output the skills read belongs in `jsonOutputs` in `scripts/skillscheck/main.go`. Text output (`leap auth whoami`, `leap server info`) is not checked.

## The skills this CLI ships with

`pkg/version/skills-ref` names the tensorleap/skills git tag this CLI ships with. `leap mcp config` installs it, `leap version -o json` reports it as `skillsRef`, and the skills update themselves to it, so users only get skills released with their CLI. To ship skill changes, tag tensorleap/skills (`skills-vN`) and set the file to that tag in a PR here. The release workflow fails if the tag doesn't exist, or if the skills at that tag call a `leap` command or flag this CLI doesn't have, or read a JSON output field removed since the previous release.
