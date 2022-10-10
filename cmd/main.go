package main

import (
	"flag"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/procutil"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/vmselectapi"
	"github.com/nexustar/vm-prometheus/pkg/selectapi"
)

var (
	vmselectAddr = flag.String("vmselectAddr", ":8401", "TCP address to accept connections from vmselect services")

	maxUniqueTimeseries          = flag.Int("search.maxUniqueTimeseries", 0, "The maximum number of unique time series, which can be scanned during every query. This allows protecting against heavy queries, which select unexpectedly high number of series. Zero means 'no limit'. See also -search.max* command-line flags at vmselect")
	maxTagKeys                   = flag.Int("search.maxTagKeys", 100e3, "The maximum number of tag keys returned per search")
	maxTagValues                 = flag.Int("search.maxTagValues", 100e3, "The maximum number of tag values returned per search")
	maxTagValueSuffixesPerSearch = flag.Int("search.maxTagValueSuffixesPerSearch", 100e3, "The maximum number of tag value suffixes returned from /metrics/find")

	disableRPCCompression = flag.Bool(`rpc.disableCompression`, false, "Whether to disable compression of the data sent from vmstorage to vmselect. "+
		"This reduces CPU usage at the cost of higher network bandwidth usage")
	denyQueriesOutsideRetention = flag.Bool("denyQueriesOutsideRetention", false, "Whether to deny queries outside of the configured -retentionPeriod. "+
		"When set, then /api/v1/query_range would return '503 Service Unavailable' error for queries with 'from' value outside -retentionPeriod. "+
		"This may be useful when multiple data sources with distinct retentions are hidden behind query-tee")
)

func main() {
	limits := vmselectapi.Limits{
		MaxLabelNames:       *maxTagKeys,
		MaxLabelValues:      *maxTagValues,
		MaxTagValueSuffixes: *maxTagValueSuffixesPerSearch,
	}
	svc, _ := vmselectapi.NewServer(*vmselectAddr, &selectapi.PromstorageAPI{}, limits, *disableRPCCompression)
	procutil.WaitForSigterm()
	svc.MustStop()

}
