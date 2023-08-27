package types

type Quote struct {
	ID         int64
	Quote      string
	Author     string
	CreatedAt  string
	Owner_uuid string
}

type SaveQuote struct {
	Owner_uuid string
	Quote      string
	Author     string
}

type QuoteResponse struct {
	ID           string   `json:"_id,omitempty"`
	Content      string   `json:"content,omitempty"`
	Author       string   `json:"author,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	AuthorSlug   string   `json:"authorSlug,omitempty"`
	Length       int      `json:"length,omitempty"`
	DateAdded    string   `json:"dateAdded,omitempty"`
	DateModified string   `json:"dateModified,omitempty"`
}
