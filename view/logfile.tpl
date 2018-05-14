||||||||||||||||||||||||||||||||||| START ||||||||||||||||||||||||||||||||||||||
| Date {{.time}}
|
| {{.method}} {{.url}} {{.protocol}}
|
| Request ----------------------------------------------------------------------
| Host: {{.host}}
| Method: {{.method}}
| URL: {{.fullUrl}}
| Path: {{.url}}
| Cookie: {{.cookie}}
| ContentLength: {{.contentLength}}
| Client IP: {{.remoteAddr}}
| 
| Headers ----------------------------------------------------------------------
{{range $key, $value := .headers}}| {{$key}}: {{$value}}
{{end}}|
| URL parameters ---------------------------------------------------------------
{{range $key, $value := .urlParams}}| {{$key}}: {{$value}}
{{end}}|
| Body -------------------------------------------------------------------------
{{if .body}}| {{.body}}
{{else}}| (No request body)
{{end}}|
| Date {{.time}}
+----------------------------------- END ---------------------------------------
