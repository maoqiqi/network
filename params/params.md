### 1.GET请求 

`curl http://127.0.0.1:8080/test`

```
{
    "query_form":{

    },
    "form":{

    },
    "post_form":{

    },
    "multipart_form":null
}
```

### 2.GET请求带参数 

`curl http://127.0.0.1:8080/test?user_name=admin&password=123456`

```
{
    "query_form":{
        "password":[
            "123456"
        ],
        "user_name":[
            "admin"
        ]
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

    },
    "multipart_form":null
}
```

### 3.POST请求带参数

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
    "multipart_form":null
}
```

### 4.POST请求带参数（地址也带参数）

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
    "multipart_form":null
}
```

### 5.POST请求带参数并上传文件

```
curl -X POST --form user_name=admin --form password=123456 \
--form "upload[]=@/Users/mac/Downloads/csb-db-2019-10-24.txt" \
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
                }
            ]
        }
    }
}
```