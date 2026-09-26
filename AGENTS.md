# Contributor Guide

This repository is a reusable Kaggle competition starter. Keep changes small and make the project's competition assumptions explicit before building on them.

## Before changing the project

- Read `docs/Competition.md` before changing training, validation, or submission behavior.
- Read relevant code and nearby tests before editing implementation.
- Check `experiments/README.md` and related findings before changing an experiment decision.
- Avoid data leakage. Treat target-derived features, future information, and competition test data as boundaries to verify.
- Keep downloaded data, model checkpoints, and generated evaluation output out of Git unless a small fixture is needed for a reproducible check.

## Experiments

- Give each experiment a stable ID and add it to `experiments/README.md`.
- Keep authored notebooks under `notebooks/`, grouped by purpose. Put downloaded examples in `references/`.
- Put reusable implementation in `src/`, commands in `cli/`, and maintenance scripts in `scripts/`.
- Store versioned configs under `configs/<experiment-id>/` and local generated output under `artifacts/<experiment-id>/`.
- Record the question, baseline, validation method, data/version, seed, result, caveats, and decision in `docs/findings/`.
- Record promoted metrics in `docs/Score.md`; record a small manifest for each promoted submission in `submissions/`.

## Validation and submission

- Run the focused checks that cover the change. The pull request workflow runs the Go CLI tests and validates the starter notebook JSON.
- Execute notebooks locally before relying on their results; do not make Kaggle credentials a CI requirement.
- Before upload, confirm the exact notebook or package, competition policy, and output format. Never commit API tokens or `.env` files.

## Layout

See `docs/Repository-Structure.md` for directory responsibilities and optional-folder guidance. Update the README map when the starter layout changes.
