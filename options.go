package g600

import "net/http"

// Options contains main options for a client creation. By default you should only provide
// ApiKey, other will be set without you.
type Options struct {
	ApiKey     string
	AgentID    string
	URL        string
	HttpClient *http.Client
}

func (o *Options) init() {
	if o.URL == "" {
		o.URL = DefaultURL
	}

	if o.HttpClient == nil {
		o.HttpClient = DefaultClient
	}

	if o.AgentID == "" {
		o.AgentID = DefaultAgent
	}
}
