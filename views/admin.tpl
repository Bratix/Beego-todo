{{define "content"}}

<h2><a href="{{urlfor "LoginController.Logout"}}"> Logout </a> <br></h2>
    {{range $id,$user := .Users}}
    Todo id is {{$user.Id}} <br>
    {{$user.Username}}<br>
        {{if $user.IsAdmin }}
        <a href="{{urlfor "AdminController.RemoveAdmin" ":id" $user.Id}}"> Remove Admin </a> <br>
        
        {{else}}
        <a href="{{urlfor "AdminController.AddAdmin" ":id" $user.Id}}"> Add Admin </a> <br>
        
        {{end}}

        {{if $user.IsStaff }}
        <a href="{{urlfor "AdminController.RemoveStaff" ":id" $user.Id}}"> Remove Staff </a> <br>
        
        {{else}}
        <a href="{{urlfor "AdminController.AddStaff" ":id" $user.Id}}"> Add Staff </a> <br>
        
        {{end}}
        ____________________<br>

    {{end}}
{{end}}
{{template "_base.tpl" .}}