{{define "content"}}
Todo detail <br>
{{.Todo.Id}}<br>
{{.Todo.Todo}}<br>
<a href="{{urlfor "MainController.Get"}}"> Home </a> <br>
<a href="{{urlfor "TodoController.EditTodo" ":id" .Todo.Id}}"> Edit Todo </a> <br>
<a href="{{urlfor "TodoController.DeleteTodo" ":id" .Todo.Id}}"> Delete Todo </a> <br>
{{end}}
{{template "_base.tpl" .}}