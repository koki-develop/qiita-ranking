{{ define "tags" -}}
{{ range . }}[`{{ .Tag }}`](https://qiita.com/items/{{ .ItemID }}) {{ end }}
{{- end }}
