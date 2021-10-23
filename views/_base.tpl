<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  
</head>

<body>
  
    {{block "content" .}} {{end}}
<a href="{{urlfor "AdminController.GetUsers"}}"> Admin </a> <br>
<a href="{{urlfor "StaffController.Get"}}"> Staff </a> <br>
<a href="{{urlfor "MainController.Get"}}"> Home </a> <br>
</body>
</html>