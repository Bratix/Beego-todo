{{define "content"}}

<h1>Hello {{ .Username }} </h1>
{{range $id,$todo := .Todos}}
Todo id is {{$todo.Id}} <br>
{{$todo.Todo}}<br>
<a href="{{urlfor "TodoController.Get" ":id" $todo.Id}}"> Details </a> <br>
<a href="{{urlfor "TodoController.EditTodo" ":id" $todo.Id}}"> Edit Todo </a> <br>
<a href="{{urlfor "TodoController.DeleteTodo" ":id" $todo.Id}}"> Delete Todo </a> <br>
____________________<br>



{{end}}
<a href="{{urlfor "TodoController.AddTodo"}}"> Add Todo </a> <br>
{{end}}
{{template "_base.tpl" .}}