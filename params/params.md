# GO获取请求参数


## 目录

* [Request请求结构](#Request请求结构)
* [模拟请求查看参数](#模拟请求查看参数)
  * [GET请求](#GET请求)
  * [GET请求带参数](#GET请求带参数) 
  * [POST请求带参数](#POST请求带参数)
  * [POST请求带参数并且地址也带参数](#POST请求带参数并且地址也带参数)
  * [POST请求带参数并上传文件](#POST请求带参数并上传文件)
* [About](#About)
* [License](#License)

## Request请求结构

* `Request.URL.RawQuery`：地址问号后的参数
* `Request.Form`：存储了get,post,put参数,在使用之前需要调用ParseForm方法
* `Request.PostForm`：存储了post,put参数,在使用之前需要调用ParseForm方法
* `Request.MultipartForm`：存储了包含了文件上传的表单的post参数,在使用之前需要调用ParseMultipartForm方法

## GO代码示例

```
_ = c.Request.ParseForm()
_ = c.Request.ParseMultipartForm(32 << 20)

// 地址问号后的参数
queryForm, _ := url.ParseQuery(c.Request.URL.RawQuery)
// 存储了get,post,put参数,在使用之前需要调用ParseForm方法
form := c.Request.Form
// 存储了post,put参数,在使用之前需要调用ParseForm方法
postForm := c.Request.PostForm
// 存储了包含了文件上传的表单的post参数,在使用之前需要调用ParseMultipartForm方法
multipartForm := c.Request.MultipartForm
```

> **注意:** 其中Request.Form和Request.PostForm必须在调用ParseForm之后,才会有数据,否则则是空数组。
而Request.FormValue和Request.PostFormValue()无需调用ParseForm就能读取数据。

## 模拟请求查看参数

### GET请求 

`curl http://127.0.0.1:8080/test`

```
{
    "query_form":{

    },
    "form":{

    },
    "post_form":{

    },
    "multipart_form":null,
    "body":""
}
```

### GET请求带参数 

`curl http://127.0.0.1:8080/test?user_name=admin&password=123456`

```
{
    "query_form":{

    },
    "form":{

    },
    "post_form":{

    },
    "multipart_form":null,
    "body":""
}
```

### POST请求带参数

`curl -X POST --data "user_name=admin&password=123456" http://127.0.0.1:8080/test`

普通的post表单请求:`Content-Type: application/x-www-form-urlencoded`

```
{
    "query_form":{

    },
    "form":{
        "password":[
            "123456"
        ],
        "user_name":[
            "admin"
        ]
    },
    "post_form":{
        "password":[
            "123456"
        ],
        "user_name":[
            "admin"
        ]
    },
    "multipart_form":null,
    "body":""
}
```

### POST请求带参数并且地址也带参数

`curl -X POST --data "user_name=admin&password=123456" http://127.0.0.1:8080/test?user_name=zhangsan`

普通的post表单请求:`Content-Type: application/x-www-form-urlencoded`

```
{
    "query_form":{
        "user_name":[
            "zhangsan"
        ]
    },
    "form":{
        "password":[
            "123456"
        ],
        "user_name":[
            "admin",
            "zhangsan"
        ]
    },
    "post_form":{
        "password":[
            "123456"
        ],
        "user_name":[
            "admin"
        ]
    },
    "multipart_form":null,
    "body":""
}
```

### POST请求带参数并上传文件

```
curl -X POST --form user_name=admin --form password=123456 \
--form "upload[]=@/Users/mac/Downloads/csb-db-2019-10-24.txt" \
--form "upload[]=@/Users/mac/Downloads/QQ20191028-161857.png" \
http://127.0.0.1:8080/test?user_name=zhangsan
```

有文件上传的表单:`Content-Type: multipart/form-data`

```
{
    "query_form":{
        "user_name":[
            "zhangsan"
        ]
    },
    "form":{
        "password":[
            "123456"
        ],
        "user_name":[
            "zhangsan",
            "admin"
        ]
    },
    "post_form":{
        "password":[
            "123456"
        ],
        "user_name":[
            "admin"
        ]
    },
    "multipart_form":{
        "Value":{
            "password":[
                "123456"
            ],
            "user_name":[
                "admin"
            ]
        },
        "File":{
            "upload[]":[
                {
                    "Filename":"csb-db-2019-10-24.txt",
                    "Header":{
                        "Content-Disposition":[
                            "form-data; name="upload[]"; filename="csb-db-2019-10-24.txt""
                        ],
                        "Content-Type":[
                            "text/plain"
                        ]
                    },
                    "Size":189
                },
                {
                    "Filename":"QQ20191028-161857.png",
                    "Header":{
                        "Content-Disposition":[
                            "form-data; name="upload[]"; filename="QQ20191028-161857.png""
                        ],
                        "Content-Type":[
                            "application/octet-stream"
                        ]
                    },
                    "Size":159513
                }
            ]
        }
    },
    "body":""
}
```

### JSON数据

```
curl -X POST --header 'content-type: application/json' \
--data '{"user_name":"admin","password":"123456"}' \
http://127.0.0.1:8080/test?user_name=zhangsan
```

```
{
    "query_form":{
        "user_name":[
            "zhangsan"
        ]
    },
    "form":{
        "user_name":[
            "zhangsan"
        ]
    },
    "post_form":{

    },
    "multipart_form":null,
    "body":"{"user_name":"admin","password":"123456"}"
}
```

### XML数据

```
curl -X POST --header "content-type: application/xml" --data \
'<?xml version="1.0" encoding="UTF-8"?><root>
  <user_name>admin</user_name>
  <password>123456</password>
</root>' \
http://127.0.0.1:8080/test?user_name=zhangsan
```

```
{
    "query_form":{
        "user_name":[
            "zhangsan"
        ]
    },
    "form":{
        "user_name":[
            "zhangsan"
        ]
    },
    "post_form":{

    },
    "multipart_form":null,
    "body":"<?xml version="1.0" encoding="UTF-8"?><root>
  <user_name>admin</user_name>
  <password>123456</password>
</root>"
}
```

## About

* **作者**：March
* **邮箱**：fengqi.mao.march@gmail.com
* **头条**：https://toutiao.io/u/425956/subjects
* **简书**：https://www.jianshu.com/u/02f2491c607d
* **掘金**：https://juejin.im/user/5b484473e51d45199940e2ae
* **CSDN**：http://blog.csdn.net/u011810138
* **SegmentFault**：https://segmentfault.com/u/maoqiqi
* **StackOverFlow**：https://stackoverflow.com/users/8223522

## License

```
   Copyright 2019 maoqiqi

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```