from collections.abc import Callable, Mapping

import numpy as np


Metric = Callable[[np.ndarray, np.ndarray], float]


def evaluate(
    y_true: np.ndarray,
    y_pred: np.ndarray,
    metrics: Mapping[str, Metric],
) -> dict[str, float]:
    """Score predictions; each metric owns its input requirements and errors."""
    return {
        name: float(metric(y_true, y_pred))
        for name, metric in metrics.items()
    }
