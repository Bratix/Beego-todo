{{define "content"}}
 <h1>Todo detail</h1> <br>
<h2>{{.Todo.User.Username}} is the owner</h2><br>
{{.Todo.Id}}<br>
{{.Todo.Todo}}<br>
<a href="{{urlfor "MainController.Get"}}"> Home </a> <br>
<a href="{{urlfor "TodoController.EditTodo" ":id" .Todo.Id}}"> Edit Todo </a> <br>
<a href="{{urlfor "TodoController.DeleteTodo" ":id" .Todo.Id}}"> Delete Todo </a> <br>
{{end}}
{{template "_base.tpl" .}}