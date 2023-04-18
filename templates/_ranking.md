{{ define "ranking" -}}
{{ if . -}}
{{ range $i, $item := . }}
## {{ inc $i }} 位: [{{ $item.Title }}]({{ $item.URL }})

{{ range $item.Tags }}[`{{ .Name }}`](https://qiita.com/tags/{{ .Name }}) {{ end }}
**{{ $item.LikesCount }}** いいね　**{{ $item.StocksCount }}** ストック
[@{{ $item.User.ID }}](https://qiita.com/{{ $item.User.ID }}) さん ( {{ $item.CreatedAt.Format "2006-01-02 15:04" }} に投稿 )
{{ end -}}
{{ else }}
ランキングに入る記事が見つかりませんでした。
{{ end -}}
{{- end }}
