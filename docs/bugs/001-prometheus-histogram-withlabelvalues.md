# Bug #001: menuSelectionDuration.WithLabelValues undefined

## Status
**Resolved**

## Summary
Build error after adding Prometheus performance monitoring: `menuSelectionDuration` was declared as a plain `prometheus.Histogram` but the code called `WithLabelValues(action)`, which exists only on `prometheus.HistogramVec`.

## Error message
```
internal/obs/prometheus.go:120:24: menuSelectionDuration.WithLabelValues undefined (type prometheus.Histogram has no field or method WithLabelValues)
```

## Root cause
- **`prometheus.Histogram`** – single metric, no labels; you observe with `.Observe(value)`.
- **`prometheus.HistogramVec`** – same metric with labels (e.g. `action`); you get a labeled observer with `.WithLabelValues(labelValues...)` and then call `.Observe(value)`.

We wanted per-action latency (e.g. `scales_list`, `scale_detail`), so the metric must be a **HistogramVec** with an `action` label. The type was mistakenly declared as `Histogram` and the code used `WithLabelValues(action)`.

## Fix
1. **Type**: Change `menuSelectionDuration` from `prometheus.Histogram` to `*prometheus.HistogramVec`.
2. **Construction**: Use `prometheus.NewHistogramVec(..., []string{"action"})` instead of `prometheus.NewHistogram(...)`.
3. **Usage**: Keep `menuSelectionDuration.WithLabelValues(action).Observe(duration.Seconds())` — this is valid for `HistogramVec`.

## Files changed
- `internal/obs/prometheus.go`: variable type and `NewHistogram` → `NewHistogramVec` with label `"action"`.

## Verification
After the fix, `go build ./cmd/server` completes successfully and `/metrics` exposes:
`guitar_training_menu_selection_duration_seconds_bucket{action="scales_list",...}`, etc.
