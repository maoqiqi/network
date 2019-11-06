# CURL


## 目录

* [简介](#简介)
* [CURL网站开发指南](#CURL网站开发指南)
  * [1.查看网页源码](#1.查看网页源码)
  * [2.自动跳转](#2.自动跳转)
  * [3.显示头信息](#3.显示头信息)
  * [4.显示通信过程](#4.显示通信过程)
  * [5.发送表单信息](#5.发送表单信息)
  * [6.HTTP动词](#6.HTTP动词)
  * [7.文件上传](#7.文件上传)
  * [8.Referer字段](#8.Referer字段)
  * [9.UserAgent字段](#9.UserAgent字段)
  * [10.cookie](#10.cookie)
  * [11.增加头信息](#11.增加头信息)
  * [12.HTTP认证](#12.HTTP认证)
* [常用参数](#常用参数)

## 简介

curl 是常用的命令行工具,用来请求 Web 服务器,然后得到和提取数据,显示在"标准输出"（stdout）上面。
它的名字就是客户端（client）的 URL 工具的意思。它支持多种协议,功能非常强大,命令行参数多达几十种。
如果熟练的话,完全可以取代 Postman 这一类的图形界面工具。


## CURL网站开发指南

我一向以为,curl只是一个编程用的函数库。

最近才发现,这个命令本身,就是一个无比有用的网站开发工具,下面举例讲解如何将它用于网站开发。

### 1.查看网页源码

直接在curl命令后加上网址,就可以看到网页源码。我们以网址www.sina.com为例（选择该网址,主要因为它的网页代码较短）：

```
curl www.sina.com
```

```
<html>
<head><title>301 Moved Permanently</title></head>
<body bgcolor="white">
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx</center>
</body>
</html>
```

如果要把这个网页保存下来,可以使用`-o`参数,这就相当于使用wget命令了。

```
curl -o [文件名] www.sina.com
```

例如：

```
curl -o sina.txt www.sina.com
```

### 2.自动跳转

有的网址是自动跳转的。使用`-L`参数,curl就会跳转到新的网址。

```
curl -L www.sina.com
```

键入上面的命令,结果就自动跳转为www.sina.com.cn。

### 3.显示头信息

`-i`参数可以显示http response的头信息,连同网页代码一起。

```
curl -i www.sina.com
```

```
HTTP/1.1 301 Moved Permanently
Server: nginx
Date: Wed, 30 Oct 2019 06:47:09 GMT
Content-Type: text/html
Content-Length: 178
Connection: keep-alive
Location: http://www.sina.com.cn/
Expires: Wed, 30 Oct 2019 06:48:20 GMT
Cache-Control: max-age=120
X-Via-SSL: ssl.22.sinag1.qxg.lb.sinanode.com
Age: 49
Via: https/1.1 ctc.guangzhou.ha2ts4.182 (ApacheTrafficServer/6.2.1 [cRs f ]), https/1.1 cnc.guangzhou.ha2ts4.41 (ApacheTrafficServer/6.2.1 [cRs f ])
X-Via-Edge: 15724180293643a06f63af0065a701af3a493
X-Cache: HIT.41
X-Via-CDN: f=edge,s=cnc.guangzhou.ha2ts4.51.nb.sinaedge.com,c=58.246.6.58;f=Edge,s=cnc.guangzhou.ha2ts4.41,c=112.90.6.51

<html>
<head><title>301 Moved Permanently</title></head>
<body bgcolor="white">
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx</center>
</body>
</html>
```

`-I`参数则是只显示http response的头信息。

### 4.显示通信过程

`-v`参数可以显示一次http通信的整个过程,包括端口连接和http request头信息。

```
* Rebuilt URL to: www.sina.com/
*   Trying 202.102.94.124...
* TCP_NODELAY set
* Connected to www.sina.com (202.102.94.124) port 80 (#0)
> GET / HTTP/1.1
> Host: www.sina.com
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
< Server: nginx
< Date: Wed, 30 Oct 2019 06:50:29 GMT
< Content-Type: text/html
< Content-Length: 178
< Connection: keep-alive
< Location: http://www.sina.com.cn/
< Expires: Wed, 30 Oct 2019 06:52:06 GMT
< Cache-Control: max-age=120
< X-Via-SSL: ssl.22.sinag1.qxg.lb.sinanode.com
< Age: 23
< Via: https/1.1 ctc.guangzhou.ha2ts4.182 (ApacheTrafficServer/6.2.1 [cRs f ]), https/1.1 ctc.nanjing.ha2ts4.32 (ApacheTrafficServer/6.2.1 [cRs f ])
< X-Via-Edge: 15724182290813a06f63a7c5e66ca54557499
< X-Cache: HIT.32
< X-Via-CDN: f=edge,s=ctc.nanjing.ha2ts4.75.nb.sinaedge.com,c=58.246.6.58;f=Edge,s=ctc.nanjing.ha2ts4.32,c=202.102.94.75
<
<html>
<head><title>301 Moved Permanently</title></head>
<body bgcolor="white">
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx</center>
</body>
</html>
* Connection #0 to host www.sina.com left intact
```

如果你觉得上面的信息还不够,那么下面的命令可以查看更详细的通信过程。

```
curl --trace output.txt www.sina.com
```

或者

```
curl --trace-ascii output.txt www.sina.com
```

运行后,请打开output.txt文件查看。

### 5.发送表单信息

发送表单信息有GET和POST两种方法。GET方法相对简单,只要把数据附在网址后面就行。

```
curl example.com/form.cgi?data=xxx
```

POST方法必须把数据和网址分开,curl就要用到--data参数。

```
curl -X POST --data "data=xxx" [URL]
```

如果你的数据没有经过表单编码,还可以让curl为你编码,参数是`--data-urlencode`。

```
curl -X POST --data-urlencode "date=xx xx" [URL]
```

### 6.HTTP动词

curl默认的HTTP动词是GET,使用`-X`参数可以支持其他动词。

```
curl -X POST www.example.com
```

```
curl -X DELETE www.example.com
```

### 7.文件上传

假定文件上传的表单是下面这样：

```
<form method="POST" enctype='multipart/form-data' action="upload.cgi">
　　　<input type=file name=upload>
　　　<input type=submit name=press value="OK">
</form>
```

你可以用curl这样上传文件：

```
curl --form upload=@localfilename --form press=OK [URL]
```

### 8.Referer字段

有时你需要在http request头信息中,提供一个referer字段,表示你是从哪里跳转过来的。

```
curl --referer [URL] [URL]
```

### 9.UserAgent字段

这个字段是用来表示客户端的设备信息。服务器有时会根据这个字段,针对不同设备,返回不同格式的网页,比如手机版和桌面版。

iPhone4的User Agent是:

```
Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7
```

curl可以这样模拟：

```
curl --user-agent "[User Agent]" [URL]
```

### 10.cookie

使用`--cookie`参数,可以让curl发送cookie。

```
curl --cookie "name=xxx" [URL]
```

至于具体的cookie的值,可以从http response头信息的`Set-Cookie`字段中得到。

`-c cookie-file`可以保存服务器返回的cookie到文件,`-b cookie-file`可以使用这个文件作为cookie信息,进行后续的请求。


```
curl -c cookies [URL]
curl -b cookies [URL]
```

### 11.增加头信息

有时需要在http request之中,自行增加一个头信息。`--header`参数就可以起到这个作用。

```
curl --header "Content-Type:application/json" [URL]
```

### 12.HTTP认证

有些网域需要HTTP认证,这时curl需要用到`--user`参数。

```
curl --user name:password [URL]
```


## 常用参数

### -A

`-A`参数指定客户端的用户代理标头,即`User-Agent`。curl 的默认用户代理字符串是`curl/[version]`。

```
curl -A 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36' https://google.com
```

上面命令将`User-Agent`改成 Chrome 浏览器。

```
curl -A '' https://google.com
```

上面命令会移除`User-Agent`标头。

也可以通过`-H`参数直接指定标头,更改`User-Agent`。

```
curl -H 'User-Agent: php/1.0' https://google.com
```

### -b

`-b`参数用来向服务器发送 Cookie。

```
curl -b 'foo=bar' https://google.com
```

上面命令会生成一个标头`Cookie: foo=bar`,向服务器发送一个名为`foo`、值为`bar`的 Cookie。

```
curl -b 'foo1=bar' -b 'foo2=baz' https://google.com
```

上面命令发送两个 Cookie。

```
curl -b cookies.txt https://www.google.com
```

上面命令读取本地文件`cookies.txt`,里面是服务器设置的 Cookie（参见`-c`参数）,将其发送到服务器。

### -c

`-c`参数将服务器设置的 Cookie 写入一个文件。

```
curl -c cookies.txt https://www.google.com
```

上面命令将服务器的 HTTP 回应所设置 Cookie 写入文本文件`cookies.txt``。

### -d

`-d`参数用于发送 POST 请求的数据体。

```
$ curl -d 'login=emma&password=123' -X POST https://google.com/login
# 或者
$ curl -d 'login=emma' -d 'password=123' -X POST  https://google.com/login
```

使用`-d`参数以后,HTTP 请求会自动加上标头`Content-Type : application/x-www-form-urlencoded`。并且会自动将请求转为 POST 方法,因此可以省略`-X POST`。

`-d`参数可以读取本地文本文件的数据,向服务器发送。

```
curl -d '@data.txt' https://google.com/login
```

上面命令读取data.txt文件的内容,作为数据体向服务器发送。

### --data-urlencode

`--data-urlencode`参数等同于-d,发送 POST 请求的数据体,区别在于会自动将发送的数据进行 URL 编码。

```
curl --data-urlencode 'comment=hello world' https://google.com/login
```

上面代码中,发送的数据`hello world`之间有一个空格,需要进行 URL 编码。

### -e

`-e`参数用来设置 HTTP 的标头Referer,表示请求的来源。

```
curl -e 'https://google.com?q=example' https://www.example.com
```

上面命令将`Referer`标头设为`https://google.com?q=example`。

`-H`参数可以通过直接添加标头`Referer`,达到同样效果。

```
curl -H 'Referer: https://google.com?q=example' https://www.example.com
```

### -F

`-F`参数用来向服务器上传二进制文件。

```
curl -F 'file=@photo.png' https://google.com/profile
```

上面命令会给 HTTP 请求加上标头`Content-Type: multipart/form-data`,然后将文件photo.png作为file字段上传。

`-F`参数可以指定 MIME 类型。

```
curl -F 'file=@photo.png;type=image/png' https://google.com/profile
```

上面命令指定 MIME 类型为`image/png`,否则 curl 会把 MIME 类型设为`application/octet-stream`。

`-F`参数也可以指定文件名。

```
curl -F 'file=@photo.png;filename=me.png' https://google.com/profile
```

上面命令中,原始文件名为`photo.png`,但是服务器接收到的文件名为`me.png`。

### -G

`-G`参数用来构造 URL 的查询字符串。

```
curl -G -d 'q=kitties' -d 'count=20' https://google.com/search
```

上面命令会发出一个 GET 请求,实际请求的 URL 为`https://google.com/search?q=kitties&count=20`。如果省略`-G`,会发出一个 POST 请求。

如果数据需要 URL 编码,可以结合`--data--urlencode`参数。

```
curl -G --data-urlencode 'comment=hello world' https://www.example.com
```

### -H

`-H`参数添加 HTTP 请求的标头。

```
curl -H 'Accept-Language: en-US' https://google.com
```

上面命令添加 HTTP 标头`Accept-Language: en-US`。

```
curl -H 'Accept-Language: en-US' -H 'Secret-Message: xyzzy' https://google.com
```

上面命令添加两个 HTTP 标头。

```
curl -d '{"login": "emma", "pass": "123"}' -H 'Content-Type: application/json' https://google.com/login
```

上面命令添加 HTTP 请求的标头是`Content-Type: application/json`,然后用`-d`参数发送 JSON 数据。

### -i

`-i`参数打印出服务器回应的 HTTP 标头。

```
curl -i https://www.example.com
```

上面命令收到服务器回应后,先输出服务器回应的标头,然后空一行,再输出网页的源码。

### -I

`-I`参数向服务器发出 HEAD 请求,然会将服务器返回的 HTTP 标头打印出来。

```
curl -I https://www.example.com
```

上面命令输出服务器对 HEAD 请求的回应。

`--head`参数等同于`-I`。

```
curl --head https://www.example.com
```

### -k

`-k`参数指定跳过 SSL 检测。

```
curl -k https://www.example.com
```

上面命令不会检查服务器的 SSL 证书是否正确。

### -L

`-L`参数会让 HTTP 请求跟随服务器的重定向。curl 默认不跟随重定向。

```
curl -L -d 'tweet=hi' https://api.twitter.com/tweet
```

### --limit-rate

`--limit-rate`用来限制 HTTP 请求和回应的带宽,模拟慢网速的环境。

```
curl --limit-rate 200k https://google.com
```

上面命令将带宽限制在每秒 200K 字节。

### -o

`-o`参数将服务器的回应保存成文件,等同于`wget`命令。

```
curl -o example.html https://www.example.com
```

上面命令将`www.example.com`保存成`example.html`。

### -O

`-O`参数将服务器回应保存成文件,并将 URL 的最后部分当作文件名。

```
curl -O https://www.example.com/foo/bar.html
```

上面命令将服务器回应保存成文件,文件名为`bar.html`。

### -s

`-s`参数将不输出错误和进度信息。

```
curl -s https://www.example.com
```

上面命令一旦发生错误,不会显示错误信息。不发生错误的话,会正常显示运行结果。

如果想让 curl 不产生任何输出,可以使用下面的命令。

```
curl -s -o /dev/null https://google.com
```

### -S

`-S`参数指定只输出错误信息,通常与`-s`一起使用。

```
curl -s -o /dev/null https://google.com
```

上面命令没有任何输出,除非发生错误。

### -u

`-u`参数用来设置服务器认证的用户名和密码。

```
curl -u 'bob:12345' https://google.com/login
```

上面命令设置用户名为`bob`,密码为`12345`,然后将其转为 HTTP 标头`Authorization: Basic Ym9iOjEyMzQ1`。

curl 能够识别 URL 里面的用户名和密码。

```
curl https://bob:12345@google.com/login
```

上面命令能够识别 URL 里面的用户名和密码,将其转为上个例子里面的 HTTP 标头。

```
curl -u 'bob' https://google.com/login
```

上面命令只设置了用户名,执行后,curl 会提示用户输入密码。

### -v

`-v`参数输出通信的整个过程,用于调试。

```
curl -v https://www.example.com
```

`--trace`参数也可以用于调试,还会输出原始的二进制数据。

```
curl --trace - https://www.example.com
```

### -x

`-x`参数指定 HTTP 请求的代理。

```
curl -x socks5://james:cats@myproxy.com:8080 https://www.example.com
```

上面命令指定 HTTP 请求通过`myproxy.com:8080`的 socks5 代理发出。

如果没有指定代理协议,默认为 HTTP。

```
curl -x james:cats@myproxy.com:8080 https://www.example.com
```

上面命令中,请求的代理使用 HTTP 协议。

### -X

`-X`参数指定 HTTP 请求的方法。

```
curl -X POST https://www.example.com
```

上面命令对https://www.example.com发出 POST 请求。




|参数|描述|
|:-----|:-----|
|-A|指定客户端的用户代理标头,即User-Agent。默认为`curl/[version]`。|
| |`curl -v -A '' http://127.0.0.1:8080/test` 移除User-Agent|
| |`curl -v -H 'User-Agent: golang/1.0' http://127.0.0.1:8080/test` 通过-H更改User-Agent|
|-b/--cookie|向服务器发送Cookie|
| |`curl -v -b user_name=admin http://127.0.0.1:8080/test` 生成标头`Cookie: user_name=admin`|
| |`curl -v --cookie user_name=admin http://127.0.0.1:8080/test` 效果同上|
|-c|将服务器设置的Cookie写入一个文件|
| |`curl -v -c cookies.txt http://127.0.0.1:8080/test` 将Cookie写入[cookies.txt](#cookies文本内容)|
|-d/--data|发送POST请求的数据体|
| |`curl -v -d user_name=admin -d password=123456 http://127.0.0.1:8080/test`|
| |`curl -v -d 'user_name=admin&password=123456' http://127.0.0.1:8080/test`|
| |两种方式一样。使用`-d`参数以后,会自动将请求转为POST,并且会自动加上标头`Content-Type: application/x-www-form-urlencoded`。|
| |`curl -v -d '@data.txt' http://127.0.0.1:8080/test` 读取本地[data.txt](#data文本内容)文件的数据,向服务器发送。|
|--data-urlencode|等同于`-d`,发送POST请求的数据体,区别在于会自动将发送的数据进行URL编码。|
| |`curl -v --data-urlencode 'user_name=ad min&password=123456' http://127.0.0.1:8080/test`|
| |上面代码中,发送的数据`user_name=ad min&password=123456`之间有一个空格,需要进行URL编码。|
|-e|用来设置标头`Referer`,表示请求的来源|
| |`curl -v -e 'http://127.0.0.1:8080/test' http://127.0.0.1:8080/test`|
| |`curl -v -H 'Referer: http://127.0.0.1:8080/test'  http://127.0.0.1:8080/test` 效果同上|
|-F/--file|向服务器上传二进制文件|
| |`curl -v -F 'file=@logo.png'  http://127.0.0.1:8080/test`|
| |使用-F参数以后,请求会自动加上标头`Content-Type: multipart/form-data`。|
| |`curl -v -F 'file=@logo.png;type=image/png'  http://127.0.0.1:8080/test` 指定MIME类型|
| |上面命令指定MIME类型为`image/png`,否则会把MIME类型设为`application/octet-stream`。|
| |`curl -v -F 'file=@logo.png;filename=test.png;type=image/png'  http://127.0.0.1:8080/test 指定文件名`|
|-G|构造URL的查询字符串|
| |`curl -v -G -d user_name=admin -d password=123456 http://127.0.0.1:8080/test`|
| |`curl -v -G -d 'user_name=admin&password=123456' http://127.0.0.1:8080/test`|
| |以上两种方式一样。使用`-G`参数后会发出一个GET请求|
| |`curl -v -G --data-urlencode 'user_name=ad min&password=123456' http://127.0.0.1:8080/test` 如果数据需要URL编码,可以结合`--data-urlencode`参数。|
|-H|添加请求的标头|
| |`curl -v -H 'Accept-Language: en-US' -H 'Secret-Message: xyzzy' http://127.0.0.1:8080/test`|
| |`curl -v -d '{"user_name":"admin","password":"123456"}' -H 'content-type: application/json' http://127.0.0.1:8080/test`|
| |上面命令添加HTTP请求的标头是`Content-Type: application/json`,然后用`-d`参数发送JSON数据。|
|-i|打印出服务器回应的HTTP标头|
| |`curl -i http://127.0.0.1:8080/test` 上面命令收到服务器响应后,先输出服务器响应的标头,然后空一行,再输出网页的源码。|
|-I/--head|向服务器发出HEAD请求,然会将服务器响应的HTTP标头打印出来。|
| |`curl -v -I http://127.0.0.1:8080/test`|
| |`curl -v --head http://127.0.0.1:8080/test`|
|-k|指定跳过SSL检测|
| |`curl -v -k http://127.0.0.1:8080/test`|
|-L|让HTTP请求跟随服务器重定向,默认不跟随重定向|
| |`curl -v -L http://127.0.0.1:8080/test`|
|--limit-rate|限制HTTP请求和回应的带宽,模拟慢网速的环境|
| |`curl -v --limit-rate 200k http://127.0.0.1:8080/test` 将带宽限制在每秒200K字节|
|-o|将服务器的回应保存成文件,等同于`wget`命令|
| |`curl -v -o test.html http://127.0.0.1:8080/test`|
|-O|将服务器回应保存成文件,并将URL的最后部分当作文件名|
| |`curl -v -O http://127.0.0.1:8080/test.html`|
|-s|不输出错误和进度信息|
| |`curl -s http://127.0.0.1:8080/test` 不会显示错误信息,不发生错误的话,会正常显示运行结果|
|-S|只输出错误信息|
| |`curl -S http://127.0.0.1:8080/test`|
|-u|设置服务器认证的用户名和密码|
| |`curl -v -u 'admin"123456' http://127.0.0.1:8080/test`|
| |设置用户名为`admin`,密码为`123456`,然后将其转为HTTP标头`Authorization: Basic YWRtaW4iMTIzNDU2OjEyMzQ1Ng==`。|
| |`curl -v  http://admin:123456@127.0.0.1:8080/test` curl能够识别URL里面的用户名和密码|
| |`curl -v -u admin http://127.0.0.1:8080/test` 只设置用户名,执行后,curl会提示用户输入密码|
|-v/--trace|输出通信的整个过程,用于调试|
| |`curl -v http://127.0.0.1:8080/test`|
| |`curl --trace -  http://127.0.0.1:8080/test` `--trace`也可以用于调试,还会输出原始的二进制数据|
|-x|指定HTTP请求的代理|
| |`curl -v -x http://10.64.144.3:8123 https://www.google.com`|
|-X|指定HTTP请求的方法|
| |`curl -v -X POST http://127.0.0.1:8080/test` 发出POST请求|
|||
|||
|||
|||
|||
|||
|||

### cookies文本内容

```
# Netscape HTTP Cookie File
# https://curl.haxx.se/docs/http-cookies.html
# This file was generated by libcurl! Edit at your own risk.

127.0.0.1	FALSE	/	FALSE	61572939117	user_id	10001
```

#### data文本内容

```
user_name=admin&password=123456
```