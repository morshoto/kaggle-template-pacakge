# Finding: <short-title>

Use this template for a controlled experiment or a policy/heuristic change.
Use `kaggle_replay_forensic_template.md` for submission replay analysis and
`submission_audit_template.md` for package or evaluation audits.

**Date:** YYYY-MM-DD
**Status:** Draft / Promising / Adopted / Rejected / Superseded / Blocked
**Owner:** <name>
**Related notebook / script:** `nb/<notebook>.ipynb` / `src/<path>.py`
**Experiment / train id:** `<id>`
**Canonical source:** `<path>@<commit or SHA-256>`
**Baseline source:** `<path>@<commit or SHA-256>`
**Related submission:** <Kaggle submission id or N/A>
**Related score entry:** `docs/Score.md` → <model/name row>

---

## 1. Decision first

- **Decision:** Adopt / Keep as candidate / Retry / Do not adopt / Reference only
- **Evidence status:** Pass / Unavailable / Not run / Blocked
- **One-sentence finding:** <what changed and what the evidence shows>
- **Promotion boundary:** <what must pass before this can affect canonical policy or submission>

Do not claim causality, leaderboard quality, or hidden-state behavior unless the
validation below directly supports that claim.

---

## 2. Question and hypothesis

- **Question:** <clear research or engineering question>
- **Hypothesis:** <expected result before running the experiment>
- **Decision needed:** <future choice this finding should guide>

---

## 3. Scope and context

- Competition / simulator constraint:
- Previous finding or baseline:
- Relevant hidden-information / probability / strategy issue:
- In scope:
- Explicitly out of scope:
- Data leakage risk checked: Yes / No / N/A

---

## 4. Method

| Item | Value |
| ---- | ----- |
| Notebook / script | `nb/...` |
| Agent / policy | `<name and source identity>` |
| Baseline | `<source identity and configuration>` |
| Candidate | `<source identity and configuration>` |
| Dataset / decks | `<description>` |
| Opponents / profiles | `<exact list>` |
| Validation tier | Focused replay / paired screen / unpaired screen / leaderboard |
| Random seed(s) | `<seed or N/A>` |
| Seat handling | `<seat 0/1, balanced, or N/A>` |
| Number of games / folds | `<n>` |
| Engine seed applied | Yes / No / Unknown / N/A |
| Key parameters | `<important params>` |

### Procedure

1. <step>
2. <step>
3. <step>

---

## 5. Results

### Metrics

| Metric | Baseline | Candidate | Delta | Evidence / caveat |
| ------ | -------- | --------- | ----- | ----------------- |
| Win rate / score rate |  |  |  |  |
| Draw rate |  |  |  |  |
| Runtime errors |  |  |  |  |
| Search errors |  |  |  |  |
| Fallbacks |  |  |  |  |
| Runtime |  |  |  |  |

### Observed changes

- <important observation>
- <failure mode or guard fired>
- <what was not measured>

---

## 6. Validation and evidence limits

### Focused checks

| Check | Result | Artifact / test |
| ----- | ------ | --------------- |
| Decision replay | Pass / Fail / Not run | `src/eval/...` |
| Replay mismatches | `<count>` | `<log or JSON>` |
| Source/archive identity | Pass / Fail / N/A | `<SHA-256 or validator>` |
| Package validation | Pass / Fail / N/A | `<command or artifact>` |
| Runtime / timeout gate | Pass / Fail / N/A | `<summary>` |

### Interpretation boundary

- **Supported:** <claims directly supported by this experiment>
- **Not supported:** <claims this experiment cannot establish>
- **Confounders:** <variance, matchup bias, seat, simulator behavior, seed sensitivity>
- **Confidence:** Low / Medium / High

If a check was not run or unavailable, preserve that state instead of treating
it as a pass.

---

## 7. Decision and follow-ups

**Reason:** <why the decision follows from the evidence>

**Applies to:** EDA / training / canonical policy / submission / evaluation

**Supersedes:** `<docs/findings/old-finding.md>` or N/A

| Priority | Task | Owner | Link / Issue |
| -------- | ---- | ----- | ------------ |
| High |  |  |  |
| Medium |  |  |  |
| Low |  |  |  |

---

## 8. Artifacts and references

- Canonical source: `<path>`
- Candidate / archive: `<path>`
- Evaluation summary: `<path>`
- Decision logs / replay snapshots: `<path>`
- Competition notes: `docs/Competition.md`
- Score table: `docs/Score.md`
- Daily log: `docs/Log.md`
- Paper / external idea: `docs/Paper.md` or `<URL>`

### Reproducibility checklist

- [ ] Exact source identities are recorded for baseline and candidate.
- [ ] Exact parameters, profiles, opponents, seats, and seeds are recorded.
- [ ] Validation tier and engine-seed behavior are recorded.
- [ ] Runtime, search-error, and fallback status is recorded.
- [ ] Focused replay or package checks are linked, if applicable.
- [ ] Evidence limits and unrun/unavailable checks are explicit.
- [ ] Score was added to `docs/Score.md`, if applicable.
- [ ] Daily work was added to `docs/Log.md`, if applicable.
