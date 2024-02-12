package v1

type Bar struct {
	Name        string `json:"name"`
	NumOfTables int    `json:"numOfTables,omitempty"`
}
