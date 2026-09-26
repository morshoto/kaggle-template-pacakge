# Submission / package audit: <short-title>

Use this template for submission-package, source/archive identity, evaluator,
runtime, or promotion-gate audits. This report verifies delivery state; it does
not by itself establish strategic quality or leaderboard performance.

**Date:** YYYY-MM-DD
**Status:** Draft / Ready / Not ready / Blocked / Superseded
**Owner:** <name>
**Submission:** `<Kaggle submission id or N/A>`
**Canonical source:** `<path>@<commit or SHA-256>`
**Candidate source:** `<path>@<commit or SHA-256>`
**Archive:** `<path>`

---

## 1. Audit conclusion

- **Delivery decision:** Ready / Not ready / Hold / Retry / Reference only
- **Evidence status:** Pass / Unavailable / Not run / Blocked
- **One-sentence conclusion:** <what was verified>
- **Unverified boundary:** <what this audit does not establish>
- **Required next action:** <smallest next action or N/A>

Do not promote `ready` package state into a claim that the policy is good,
causal, native, live, or leaderboard-competitive.

---

## 2. Artifact identity

| Artifact | Path | Commit / SHA-256 | Matches expected source? | Notes |
| -------- | ---- | ---------------- | ------------------------ | ----- |
| Canonical source | `<path>` | `<identity>` | Yes / No / N/A |  |
| Rendered candidate | `<path>` | `<identity>` | Yes / No / N/A |  |
| Embedded archive source | `<path>` | `<identity>` | Yes / No / N/A |  |
| Deck / config | `<path>` | `<identity>` | Yes / No / N/A |  |
| Submission artifact | `<path or id>` | `<identity>` | Yes / No / N/A |  |

### Identity checks

- Source/archive match: Pass / Fail / Not run / Unavailable
- Archive contains required entrypoints: Pass / Fail / Not run / Unavailable
- Deck/config matches intended candidate: Pass / Fail / Not run / Unavailable
- Working tree or commit identity recorded: Yes / No

---

## 3. Package and runtime validation

| Gate | Result | Command / artifact | Notes |
| ---- | ------ | ------------------ | ----- |
| Syntax / import check | Pass / Fail / Not run | `<command>` |  |
| Focused tests | Pass / Fail / Not run | `<command>` |  |
| Package validator | Pass / Fail / Not run | `<command>` |  |
| Archive validation | Pass / Fail / Not run | `<command>` |  |
| Smoke evaluation | Pass / Fail / Not run | `<command>` |  |
| Runtime errors | `<count>` | `<summary path>` |  |
| Search errors | `<count>` | `<summary path>` |  |
| Fallbacks | `<count>` | `<summary path>` |  |
| Timeout / incomplete games | `<count>` | `<summary path>` |  |

Preserve `unavailable`, `not run`, and `blocked` rather than converting them to
passes.

---

## 4. Evaluation context

| Item | Value |
| ---- | ----- |
| Evaluator / harness | `<path and version/commit>` |
| Evaluation profile | `<profile>` |
| Opponents / decks | `<exact list>` |
| Random seed(s) | `<seed or N/A>` |
| Seat handling | `<seat 0/1, balanced, or N/A>` |
| Games per arm | `<n>` |
| Engine seed applied | Yes / No / Unknown / N/A |
| Baseline identity | `<path>@<identity>` |
| Candidate identity | `<path>@<identity>` |
| Evaluation artifacts | `<path(s)>` |

### Result classification

Mark the strongest claim actually supported:

- [ ] Package/source identity only
- [ ] Focused decision replay
- [ ] Runtime or package smoke check
- [ ] Unpaired local screen
- [ ] Same-seed paired or counterfactual comparison
- [ ] Remote Kaggle submission result
- [ ] Stable leaderboard or rating claim

Do not call an unpaired screen a causal comparison. Record when the evaluator
reports `engine_seed_applied = false` or an equivalent limitation.

---

## 5. Evaluation results

| Arm | Opponent / profile | Games | Wins | Score rate | Runtime errors | Search errors | Fallbacks | Notes |
| --- | ----------------- | ----- | ---- | ---------- | -------------- | ------------- | --------- | ----- |
| Baseline |  |  |  |  |  |  |  |  |
| Candidate |  |  |  |  |  |  |  |  |

### Gate interpretation

- **Passed gates:** <list>
- **Failed gates:** <list>
- **Unrun or unavailable gates:** <list>
- **What the results support:** <claim>
- **What the results do not support:** <claim>

---

## 6. Promotion decision

- **Canonical policy changed:** Yes / No
- **Submission created or pushed:** Yes / No / N/A
- **Decision:** Promote / Do not promote / Hold / Retry
- **Reason:** <evidence-backed rationale>
- **Rollback or control:** `<source/archive identity>` or N/A
- **Next boundary:** <required validation before the next delivery step>

If a candidate is held or rejected, state which artifact remains canonical.

---

## 7. Artifacts and references

- Canonical source: `<path>`
- Candidate source: `<path>`
- Archive / package: `<path>`
- Validator output: `<path>`
- Evaluation summary: `<path>`
- Decision logs: `<path>`
- Related finding: `docs/findings/<file>.md`
- Score table: `docs/Score.md`
- Daily log: `docs/Log.md`

### Reproducibility checklist

- [ ] Every artifact has a path and identity.
- [ ] Source/archive equality was checked or marked unavailable.
- [ ] Required package and runtime gates have explicit states.
- [ ] Evaluator, profile, opponents, seats, seeds, and game counts are recorded.
- [ ] Engine-seed behavior is recorded.
- [ ] Errors, fallbacks, timeouts, and incomplete runs are reported.
- [ ] Evaluation evidence is classified by strength.
- [ ] Canonical promotion decision is explicit.
