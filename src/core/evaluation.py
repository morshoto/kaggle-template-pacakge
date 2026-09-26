from collections.abc import Callable, Mapping
from typing import SupportsFloat

import numpy as np


Metric = Callable[[np.ndarray, np.ndarray], SupportsFloat]


def evaluate(
    y_true: np.ndarray,
    y_pred: np.ndarray,
    metrics: Mapping[str, Metric],
) -> dict[str, float]:
    """Score predictions; each metric owns its input requirements and errors."""
    if not metrics:
        raise ValueError("At least one metric must be provided.")

    return {
        name: float(metric(y_true, y_pred))
        for name, metric in metrics.items()
    }
