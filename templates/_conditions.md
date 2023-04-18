{{ define "conditions" -}}
{{ range $key, $value := . -}}
- {{ $key }}: {{ $value }}
{{ end }}
{{- end }}
