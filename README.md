# Kaggle Competition Starter

A starter repository for Kaggle competitions. Replace this introduction with the competition name, objective, evaluation metric, and the current baseline once a project begins.

## Start here

- [Competition brief](docs/Competition.md): rules, metric, deadlines, and submission format.
- [Setup guide](docs/Setup.md): local environment and Kaggle credentials.
- [Experiment index](experiments/README.md): active questions, status, and links to evidence.
- [Repository structure](docs/Repository-Structure.md): where project files belong.
- [Score table](docs/Score.md): comparable validation and leaderboard results.

## Repository map

```text
configs/                 Versioned experiment and evaluation settings
data/                    Local raw, processed, and external data (ignored by Git)
notebooks/               Authored notebooks, grouped by purpose
  exploration/           EDA and hypothesis checks
  training/              Training and model development
  analysis/              Error analysis and evaluation
  submission/            Notebook based submission generation
references/              Curated external notebooks, papers, and discussions
experiments/              Experiment index and durable findings
artifacts/                Local generated runs, models, and reports (ignored)
src/                      Reusable implementation
cli/                      User-facing command line tools
scripts/                  Maintenance and automation scripts
submissions/              Small manifests for promoted submissions
docs/                      Competition brief, setup, findings, log, and score tracking
.github/workflows/         Lightweight repository checks
```

Optional folders such as `notebooks/training/`, `notebooks/analysis/`, `notebooks/submission/`, and `references/notebooks/` can be created when first needed. The template keeps only folders with starter content or an explicit README.

## Experiment flow

1. Add a row with a stable ID to [the experiment index](experiments/README.md).
2. Put the notebook under `notebooks/` and any reusable logic under `src/`.
3. Save the exact configuration under `configs/` and generated output under `artifacts/<experiment-id>/`.
4. Summarize the result, baseline, validation method, and limitations in `docs/findings/`.
5. Record promoted results in `docs/Score.md` and capture the submitted artifact manifest under `submissions/`.

## Useful commands

```bash
uv sync
uv run python scripts/download.py --competition <competition-slug> --dest data/raw
go run cli/validate.go notebooks/exploration/exp-001-data.ipynb
```

Kaggle credentials belong in a local `.env` file and must not be committed. See [the setup guide](docs/Setup.md) before downloading data or running notebooks.
