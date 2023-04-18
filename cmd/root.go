package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/koki-develop/qiita-ranking/internal/builder"
	"github.com/koki-develop/qiita-ranking/internal/config"
	"github.com/koki-develop/qiita-ranking/internal/qiita"
	"github.com/spf13/cobra"
)

var (
	flagTag string
)

var rootCmd = &cobra.Command{
	Use: "cli",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

type updateParameters struct {
	Config   *config.Config
	Title    string
	Template builder.Template
	Item     *config.ConfigItem
	Tags     config.ConfigItems
	From     time.Time
	Stocks   int
	Tag      string
}

func update(params *updateParameters) error {
	cl := qiita.New(params.Config.QiitaAccessToken)
	query := []string{fmt.Sprintf("created:>=%s", params.From.Format(time.DateOnly))}
	if params.Tag != "" {
		query = append(query, fmt.Sprintf("tag:%s", params.Tag))
	}
	if params.Stocks > 0 {
		query = append(query, fmt.Sprintf("stocks:>=%d", params.Stocks))
	}
	items, err := cl.ListItemsWithPagination(strings.Join(query, " "))
	if err != nil {
		return err
	}

	conds := builder.Conditions{{Key: "集計期間", Value: fmt.Sprintf("%s ~ %s", params.From.Format(time.DateOnly), time.Now().Format(time.DateOnly))}}
	if params.Stocks > 0 {
		conds = append(conds, &builder.Condition{Key: "条件", Value: fmt.Sprintf("ストック数が **%d** 以上の記事", params.Stocks)})
	}

	b := builder.New()
	body, err := b.Build(&builder.BuildParameters{
		Template:   params.Template,
		Tags:       params.Tags,
		Conditions: conds,
		Items:      items,
	})
	if err != nil {
		return err
	}

	p := &qiita.UpdateItemParameters{
		Title: params.Title,
		Tags:  qiita.Tags{{Name: "Qiita"}, {Name: "いいね"}, {Name: "lgtm"}, {Name: "ランキング"}},
		Body:  string(body),
	}
	if params.Tag != "" {
		p.Tags = append(p.Tags, &qiita.Tag{Name: params.Tag})
	}
	if err := cl.UpdateItem(params.Item.ItemID, p); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(likesCmd)
	likesCmd.AddCommand(
		likesDailyCmd,
		likesWeeklyCmd,
	)

	for _, cmd := range []*cobra.Command{likesDailyCmd, likesWeeklyCmd} {
		cmd.Flags().StringVar(&flagTag, "tag", "", "")
	}
}
