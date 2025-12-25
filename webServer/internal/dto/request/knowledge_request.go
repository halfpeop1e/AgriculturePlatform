package request

type KnowledgeRequest struct {
	Title string `json:"title"`

	Excerpt string `json:"excerpt"`

	Content string `json:"content"`

	Tags []string `json:"tags"`

	Author string `json:"author"`
	Date   string `json:"date"`
}
