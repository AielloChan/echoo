# Echoo

一个用 Go 语言编写的 echo 服务器

## 四大模式

[echo 模式](#echo) [terminal 模式](#terminal) [file 模式](#file) [websocket 模式](#websocket)

### echo

此模式将会直接以 html 模式返回你发送的请求信息

### terminal

该模式会直接在命令行中输出访问 url 时的请求信息

### file

其实这个模式就是将上面 terminal 中输出的内容保存在文件里（以后会将此模式直接迁移为 log，
从而在所有模式中使用此功能）。

### websocket

通过 websocket 技术，你可以将指定链接被访问时使用的请求信息，直接显示在另一个页面中，且是实时更新的