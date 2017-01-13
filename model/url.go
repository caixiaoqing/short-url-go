package model

type Url struct {
	Url	string	`json:"url"`
}

type Urls []Url

type ShortUrl struct {
	Short         string `json:"short"`
	RedirectShort string `json:"redirect"`
}

type OriginalUrl struct {
	Original string `json:"original"`
}
