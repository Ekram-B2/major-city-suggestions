package results



// resultParser is applied to define variants of parse logic
type resultParser interface {
	// parseResult created structured results that can support the
	parseResult(map[string]interface{}) Results
}

func getParser(dataPoint string) resultParser {
	switch dataPoint {
	case "city":
		return cityParser{}
	default:
		return cityParser{}
	}
}
