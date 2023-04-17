package qiita

import (
	"net/http"
)

type ListItemsParameters struct {
	Query   string `url:"query"`
	Page    int    `url:"page"`
	PerPage int    `url:"per_page"`
}

func (cl *Client) ListItems(p *ListItemsParameters) (Items, error) {
	req, err := cl.newRequest(http.MethodGet, "items", p, nil)
	if err != nil {
		return nil, err
	}

	var items Items
	if err := cl.doRequest(req, &items); err != nil {
		return nil, err
	}

	return items, nil
}
