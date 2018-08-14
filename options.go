package g600

import "net/http"

// Options contains main options for a client creation. By default you should only provide
// ApiKey, other will be set without you.
type Options struct {
	APIKey     string
	AgentID    string
	URL        string
	HTTPClient *http.Client
}

func (o *Options) init() {
	if o.URL == "" {
		o.URL = DefaultURL
	}

	if o.HTTPClient == nil {
		o.HTTPClient = DefaultClient
	}

	if o.AgentID == "" {
		o.AgentID = DefaultAgent
	}
}
