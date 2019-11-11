# brew

https://brew.sh/

安装 Homebrew

`/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`

brew 是 Mac 下的一个包管理工具，类似于 centos 下的 yum，可以很方便地进行安装/卸载/更新各种软件包。


基本用法

安装/卸载/更新

以 nodejs 为例，执行下面命令即可，安装目录在 /usr/local/Cellar

`brew install nodejs`

如果需要更新或卸载

```
brew upgrade nodejs
brew remove nodejs
```

其他命令

```
brew list                   # 列出当前安装的软件
brew search nodejs          # 查询与 nodejs 相关的可用软件
brew info nodejs            # 查询 nodejs 的安装信息
```


Mac安装brew并更改源

在Mac上安装brew时，如果使用官方推荐的方式，会耗费很长时间，并且也不一定能成功。

将安装源换成国内源。

1.将brew的install文件下载本地

`curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install >> brew_install`

2.修改install文件的镜像源

`vim brew_install`

将brew_install文件里面的两行代码替换掉，待替换的代码为：

```
BREW_REPO = "https://github.com/Homebrew/brew".freeze
CORE_TAP_REPO = "https://github.com/Homebrew/homebrew-core".freeze
```

替换为：

```
BREW_REPO = "git://mirrors.ustc.edu.cn/brew.git".freeze
CORE_TAP_REPO = "git://mirrors.ustc.edu.cn/homebrew-core.git".freeze
```

修改完成之后保存好修改后的brew_install文件，并退出。


安装brew

`/usr/local/bin/ruby ~/brew_install`

