# Kaggle replay / forensic finding: <short-title>

Use this template for replay-backed analysis of a Kaggle submission or public
episode corpus. This is a diagnostic report, not a policy-change report.

**Date:** YYYY-MM-DD
**Status:** Draft / Complete / Follow-up required / Superseded / Blocked
**Owner:** <name>
**Submission:** `<Kaggle submission id or name>`
**Target agent:** `<agent name and seat rule>`
**Simulator module:** `<version or Unknown>`
**Parser / analyzer:** `<script path and version/commit>`
**Related policy source:** `<path>@<commit or SHA-256>` or N/A

---

## 1. Executive summary

- **Public result:** `<wins/losses or N/A>`
- **Validation self-play result:** `<separate result or N/A>`
- **Evidence status:** Pass / Unavailable / Not run / Blocked
- **Finding:** <one or two sentences describing the strongest replay-backed signal>
- **Decision:** Diagnostic only / Follow-up investigation / Policy change proposed

Public episodes are a small diagnostic sample. Do not present them as a stable
rating, matchup estimate, or causal policy evaluation.

---

## 2. Corpus and acquisition

| Item | Value |
| ---- | ----- |
| Submission | `<id>` |
| Replay directory | `data/kaggle_episode_logs/<id>/` |
| Download index | `<path>` |
| Analysis index | `<path>` |
| Target agent | `<name>` |
| Public episodes | `<n>` |
| Validation episodes | `<n>` |
| Download failures | `<n>` |
| Parser failures | `<n>` |
| Action/observation mismatches | `<n>` |
| Terminal status | `<statuses>` |
| Hidden state policy | Visible evidence only / Other: `<explain>` |

### Acquisition and analysis commands

```text
<exact command used to download or locate replays>
<exact command used to parse and analyze them>
```

Keep validation self-play separate from public episodes in every aggregate.

---

## 3. Episode inventory

| Episode | Type | Seat | Opponent signature | Result | Steps | Target actions | Key resource | Final state |
| ------- | ---- | ---- | ------------------ | ------ | ----- | -------------- | ------------ | ----------- |
| `<id>` | public / validation | 0 / 1 | `<visible signature>` | win / loss / draw |  |  |  |  |

Opponent signatures are descriptive and based only on visible cards. Do not
reconstruct hidden decklists from absent cards.

---

## 4. Aggregate results

| Group | Games | Wins | Losses | Draws | Mean / median metric | Caveat |
| ----- | ----- | ---- | ------ | ----- | -------------------- | ------ |
| Public wins |  |  |  |  |  |  |
| Public losses |  |  |  |  |  |  |
| Validation self-play |  |  |  |  |  |  |

### Resource or route comparison

| Signal | Wins | Losses | Interpretation limit |
| ------ | ---- | ------ | -------------------- |
| `<visible signal>` |  |  |  |

These are descriptive comparisons unless the corpus and design support a
paired or controlled claim.

---

## 5. Replay-backed observations

### Observation 1: <short title>

- **Episodes:** `<episode ids>`
- **Raw steps:** `<step range or exact steps>`
- **Trace/report artifacts:** `<paths>`
- **Visible evidence:** <what the replay directly shows>
- **Interpretation:** <what it may indicate>
- **Not proven:** <causal or hidden-state claims that remain unsupported>

### Observation 2: <short title>

- **Episodes:** `<episode ids>`
- **Raw steps:** `<step range or exact steps>`
- **Trace/report artifacts:** `<paths>`
- **Visible evidence:** <what the replay directly shows>
- **Interpretation:** <what it may indicate>
- **Not proven:** <what remains unsupported>

Do not label an action strategically wrong solely because the replay exposes an
option index. Inspect legal options, public state, and the relevant policy
context before making that claim.

---

## 6. Hypotheses and confidence

| Hypothesis | Confidence | Supporting episodes / fields | What remains unproven |
| ---------- | ---------- | --------------------------- | --------------------- |
| `<hypothesis>` | Low / Medium / High | `<evidence>` | `<limit>` |

Keep separate categories separate, for example:

- failure to establish the main attacker;
- attacker established but resource conversion failed;
- opponent pressure or target selection;
- simulator termination or runtime behavior;
- draw variance or hidden information.

---

## 7. Decision and follow-up

- **Policy changed:** Yes / No
- **Canonical source changed:** Yes / No
- **Decision:** <diagnostic only / investigate / retry / reject hypothesis>
- **Reason:** <why the evidence supports this scope>
- **Promotion gate:** <controlled experiment, paired replay, or other required evidence>

| Priority | Follow-up | Owner | Artifact / issue |
| -------- | --------- | ----- | ---------------- |
| High |  |  |  |
| Medium |  |  |  |
| Low |  |  |  |

---

## 8. Limitations

- <sample size or repeated-matchup limitation>
- <visible versus hidden information limitation>
- <seat, opponent, or termination confounder>
- <parser or simulator limitation>
- <why this does not establish policy causality>

---

## 9. Artifacts and checklist

- Download index: `<path>`
- Analysis index: `<path>`
- Per-episode reports: `<path pattern>`
- Per-episode traces: `<path pattern>`
- Parser drafts: `<path pattern>`
- Findings output: `<path>`

- [ ] Public and validation episodes are separated.
- [ ] Download and parser failures are reported.
- [ ] Episode IDs and raw trace/report locations are recorded.
- [ ] Visible evidence is separated from hidden-state inference.
- [ ] Public sample is not presented as a rating estimate.
- [ ] Causal and policy claims are explicitly bounded.
- [ ] No policy or score-table change was implied without a separate gate.
