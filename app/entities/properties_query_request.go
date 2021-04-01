package entities

type PropertiesQueryRequest struct {
	Data []PropertiesDataQueryRequest `json:"data"`
}

type PropertiesDataQueryRequest struct {
	Did        string   `json:"did"`
	Properties []string `json:"properties"`
}

type NoPropertiesDataRequest struct {
	Did string `json:"did"`
}

type NoPropertiesRequest struct {
	Data []NoPropertiesDataRequest `json:"data"`
}

type PropertiesDataRequest struct {
	Did string `json:"did"`
}

type PropertiesRequest struct {
	Data []PropertiesDataRequest `json:"data"`
}
