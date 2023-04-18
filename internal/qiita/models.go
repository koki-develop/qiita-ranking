package qiita

import (
	"sort"
	"time"
)

type User struct {
	ID string `json:"id"`
}

type Item struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	LikesCount  int       `json:"likes_count"`
	StocksCount int       `json:"stocks_count"`
	URL         string    `json:"url"`
	User        User      `json:"user"`
	Tags        Tags      `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
}

type Items []*Item

func (items Items) Sort() {
	sort.Slice(items, func(i, j int) bool {
		if items[i].LikesCount != items[j].LikesCount {
			return items[i].LikesCount > items[j].LikesCount
		}
		if items[i].StocksCount != items[j].StocksCount {
			return items[i].StocksCount > items[j].StocksCount
		}
		return items[i].CreatedAt.After(items[j].CreatedAt)
	})
}

func (items Items) Filter() Items {
	var rtn Items

	for _, item := range items {
		if item.LikesCount > 0 {
			rtn = append(rtn, item)
		}
	}

	return rtn
}

type Tag struct {
	Name string `json:"name"`
}

type Tags []*Tag
