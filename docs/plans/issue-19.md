## TDD Plan: Reusable model-agnostic evaluation (#19)

**Type**: Feature
**Issue**: https://github.com/morshoto/kaggle-template-pacakge/issues/19
**Complexity**: Low
**TDD Entry Point**: `test_single_metric_scores_predictions` in `tests/core/test_evaluation.py`.

### Issue Summary and Excerpt

Provide a reusable prediction-scoring boundary with named metrics and plain float results.

> Evaluation should depend on predictions, not on the model implementation.

### Scope

Add a small runner, unit tests, README usage, and an executable training-module
example. Keep existing metrics and training-stub behavior unchanged. No model
inspection, splitting, CV, tracking integration, registries, or shape validation.

### Behaviour Inventory / Acceptance Criteria as Tests

| ID | Observable behaviour | Unit test in tests/core/test_evaluation.py |
| --- | --- | --- |
| B1 | Existing RMSE scores supplied predictions | test_single_metric_scores_predictions |
| B2 | Multiple metrics preserve arbitrary names | test_multiple_metrics_preserve_names |
| B3 | NumPy scalar results become Python floats | test_metric_outputs_are_python_floats |
| B4 | Empty mappings raise a clear ValueError | test_empty_metrics_raise_clear_error |
| B5 | Metric exceptions propagate unchanged | test_metric_errors_propagate |
| B6 | Configured metrics receive original prediction matrices | test_partial_receives_original_arrays |

B1–B6 exercise the public function with predictions only; review verifies the
runner has no model, split, metric-specific, or framework dependencies.
README and training-module doctests demonstrate fit -> predict -> evaluate.

### Test-First Implementation Cycles

1. **Red**: Add B1–B3, B5–B6; run discovery and observe missing evaluation module.
   **Green**: Add `src/core/evaluation.py` with a Callable alias, Mapping input,
   and a dictionary comprehension converting each result to float.
   **Refactor**: Review names and duplication while focused tests stay green.
2. **Red**: Add B4; verify failure because an empty mapping returns `{}`.
   **Green**: Reject empty metrics with `ValueError` before executing metrics.
   **Refactor**: Keep the runner minimal; rerun focused tests.
3. Add executable documentation in README and `src/pipelines/train.py` using
   the existing baseline. No production logic change; validate with doctest.

Commit the plan, each red checkpoint, each green checkpoint, and documentation
with separate `update:` messages of at most eight words.

### Affected Files

- `tests/core/test_evaluation.py`: new unit coverage (red).
- `src/core/evaluation.py`: new scoring runner (green).
- `README.md`, `src/pipelines/train.py`: executable usage documentation.

### Test Commands

Use an isolated environment with only the existing NumPy dependency:

```bash
uv run --no-project --with 'numpy>=2.2.6' python -m unittest discover -s tests/core -v
uv run --no-project --with 'numpy>=2.2.6' python -m doctest README.md src/pipelines/train.py
```

Final verification runs both commands and `git diff --check`. There is no
configured Python lint or test suite in the current repository.

### Risks and Rollout

- Prediction shapes vary: B6 verifies matrices pass through untouched; metrics
  own shape rules.
- Metric failures must remain visible: B5 asserts the original exception.
- Training is a placeholder: documentation only preserves its artifact contract.
- No dependencies or migration required; callers opt into the new function.

### Definition of Done

All six behaviours pass; red failures recorded before corresponding logic;
refactoring only on green; examples execute; final checks pass; PR documents
validation and the unchanged training stub.
