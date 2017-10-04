<!DOCTYPE html>
<html>
    <head>
        <title>Support</title>
        {{ range $value := .Scripts }}
        <script href = "{{ $value }}" ></script>
        {{ end }}
        {{ range $value := .Css }}
        <link href="{{ $value }}" rel="stylesheet">
        {{ end }}
        {{ range $value := .Heads }}
        {{ $value | unescape }}
        {{ end }}
    </head>
    <body>
        <div class="Body">
            {{ yield }}
        </div>
    </body>
</html>