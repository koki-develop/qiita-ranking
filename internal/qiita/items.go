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

func (cl *Client) ListItemsWithPagination(query string) (Items, error) {
	var rtn Items

	for i := 0; i < 100; i++ {
		items, err := cl.ListItems(&ListItemsParameters{Query: query, Page: i + 1, PerPage: 100})
		if err != nil {
			return nil, err
		}
		rtn = append(rtn, items...)

		if len(items) < 100 {
			break
		}
	}

	return rtn, nil
}
