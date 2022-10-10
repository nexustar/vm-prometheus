package selectapi

import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/querytracer"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/vmselectapi"
)

// promstorageAPI impelements github.com/VictoriaMetrics/VictoriaMetrics/lib/vmselectapi.API
type PromstorageAPI struct {
	s int
}

func (api *PromstorageAPI) InitSearch(qt *querytracer.Tracer, sq *storage.SearchQuery, deadline uint64) (vmselectapi.BlockIterator, error) {
	return nil, nil
}

func (api *PromstorageAPI) SearchMetricNames(qt *querytracer.Tracer, sq *storage.SearchQuery, deadline uint64) ([]string, error) {
	return nil, nil
}

func (api *PromstorageAPI) LabelValues(qt *querytracer.Tracer, sq *storage.SearchQuery, labelName string, maxLabelValues int, deadline uint64) ([]string, error) {
	return nil, nil
}

func (api *PromstorageAPI) TagValueSuffixes(qt *querytracer.Tracer, accountID, projectID uint32, tr storage.TimeRange, tagKey, tagValuePrefix string, delimiter byte, maxSuffixes int, deadline uint64) ([]string, error) {
	return nil, nil
}

func (api *PromstorageAPI) LabelNames(qt *querytracer.Tracer, sq *storage.SearchQuery, maxLableNames int, deadline uint64) ([]string, error) {
	return nil, nil
}

func (api *PromstorageAPI) SeriesCount(qt *querytracer.Tracer, accountID, projectID uint32, deadline uint64) (uint64, error) {
	return 0, nil
}

func (api *PromstorageAPI) TSDBStatus(qt *querytracer.Tracer, sq *storage.SearchQuery, focusLabel string, topN int, deadline uint64) (*storage.TSDBStatus, error) {
	return nil, nil
}

func (api *PromstorageAPI) DeleteSeries(qt *querytracer.Tracer, sq *storage.SearchQuery, deadline uint64) (int, error) {
	return 0, nil
}

func (api *PromstorageAPI) RegisterMetricNames(qt *querytracer.Tracer, mrs []storage.MetricRow, deadline uint64) error {
	return nil
}
