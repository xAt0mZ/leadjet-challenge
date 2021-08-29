package query

type (
	result struct {
		Title       string
		URL         string
		Description string
	}

	queryPayload struct {
		QueryString string
		Results     []result
	}
)
