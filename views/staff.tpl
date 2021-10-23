{{define "content"}}

<h1>Hello {{ .Username }} </h1> 
<h2><a href="{{urlfor "LoginController.Logout"}}"> Logout </a> <br></h2>
    {{range $id,$todo := .Todos}}
    Todo id is {{$todo.Id}} <br>
    {{$todo.Todo}}<br>
    <a href="{{urlfor "TodoController.GetTodo" ":id" $todo.Id}}"> Details </a> <br>
    <a href="{{urlfor "TodoController.EditTodo" ":id" $todo.Id}}"> Edit Todo </a> <br>
    <a href="{{urlfor "TodoController.DeleteTodo" ":id" $todo.Id}}"> Delete Todo </a> <br>
    ____________________<br>



    {{end}}
    
{{end}}
{{template "_base.tpl" .}}