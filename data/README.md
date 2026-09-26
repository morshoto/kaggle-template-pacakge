# Data directory

This directory is ignored by Git except for this README and intentional `.gitkeep` files. Store competition inputs and derived datasets here; use `artifacts/` for generated evaluation results, models, and reports.

## Structure

- `data/raw/`: Raw files downloaded from Kaggle (train/test/sample_submission, etc.)
- `data/processed`/: Final features ready for training/inference
- `data/external`/: External datasets or resources

## Kaggle download

Configure Kaggle credentials as described in [`docs/Setup.md`](../docs/Setup.md), then download competition data into `data/raw/`.

Example:

```bash
pip install kaggle
export KAGGLE_USERNAME=your_name
export KAGGLE_KEY=your_key
export COMPETITION=your-competition
uv run python scripts/download.py --competition your-competition --dest data/raw
```

## Expected file names (example)

- data/raw/train.csv
- data/raw/test.csv
- data/raw/sample_submission.csv

Update these names as needed for each competition.
