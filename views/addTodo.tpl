{{define "content"}}
<form action="" method="post">
    {{.Form | renderform}}
    <button type="submit">Submit</button>
</form>
{{end}}
{{template "_base.tpl" .}}