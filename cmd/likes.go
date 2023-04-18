package cmd

import (
	"fmt"
	"time"

	"github.com/koki-develop/qiita-ranking/internal/builder"
	"github.com/koki-develop/qiita-ranking/internal/config"
	"github.com/koki-develop/qiita-ranking/internal/qiita"
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

		cl := qiita.New(cfg.QiitaAccessToken)
		to := time.Now()
		from := to.AddDate(0, 0, -1)
		query := fmt.Sprintf("created:>=%s stocks:>=2", from.Format(time.DateOnly))
		items, err := cl.ListItemsWithPagination(query)
		if err != nil {
			return err
		}

		b := builder.New()
		body, err := b.Build(&builder.BuildParameters{
			Template: builder.TemplateLikesDaily,
			Tags:     cfg.Likes.DailyByTag,
			Conditions: map[string]string{
				"集計期間": fmt.Sprintf("%s ~ %s", from.Format(time.DateOnly), to.Format(time.DateOnly)),
				"条件":   "ストック数が **2** 以上の記事",
			},
			Items: items,
		})
		if err != nil {
			return err
		}

		params := &qiita.UpdateItemParameters{
			Title: "Qiita デイリーいいね数ランキング【自動更新】",
			Tags:  qiita.Tags{{Name: "Qiita"}, {Name: "いいね"}, {Name: "lgtm"}, {Name: "ランキング"}},
			Body:  string(body),
		}
		if err := cl.UpdateItem(cfg.Likes.Daily.ItemID, params); err != nil {
			return err
		}
		return nil
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

		cl := qiita.New(cfg.QiitaAccessToken)
		to := time.Now()
		from := to.AddDate(0, 0, -7)
		query := fmt.Sprintf("created:>=%s stocks:>=10", from.Format(time.DateOnly))
		items, err := cl.ListItemsWithPagination(query)
		if err != nil {
			return err
		}

		b := builder.New()
		body, err := b.Build(&builder.BuildParameters{
			Template: builder.TemplateLikesWeekly,
			Tags:     cfg.Likes.WeeklyByTag,
			Conditions: map[string]string{
				"集計期間": fmt.Sprintf("%s ~ %s", from.Format(time.DateOnly), to.Format(time.DateOnly)),
				"条件":   "ストック数が **10** 以上の記事",
			},
			Items: items,
		})
		if err != nil {
			return err
		}

		params := &qiita.UpdateItemParameters{
			Title: "Qiita 週間いいね数ランキング【自動更新】",
			Tags:  qiita.Tags{{Name: "Qiita"}, {Name: "いいね"}, {Name: "lgtm"}, {Name: "ランキング"}},
			Body:  string(body),
		}
		if err := cl.UpdateItem(cfg.Likes.Weekly.ItemID, params); err != nil {
			return err
		}
		return nil
	},
}
