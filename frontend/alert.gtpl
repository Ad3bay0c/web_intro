{{define "Alert"}}
    {{if .Message}}
        <div class="alert alert-{{.Color}}">
            {{ .Message}}
        </div>
    {{end}}

{{end}}