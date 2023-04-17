package qiita

import "time"

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

type Tag struct {
	Name string `json:"name"`
}

type Tags []*Tag
