"""Training orchestration stub.

A real experiment prepares a competition-appropriate validation split, then
runs fit -> predict -> evaluate. This standalone example uses toy arrays:

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

The example does not run on import. ``run_train`` still writes a placeholder
artifact; replace its body with your data preparation and training flow.
Splitting and OOF prediction assembly belong here, outside the evaluator.
"""

from pathlib import Path

from src.core.config import PROCESSED_DIR, ensure_data_dirs


def run_train(output: str | Path = PROCESSED_DIR / "model.txt") -> Path:
    """Run training and return the model artifact path."""
    ensure_data_dirs()

    output_path = Path(output)
    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text("placeholder model artifact\n")
    return output_path
