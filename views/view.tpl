<html>
<head></head>
<body>
    <table>
        <tr>
            <th>Title</th>
            <th>Description</th>
        </tr>
        <tr>
    {{ range $idx, $ele := .Audits}}
            <td>{{ $ele.Title }}</td>
            <td>{{ $ele.Summary }}</td>
        </tr>
    {{ end }}
    </table>
</body>
</html>

