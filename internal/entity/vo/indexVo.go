package vo

type IndexVo struct {
	Id    uint     `json:"id"`
	Title string   `json:"title"`
	Post  string   `json:"post"`
	Tags  []string `json:"tags"`
}
