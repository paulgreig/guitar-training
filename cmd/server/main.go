package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/paulgreig/guitar-training/internal/obs"
	"github.com/paulgreig/guitar-training/internal/tui"
)

func main() {
	// Initialise logging and metrics.
	obs.InitLogger()
	obs.RecordAppStart()
	obs.Info("application starting")

	// Start Prometheus metrics server (CPU, RAM, GC + app metrics) for Grafana Cloud / Prometheus.
	obs.StartMetricsServer()
	defer obs.StopMetricsServer()

	start := time.Now()

	// Ensure we record app exit and log metrics summary.
	defer func() {
		obs.RecordAppExit()
		duration := time.Since(start)
		m := obs.Snapshot()
		obs.WithFields(
			obs.LevelInfo,
			"application exit",
			map[string]interface{}{
				"duration":            duration.String(),
				"app_starts":          m.AppStarts,
				"app_exits":           m.AppExits,
				"key_presses":         m.KeyPresses,
				"scales_list_views":   m.ScalesListViews,
				"lessons_list_views":  m.LessonsListViews,
				"scale_detail_views":  m.ScaleDetailViews,
				"lesson_detail_views": m.LessonDetailViews,
				"data_load_errors":    m.DataLoadErrors,
				"data_load_successes": m.DataLoadSuccesses,
			},
		)
	}()

	// Initialize the TUI application.
	p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		obs.Error("application error: %v", err)
		fmt.Printf("Error running application: %v\n", err)
		os.Exit(1)
	}

	obs.Info("application shutdown complete")
}
