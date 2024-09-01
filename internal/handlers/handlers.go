package handlers

type Link struct {
	Href        string      `json:"href"`
	Rel         string      `json:"rel"`
	DescribedBy *Link       `json:"link"`
	Title       string      `json:"title"`
	HrefLang    string      `json:"hreflang"`
	Meta        interface{} `json:"meta"`
}

type SuccessResponse[T any] struct {
	Data     T               `json:"data"`
	Metadata interface{}     `json:"metadata,omitempty"`
	Links    map[string]Link `json:"links,omitempty"`
	Included []interface{}   `json:"includes,omitempty"`
}

type GenericError struct {
	Message string `json:"message"`
}
