package obs

import (
	"sync/atomic"
	"time"
)

// Metrics holds simple in-memory counters for observability.
// This is intentionally lightweight and process-local.
type Metrics struct {
	AppStarts          uint64
	AppExits           uint64
	KeyPresses         uint64
	ScalesListViews    uint64
	LessonsListViews   uint64
	ScaleDetailViews   uint64
	LessonDetailViews  uint64
	DataLoadErrors     uint64
	DataLoadSuccesses  uint64
	LastAppStart       int64 // unix nano
	LastAppExit        int64 // unix nano
}

var globalMetrics Metrics

func incr(ptr *uint64) {
	atomic.AddUint64(ptr, 1)
}

// RecordAppStart increments the app start counter.
func RecordAppStart() {
	incr(&globalMetrics.AppStarts)
	atomic.StoreInt64(&globalMetrics.LastAppStart, time.Now().UnixNano())
}

// RecordAppExit increments the app exit counter.
func RecordAppExit() {
	incr(&globalMetrics.AppExits)
	atomic.StoreInt64(&globalMetrics.LastAppExit, time.Now().UnixNano())
}

// RecordKeyPress increments the keypress counter and Prometheus key_presses_total.
func RecordKeyPress() {
	incr(&globalMetrics.KeyPresses)
	IncKeyPressesTotal()
}

// RecordScalesListView increments the scales list view counter and Prometheus views_total.
func RecordScalesListView() {
	incr(&globalMetrics.ScalesListViews)
	IncViewsTotal("scales_list")
}

// RecordLessonsListView increments the lessons list view counter and Prometheus views_total.
func RecordLessonsListView() {
	incr(&globalMetrics.LessonsListViews)
	IncViewsTotal("lessons_list")
}

// RecordScaleDetailView increments the scale detail view counter and Prometheus views_total.
func RecordScaleDetailView() {
	incr(&globalMetrics.ScaleDetailViews)
	IncViewsTotal("scale_detail")
}

// RecordLessonDetailView increments the lesson detail view counter and Prometheus views_total.
func RecordLessonDetailView() {
	incr(&globalMetrics.LessonDetailViews)
	IncViewsTotal("lesson_detail")
}

// RecordDataLoadError increments the data load error counter and Prometheus data_load_total.
func RecordDataLoadError() {
	incr(&globalMetrics.DataLoadErrors)
	IncDataLoadTotal("error")
}

// RecordDataLoadSuccess increments the data load success counter and Prometheus data_load_total.
func RecordDataLoadSuccess() {
	incr(&globalMetrics.DataLoadSuccesses)
	IncDataLoadTotal("success")
}

// Snapshot returns a copy of the current metrics for logging/inspection.
func Snapshot() Metrics {
	return Metrics{
		AppStarts:          atomic.LoadUint64(&globalMetrics.AppStarts),
		AppExits:           atomic.LoadUint64(&globalMetrics.AppExits),
		KeyPresses:         atomic.LoadUint64(&globalMetrics.KeyPresses),
		ScalesListViews:    atomic.LoadUint64(&globalMetrics.ScalesListViews),
		LessonsListViews:   atomic.LoadUint64(&globalMetrics.LessonsListViews),
		ScaleDetailViews:   atomic.LoadUint64(&globalMetrics.ScaleDetailViews),
		LessonDetailViews:  atomic.LoadUint64(&globalMetrics.LessonDetailViews),
		DataLoadErrors:     atomic.LoadUint64(&globalMetrics.DataLoadErrors),
		DataLoadSuccesses:  atomic.LoadUint64(&globalMetrics.DataLoadSuccesses),
		LastAppStart:       atomic.LoadInt64(&globalMetrics.LastAppStart),
		LastAppExit:        atomic.LoadInt64(&globalMetrics.LastAppExit),
	}
}

