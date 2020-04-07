package confluence

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Page struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Body   Body   `json:"body"`
}
type Storage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}
type Body struct {
	Storage Storage `json:"storage"`
}
type Confluence struct {
	user, pwd string
	client    *resty.Client
}

func NewConfluence(hosturl, user, pwd string) *Confluence {
	c := &Confluence{
		user:   user,
		pwd:    pwd,
		client: resty.New(),
	}
	c.client.SetBasicAuth(user, pwd)
	c.client.SetHostURL(hosturl)
	return c
}

func (c *Confluence) GetPageById(id string) *Page {
	var p Page
	url := fmt.Sprintf("/rest/api/content/%s?expand=body.storage", id)
	c.client.R().SetResult(&p).Get(url)
	return &p
}
