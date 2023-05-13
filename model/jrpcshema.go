package model

type JRPCSchema struct {
	Version string             `json:"version"`
	Info    JRPCSchemaInfo     `json:"info"`
	Methods []JRPCSchemaMethod `json:"methods"`
}

type JRPCSchemaInfo struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type JRPCSchemaMethod struct {
	Name         string                  `json:"name"`
	Description  string                  `json:"description"`
	Notification bool                    `json:"notification"`
	Params       []JRPCSchemaMethodParam `json:"params"`
	Result       JRPCSchemaMethodResult  `json:"result"`
}

type JRPCSchemaMethodParam struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Schema      JsonSchema `json:"schema"`
	Required    bool       `json:"required"`
}

type JRPCSchemaMethodResult struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Schema      JsonSchema `json:"schema"`
}
