# Mac开发配置手册

《Mac开发配置手册》如何让一部全新的MacBook快速完成开发环境配置。


## 目录

* [Mac开发前需做的系统设置](#Mac开发前需做的系统设置)
* [必装软件清单](#必装软件清单)
* [Sudo免密码设置](#Sudo免密码设置)
* [临时上传下载文件](#临时上传下载文件)
* [About](#About)
* [License](#License)

## Mac开发前需做的系统设置

### 系统设置

在任何的操作系统中，首先你需要做一件事就是更新系统，点击窗口左上角的  > 关于本机 > 软件更新。
此外，如果这是一部新的电脑，你还需要到系统设置进行一些适当调整。如何调整，取决于个人喜好。

* 触控板(系统设置 > 触控板)
  * [ ] 光标与点击
  * [x] 轻拍来点按
  * [x] 辅助点按
  * [x] 查找
  * [x] 三指拖移
  * [ ] 滚动缩放
  * [x] 默认全选
  * [ ] 更多手势
  * [x] 默认全选
* Dock
  * 置于屏幕上的位置：底部
  * 设置Dock图标更小(大小随个人喜好)
  * ✓ 自动显示和隐藏Dock
* Finder
  * Finder > 显示
  * Finder > 偏好设置
    * 通用:开启新Finder窗口时打开:HOME「用户名」目录
    * 边栏:添加HOME用户名」目录和创建代码文件目录,将共享的(shared)和标记(tags)目录去掉
* 菜单栏
  * 去掉蓝牙等无需经常使用的图标
  * 将电池显示设置为百分比

* Spotlight
  * 去掉字体和书签与历史记录等不需要的内容
  * 设置合适的快捷键

* 互联网帐户
  * 添加iCloud用户，同步日历，联系人和 Find my mac 等等。
  

## 必装软件清单

|软件|版本|说明|
|:-----|:-----|:-----|
|brew|`*`|https://brew.sh/|
|iTerm2|`*`|https://iterm2.com/|
|ohmyzsh|`*`|https://ohmyz.sh/|
|git|`*`|`brew install git`|
|redis|`>3.*`|`brew install redis`|
|nginx|`>1.10`|`brew install nginx` 如果安装docker环境可不装|
|sequel pro|`*`|http://sequelpro.com/ 数据库管理工具|
|switchhosts|`*`|https://github.com/oldj/SwitchHosts host管理工具|
|The Unarchiver|`*`|https://theunarchiver.com/ 解压工具|


## 如何安装Homebrew与使用

Homebrew:包管理工具可以让你安装和更新程序变得更方便，是目前在OS X系统中最受欢迎的包管理工具。
类似于centos下的yum，可以很方便地进行安装/卸载/更新各种软件包。

官网:https://brew.sh/

### 安装

`/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`

#### 安装并更改源

在Mac上安装brew时，如果使用官方推荐的方式，会耗费很长时间，并且也不一定能成功。

将安装源换成国内源

* 将brew的install文件下载本地
  `curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install >> brew_install`
* 修改install文件的镜像源
  `vim brew_install`
* 将brew_install文件里面的两行代码替换掉

  待替换的代码为：
  
  ```
  BREW_REPO = "https://github.com/Homebrew/brew".freeze
  CORE_TAP_REPO = "https://github.com/Homebrew/homebrew-core".freeze
  ```
  
  替换为：
  
  ```
  BREW_REPO = "git://mirrors.ustc.edu.cn/brew.git".freeze
  CORE_TAP_REPO = "git://mirrors.ustc.edu.cn/homebrew-core.git".freeze
  ```
  
  修改完成之后保存好修改后的brew_install文件。
  
* 安装
  `/usr/local/bin/ruby ~/brew_install`

### 基本使用

* `brew update`:更新Homebrew在服务器端上的包目录
* `brew install <package_name>`:安装
* `brew upgrade <package_name>`:更新
* `brew remove`:卸载
* `brew outdated`:查看你的包是否需要更新
* `brew cleanup`:Homebrew将会把老版本的包缓存下来，以便当你想回滚至旧版本时使用。但这是比较少使用的情况，当你想清理旧版本的包缓存时，可以运行。
* `brew list`:列出当前安装的软件
* `brew list --versions`:列出当前安装的软件(包括版本号)
* `brew search <package_name>`:查询与`package_name`相关的可用软件
* `brew info <package_name>`:查询`package_name`的安装信息

### Cask

你已经感受到了使用 Homebrew 安装命令行程序的便利。
那么接下来，我们将通过Homebrew Cask优雅、简单、快速的安装和管理OS X图形界面程序，比如Google Chrome和Dropbox。


## Sudo免密码设置

* 打开命令窗口输入如下命令：`sudo visudo`
* 替换 #%username ALL=(ALL) ALL 为:`%username ALL=(ALL) NOPASSWD: ALL`

> username为当前登录用户名


## 临时上传下载文件

```
curl -F "file=@test.txt" https://file.io
{"success":true,"key":"2ojE41","link":"https://file.io/2ojE41","expiry":"14 days"}

curl https://file.io/2ojE41
This is a test

curl https://file.io/2ojE41
{"success":false,"error":404,"message":"Not Found"}
```

设置到期时间

```
curl -F "file=@test.txt" https://file.io/?expires=1w
{"success":true,"key":"aQbnDJ","link":"https://file.io/aQbnDJ","expiry":"7 days"}

sleep 604801

curl https://file.io/aQbnDJ
{"success":false,"error":404,"message":"Not Found"}
```

> 查询参数过期必须是一个正整数，默认情况下，它表示删除文件之前的天数（默认为14天）。如果跟随w，则为周数。m代表数月，y代表数年。

还可以将直接文本发送到`file.io`:

```
curl --data "text=this is a secret pw" https://file.io
{"success":true,"key":"pgiPc2","link":"https://file.io/pgiPc2","expiry":"14 days"}

$ curl https://file.io/pgiPc2
this is a secret pw

$ curl https://file.io/pgiPc2
{"success":false,"error":404,"message":"Not Found"}
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