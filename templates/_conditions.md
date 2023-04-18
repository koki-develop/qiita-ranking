{{ define "conditions" -}}
{{ range . -}}
- {{ .Key }}: {{ .Value }}
{{ end }}
{{- end }}
