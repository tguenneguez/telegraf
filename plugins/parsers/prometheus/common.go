package prometheus

import (
	"github.com/influxdata/telegraf"
	dto "github.com/prometheus/client_model/go"
)

func mapValueType(mt dto.MetricType) telegraf.ValueType {
	switch mt {
	case dto.MetricType_COUNTER:
		return telegraf.Counter
	case dto.MetricType_GAUGE:
		return telegraf.Gauge
	case dto.MetricType_SUMMARY:
		return telegraf.Summary
	case dto.MetricType_HISTOGRAM:
		return telegraf.Histogram
	default:
		return telegraf.Untyped
	}
}

func getTagsFromLabels(m *dto.Metric, defaultTags map[string]string) map[string]string {
	result := make(map[string]string, len(defaultTags)+len(m.Label))
	for key, value := range defaultTags {
		result[key] = value
	}

	for _, label := range m.Label {
		if v := label.GetValue(); v != "" {
			result[label.GetName()] = v
		}
	}

	return result
}
