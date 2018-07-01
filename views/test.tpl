<html>
<head></head>
<body>
    <p>{{.Members}}</p>
    <ul>
    {{ range $idx, $ele := .Members}}
        <li>{{ $ele }}</li>
    {{ end }}
    </ul>
    <p>{{.Numbers}}</p>
</body>
</html>
