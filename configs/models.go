package configs

type ThreadHeader struct {
	Status int64             `json:"status"`
	Config map[string]string `json:"config"`
	Kw     []string          `json:"kw"`
}
