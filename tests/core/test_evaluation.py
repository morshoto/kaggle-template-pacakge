from functools import partial
from types import MappingProxyType
import unittest

import numpy as np

from src.core.evaluation import evaluate
from src.core.metrics import rmse


class EvaluationTests(unittest.TestCase):
    def setUp(self):
        self.y_true = np.array([1.0, 2.0, 3.0])
        self.y_pred = np.array([1.0, 2.0, 4.0])

    def test_single_metric_scores_predictions(self):
        scores = evaluate(self.y_true, self.y_pred, {"rmse": rmse})
        self.assertEqual(set(scores), {"rmse"})
        self.assertAlmostEqual(scores["rmse"], np.sqrt(1.0 / 3.0))

    def test_multiple_metrics_preserve_names(self):
        scores = evaluate(
            self.y_true,
            self.y_pred,
            MappingProxyType({
                "custom RMSE": rmse,
                "absolute/error": lambda y, p: np.mean(np.abs(y - p)),
            }),
        )
        self.assertEqual(set(scores), {"custom RMSE", "absolute/error"})
        self.assertAlmostEqual(scores["custom RMSE"], np.sqrt(1.0 / 3.0))
        self.assertAlmostEqual(scores["absolute/error"], 1.0 / 3.0)

    def test_metric_outputs_are_python_floats(self):
        for value in (np.float32(0.5), np.float64(0.5), np.int64(2), 3):
            with self.subTest(value=value):
                scores = evaluate(self.y_true, self.y_pred, {"score": lambda y, p: value})
                self.assertIs(type(scores["score"]), float)
                self.assertEqual(scores["score"], float(value))

    def test_empty_metrics_raise_clear_error(self):
        with self.assertRaisesRegex(ValueError, "^At least one metric must be provided[.]$"):
            evaluate(self.y_true, self.y_pred, {})

    def test_metric_errors_propagate(self):
        error = RuntimeError("metric failed")

        def broken_metric(y, p):
            raise error

        with self.assertRaises(RuntimeError) as caught:
            evaluate(self.y_true, self.y_pred, {"broken": broken_metric})
        self.assertIs(caught.exception, error)

    def test_partial_receives_original_arrays(self):
        predictions = np.array([[0.8, 0.2], [0.1, 0.9], [0.4, 0.6]])
        calls = []

        def configured_metric(y, p, *, column):
            self.assertIs(y, self.y_true)
            self.assertIs(p, predictions)
            calls.append(column)
            return np.mean(p[:, column])

        scores = evaluate(
            self.y_true, predictions,
            {"class@1": partial(configured_metric, column=1)},
        )
        self.assertEqual(calls, [1])
        self.assertAlmostEqual(scores["class@1"], 1.7 / 3.0)


if __name__ == "__main__":
    unittest.main()
