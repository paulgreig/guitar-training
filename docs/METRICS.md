# Performance monitoring and metrics

The application exposes Prometheus metrics for performance monitoring, including menu selection latency, view counts, and Go runtime (CPU, RAM, GC). You can scrape these with Prometheus and visualize them in Grafana or Grafana Cloud.

## Enabling the metrics server

By default the metrics HTTP server listens on **port 9090**. Configure with:

- **`METRICS_PORT`** – Port for the `/metrics` endpoint (default: `9090`). Set to `0` to disable the metrics server.

```bash
# Default: metrics on :9090
go run cmd/server/main.go

# Custom port
METRICS_PORT=9091 go run cmd/server/main.go

# Disable metrics server (only file logging and in-memory metrics)
METRICS_PORT=0 go run cmd/server/main.go
```

While the TUI is running, open `http://localhost:9090/metrics` in a browser or with `curl` to see the Prometheus exposition format.

## Metrics exposed

### Application metrics (namespace `guitar_training_`)

| Metric | Type | Description |
|--------|------|-------------|
| `guitar_training_menu_selection_duration_seconds` | Histogram | Time taken for a menu selection (e.g. opening scales list or scale detail). Labels: none. |
| `guitar_training_views_total` | Counter | Total views by type. Label: `view` = `scales_list`, `scale_detail`, `lessons_list`, `lesson_detail`. |
| `guitar_training_key_presses_total` | Counter | Total key presses. |
| `guitar_training_data_load_total` | Counter | Data load attempts. Label: `result` = `success`, `error`. |

### Go runtime (from Prometheus Go collector)

- **GC**: e.g. `go_gc_duration_seconds`, `go_memstats_*`
- **Memory**: heap, stack, allocs, etc.
- **Goroutines**: `go_goroutines`

### Process (from Prometheus process collector)

- **CPU**: process CPU usage (where supported)
- **Memory**: process RSS, VMS (where supported)
- **Open FDs**: open file descriptors (where supported)

Metrics are prefixed with the `guitar_training` namespace where applicable (process collector uses the same namespace).

## Prometheus scrape config

Add a scrape job for the app. Because the TUI runs on the host (e.g. your laptop), you typically run Prometheus on the same machine or use an agent that can reach `localhost:9090`.

Example `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'guitar-training'
    static_configs:
      - targets: ['localhost:9090']
    scrape_interval: 15s
```

If Prometheus runs on another host, use a hostname/IP that can reach the machine running the app (e.g. `host.docker.internal:9090` or the machine’s LAN IP), and ensure port 9090 is not firewalled.

## Grafana Cloud

1. **Create a Grafana Cloud stack** (or use an existing one) and note the Prometheus endpoint (e.g. Prometheus “remote write” or “scrape” setup).

2. **Scraping from Grafana Cloud**  
   Grafana Cloud can scrape your app only if the app is reachable from the internet (e.g. via a tunnel or a VM with a public IP). For a local TUI, common options are:
   - Run a **Prometheus** (or Grafana Agent) instance on the same machine as the app, scrape `localhost:9090`, and **remote write** to Grafana Cloud Prometheus.
   - Use **Grafana Agent** in “host” mode to scrape `localhost:9090` and send metrics to Grafana Cloud.

3. **Grafana Cloud Prometheus remote write**  
   In Grafana Cloud:
   - Open **Connections** → **Add new connection** → **Prometheus** (or use the built-in Prometheus data source).
   - Get the **Remote Write** URL and auth (e.g. username + API token).
   - Configure your local Prometheus (or Grafana Agent) to remote write to that URL.

4. **Example: local Prometheus remote writing to Grafana Cloud**

   ```yaml
   # prometheus.yml
   global:
     external_labels:
       env: local

   scrape_configs:
     - job_name: 'guitar-training'
       static_configs:
         - targets: ['localhost:9090']
       scrape_interval: 15s

   remote_write:
     - url: 'https://prometheus-prod-XX-XX.grafana.net/api/prom/push'
       basic_auth:
         username: '<your-username>'
         password: '<your-grafana-cloud-api-token>'
   ```

   Replace the URL and credentials with the values from your Grafana Cloud stack.

5. **Dashboards**  
   In Grafana (or Grafana Cloud), create panels using the metrics above, for example:
   - **Menu selection latency**: `histogram_quantile(0.95, rate(guitar_training_menu_selection_duration_seconds_bucket[5m]))`
   - **Views by type**: `rate(guitar_training_views_total[5m])` by `view`
   - **Go memory**: `go_memstats_heap_inuse_bytes`, `go_memstats_alloc_bytes`
   - **GC**: `rate(go_gc_duration_seconds_sum[5m]) / rate(go_gc_duration_seconds_count[5m])`

## Summary

- Run the app with the metrics server enabled (default port 9090, or set `METRICS_PORT`).
- Point Prometheus (or Grafana Agent) at `http://localhost:9090/metrics`.
- Optionally remote write from that Prometheus to Grafana Cloud for visualization and alerting.
