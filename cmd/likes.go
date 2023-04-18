package cmd

import (
	"time"

	"github.com/koki-develop/qiita-ranking/internal/builder"
	"github.com/koki-develop/qiita-ranking/internal/config"
	"github.com/spf13/cobra"
)

var likesCmd = &cobra.Command{
	Use: "likes",
}

// デイリーいいね数ランキング
// https://qiita.com/koki_develop/items/fa223e1fa0ab057a54bc
var likesDailyCmd = &cobra.Command{
	Use: "daily",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		return update(&updateParameters{
			Config:   cfg,
			Title:    "Qiita デイリーいいね数ランキング【自動更新】",
			Template: builder.TemplateLikesDaily,
			Item:     cfg.Likes.Daily,
			Tags:     cfg.Likes.DailyByTag,
			From:     time.Now().AddDate(0, 0, -1),
			Stocks:   2,
		})
	},
}

// 週間いいね数ランキング
// https://qiita.com/koki_develop/items/b6cfc81906990b3a3e72
var likesWeeklyCmd = &cobra.Command{
	Use: "weekly",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		return update(&updateParameters{
			Config:   cfg,
			Title:    "Qiita 週間いいね数ランキング【自動更新】",
			Template: builder.TemplateLikesWeekly,
			Item:     cfg.Likes.Weekly,
			Tags:     cfg.Likes.WeeklyByTag,
			From:     time.Now().AddDate(0, 0, -7),
			Stocks:   10,
		})
	},
}
