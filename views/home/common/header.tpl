<!DOCTYPE html>
<html lang="cn">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{if ne "" (or .Title "")}} {{.Title}} - {{end}} {{GetSystemConf.Title}} - {{GetSystemConf.Tail}}</title>
    <meta name="description" content="{{if ne "" (or .Describe "")}}{{.Describe}}{{else}}{{GetSystemConf.Describe}}{{end}}">
    <meta name="keywords" content="{{if ne "" (or .Keyword "")}}{{.Keyword}}{{else}}{{GetSystemConf.Keyword}}{{end}}">
    <link rel="stylesheet" href="/static/home/lib/layui/css/layui.css">
    <link rel="stylesheet" href="/static/home/css/style.css">
    <link rel="stylesheet" href="/static/home/lib/iconfont/iconfont.css">
    <script src="/static/home/lib/layui/layui.js"></script>
    <script src="/static/home/lib/iconfont/iconfont.js"></script>
    <link rel=icon href=/static/favicon.ico>
</head>