package types

type Log struct {
	Level      string     `json:"level"`
	Message    string     `json:"message"`
	ResourceId string     `json:"resourceId"`
	Timestamp  string     `json:"timestamp"`
	TraceId    string     `json:"traceid"`
	SpanId     string     `json:"spanId"`
	Commit     string     `json:"commi"`
	Metadata   []Metadata `json:"metadata"`
}

type Metadata struct {
	ParentResourceId string `json:"parentResourceId"`
}
