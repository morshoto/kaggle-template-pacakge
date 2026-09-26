# Repository Structure

This layout keeps authored notebooks together while giving experiments, references, and generated output separate homes. Create optional folders when they gain content; avoid placeholder directories that do not explain a real convention.

## Directory responsibilities

| Path | Use |
| --- | --- |
| `notebooks/` | Authored notebooks, grouped as `exploration/`, `training/`, `analysis/`, or `submission/` when those categories are used. |
| `references/` | Curated external material such as downloaded notebooks, papers, and discussion threads. Add subfolders when they contain material. |
| `experiments/` | Navigation index connecting experiment IDs to notebooks, configs, findings, and results. |
| `configs/` | Versioned run settings, grouped by experiment ID. |
| `data/` | Local raw, processed, and external data. Large/private data is ignored by Git. |
| `artifacts/` | Local evaluation output, reports, checkpoints, and models. Ignored by Git. |
| `submissions/` | Small manifests describing promoted submission source, config, and validation. |
| `src/`, `cli/`, `scripts/` | Reusable implementation, user-facing commands, and maintenance automation. |
| `docs/` | Competition brief, setup instructions, findings, compact project log, and score table. |
| `.github/workflows/` | Small checks that can run without Kaggle credentials or private data. |

## Naming and traceability

Use stable IDs such as `exp-001` or `train-004` across notebook names, config directories, findings, and artifact paths. Prefer descriptive names over `copy`, `backup`, or `executed` suffixes. Keep one row per experiment in `experiments/README.md` and link to the relevant notebook, config, finding, and result.

Keep detailed findings as individual Markdown files under `docs/findings/`, using the templates there when relevant. The experiment index links to each finding. A finding should state the question, baseline, validation method, data/version, seed, result, limitations, and decision. Put raw logs under `artifacts/<experiment-id>/` and summarize patterns instead of creating one document per generated record.

## Data and outputs

`data/README.md` describes provenance and which files can be versioned. Commit only small, redistributable fixtures needed for reproducible checks. Store local run output under `artifacts/<experiment-id>/`; include the command and config needed to interpret it. Keep `.env`, credentials, large datasets, and generated model files out of Git.

## Optional directories

Create `notebooks/training/`, `notebooks/analysis/`, `notebooks/submission/`, and `references/notebooks/` only when needed. This keeps a fresh project navigable without presenting empty folders as active parts of the workflow.
