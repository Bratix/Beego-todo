{{define "content"}}
<form action="" method="post">
    {{.Form | renderform}}
    <button type="submit">Submit</button>
</form> <br>

<a href="{{urlfor "RegisterController.Get" }}"> Register </a>



{{end}}
{{template "_base.tpl" .}}
