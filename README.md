# Kaggle Template

### Description

Develop machine learning models to predict. The goal is to improve understanding.

```bash
├── data           <---- Data directory
├── cli            <---- Command line tools
├── docs           <---- Documents, logs
│   ├── Log.md     <---- Day tracking work
│   ├── Paper.md   <---- Paper research
│   └── Scoring.md <---- Score tracking table
├── paper          <---- Papers to read, get inspired
├── nb             <---- Created on jupyter notebook
├── nb_download    <---- Public notebook from kaggle
├── README.md
├── requirements.txt
├── scripts        <----Utility scripts
├── src            <----Reusable code
│   ├── cli        <----Minimal logic Command-line entry point to call pipelines
│   ├── core       <----Shared config & utilities
│   ├── data       <----Data I/O + preprocessing helpers
│   ├── features   <----Feature engineering
│   ├── models     <----Model implementations
│   └── pipelines  <----Orchestration flows
```

### CLI Usage 

**cli/get_discussion**
This command line tool let you gather kaggle discussion and let you create local markdown files. This levarages you to summarize discussions with other tools

> [!NOTE]
> Initialize `.env` file from `.env.example` file to read competition

### Dataset

The dataset provided for this competition consists of.

| Name | Detail | Size | Link     |
| ---- | ------ | ---- | -------- |
| name |        |      | [Link]() |

### Evaluation

Use **fit -> predict -> evaluate** for validation. Models are responsible for
producing predictions. Evaluation is responsible only for scoring those predictions.
This runnable example uses the placeholder mean baseline and toy validation data:

```pycon
>>> import numpy as np
>>> from src.models.baseline import fit, predict
>>> from src.core.evaluation import evaluate
>>> from src.core.metrics import rmse
>>> X_train, y_train = np.array([[0.0], [1.0]]), np.array([1.0, 3.0])
>>> X_valid, y_valid = np.array([[2.0], [3.0]]), np.array([2.0, 4.0])
>>> model = fit(X_train, y_train)
>>> pred_valid = predict(model, X_valid)
>>> scores = evaluate(y_valid, pred_valid, metrics={"rmse": rmse})
>>> round(scores["rmse"], 4)
1.4142

```

`src/core/metrics.py` defines individual metrics; `src/core/evaluation.py` executes
any mapping of metric names to callables and returns `dict[str, float]`. Add more
metrics to the mapping or configure them with closures or `functools.partial`.
An empty mapping raises `ValueError`; metric exceptions propagate to the caller.
Each metric owns its input-shape requirements, including multiclass matrices.

Keep data splitting, training, inference, and OOF assembly in your pipeline.
The same evaluator accepts baseline, ensemble, post-processed, or OOF predictions.
See `src/pipelines/train.py` for the integration example; its training function
remains a placeholder.

Run the evaluation tests and executable examples with the existing NumPy dependency:

```bash
uv run --no-project --with 'numpy>=2.2.6' python -m unittest discover -s tests/core -v
uv run --no-project --with 'numpy>=2.2.6' python -m doctest README.md src/pipelines/train.py
```
