package provider

type ErrorResponse struct {
	Errors []struct {
		Detail string `json:"detail"`
	} `json:"errors"`
}

type AtomicOperationResponse struct {
	AtomicResults []struct {
		Data struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"data"`
	} `json:"atomic:results"`
} 