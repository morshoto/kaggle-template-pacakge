# Experiment Index

Use a stable ID for each question and reuse it in the notebook, config, finding, and generated output path. Keep this index concise and current.

| ID | Question | Status | Notebook | Config | Finding / result | Decision |
| --- | --- | --- | --- | --- | --- | --- |
| `exp-001` | Inspect the competition data and submission schema | Starter | [`notebooks/exploration/exp-001-data.ipynb`](../notebooks/exploration/exp-001-data.ipynb) | — | — | — |

## Suggested workflow

1. Write the question and hypothesis before changing the model.
2. Record the baseline and validation method; check for leakage.
3. Store reusable code in `src/`, notebook work under `notebooks/`, and versioned settings in `configs/<id>/`.
4. Keep large generated output in ignored `artifacts/<id>/`.
5. Add a finding under `docs/findings/` with the measured result, caveats, and decision; link it from this index.
6. Promote only comparable results to `docs/Score.md` and record submission provenance under `submissions/`.

Use the relevant template in [`docs/findings/`](../docs/findings/) for a result that should guide future work. Small exploratory runs that do not change a decision can be summarized in the index.
