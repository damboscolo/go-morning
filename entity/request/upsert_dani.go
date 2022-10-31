package request

type UpsertDani struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
}
