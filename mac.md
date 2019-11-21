# Mac开发配置手册

《Mac开发配置手册》如何让一部全新的MacBook快速完成开发环境配置。


## 目录

* [Mac开发前需做的系统设置](#Mac开发前需做的系统设置)
* [必装软件清单](#必装软件清单)
* [brew](#brew)
* [iTerm2](#iTerm2)
* [ohmyzsh](#ohmyzsh)
  * [编辑zshrc文件](#编辑zshrc文件)
  * [安装配色方案](#安装配色方案)
  * [安装代码高亮插件](#安装代码高亮插件)
  * [安装代码补全插件](#安装代码补全插件)
* [git](#git)
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

### 开发工具

|软件|版本|说明|
|:-----|:-----|:-----|
|[brew](https://brew.sh/)|`*`||
|[iTerm2](https://iterm2.com/)|`*`||
|[ohmyzsh](https://ohmyz.sh/)|`*`||
|[git]|`*`|`brew install git`|
|[redis]|`>3.*`|`brew install redis`|
|[nginx]|`>1.10`|`brew install nginx`如果安装docker环境可不装|
|[sequelpro](http://sequelpro.com/)|`*`|数据库管理工具|
|[SwitchHosts](https://github.com/oldj/SwitchHosts)|`*`|host管理工具|
|[Google Chrome](https://google.com/intl/en/chrome/)|`*`|`brew cask install google-chrome`|
|[Sketch](https://www.sketch.com/)|`*`||
|[Dash](https://kapeli.com/dash)|`*`||
|[Parallels](https://www.parallels.com/hk)|`*`|/
|[Github](https://desktop.github.com/)|`*`||

### 生产力工具

|软件|版本|说明|
|:-----|:-----|:-----|
|[The Unarchiver](https://theunarchiver.com/)|`*`|支持多种格式(包括Windows下的格式)的压缩/解压缩工具|
|[1Password](https://blog.1password.com/accel-partnership/)||跨平台的密码管理工具|
|[Alfred 2](https://www.alfredapp.com/)||搜索工具|
|[AppCleaner](http://freemacsoft.net/appcleaner/)||应用程序卸载工具|
|[Dropbox](https://www.dropbox.com/)||文件同步工具|
|[Reeder](https://reederapp.com/)||RSS阅读工具|
|[Pocket](https://getpocket.com/)||稍后阅读工具|
|[Spectacle](https://www.spectacleapp.com/)||让窗口成比例的显示，在写代码调试的时候很方便|
|[OminiFocus](https://www.omnigroup.com/omnifocus)||时间管理工具|
|[Mou](http://25.io/mou/)||Markdown编辑器,国人出品|

### 其它工具

|软件|版本|说明|
|:-----|:-----|:-----|
|CheatSheet||http://www.grandtotal.biz/cheatsheet/ 长按Command,将能查看当前程序的快捷键|


## brew

Homebrew:包管理工具可以让你安装和更新程序变得更方便，是目前在OS X系统中最受欢迎的包管理工具。
类似于centos下的yum，可以很方便地进行安装/卸载/更新各种软件包。

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

安装Homebrew-cask是如此的简单直接，运行以下命令即可完成：

```
// 添加Github上的caskroom/cask库
brew tap caskroom/cask  
// 安装brew-cask
brew install brew-cask  
// 安装Google浏览器
brew cask install google-chrome 
// 更新
brew update && brew upgrade brew-cask && brew cleanup
```

如果你想查看cask上是否存在你需要的app，可以到[caskroom.io](https://buyinstagramlikes.io/caskroom/)进行搜索。

**文件预览插件**

有些插件可以让Mac上的文件预览更有效，比如语法高亮、markdown 渲染、json 预览等等。

```
brew cask install qlcolorcode
brew cask install qlstephen
brew cask install qlmarkdown
brew cask install quicklook-json
brew cask install qlprettypatch
brew cask install quicklook-csv
brew cask install betterzipql
brew cask install webp-quicklook
brew cask install suspicious-package  
```

OS X图形界面程序

```
brew cask install alfred
brew cask install appcleaner
brew cask install cheatsheet
brew cask install dropbox
brew cask install google-chrome
brew cask install onepassword
brew cask install sublime-text
brew cask install totalfinder
```

> 如果你经常使用Alfred启动程序，那么你会想Alfred可以搜索brew cask安装的程序，实现这些仅需运行：

```
brew cask alfred link
```

此外你还可以通过brew cask安装[launchrocket](https://github.com/jimbojsb/launchrocket)，来管理通过brew安装的service。

![launchrocket](images/launchrocket.png)


## iTerm2

作为一名开发者，我们常常花上很多时间在终端上，如同武士的剑，一出手便知高低。所以让我们安装Mac上最强大的终端iTerm2吧！

* 在`Keys -> Hotkey`中设置`command + option + i`快速显示和隐藏iTerm
* 在`Profiles -> Default -> Check silence bell`
* 下载`Solarized dark iterm colors`，在`Profiles -> Default -> Colors -> Load Presets`将其导入，作为默认颜色。
* 在`Profiles -> Text`改变游标(cursor)文字和颜色，随个人喜好。


## ohmyzsh

安装zsh，zsh的功能极其强大，只是配置过于复杂,通过oh-my-zsh可以很快配置zsh。其中Env.sh文件用于维护别名(aliases)，输出(exports)和路径改变(path changes)等等，以免影响`~/.zshrc`。

使用brew完成zsh和zsh completions的安装:`brew install zsh zsh-completions`

> Mac系统默认使用dash作为终端，可以使用命令修改默认使用zsh:

```
chsh -s /bin/zsh
```

如果想修改回默认dash，同样使用chsh命令即可：

```
chsh -s /bin/bash
```

> 查看系统当前使用的shell:`echo $SHELL`

> 查看系统是否安装了zsh:`cat /etc/shells`

安装oh-my-zsh方法有两种，可以使用curl或wget，看自己环境或喜好：

```
# curl 安装方式
sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
```

```
# wget 安装方式
sh -c "$(wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"
```

卸载oh-my-zsh，在命令行输入如下命令，回车即可:`uninstall_oh_my_zsh`

### 编辑zshrc文件

```
# 设置主题
ZSH_THEME="agnoster"

# 设置别名
alias zshconfig="vim ~/.zshrc"
alias envconfig="vim ~/.env"
alias ohmyzsh="vim ~/.oh-my-zsh"  

# 隐藏"user@hostname"
prompt_context() {}
```

> 注意：许多主题需要安装[Powerline Fonts](https://github.com/powerline/fonts)才能正确呈现。

```
# clone
git clone https://github.com/powerline/fonts.git --depth=1
# install
cd fonts
./install.sh
# clean-up a bit
cd ..
rm -rf fonts
```

安装好字体库之后，我们来设置iTerm2的字体，具体的操作是:
打开`iTerm2 -> Preferences -> Profiles -> Text`,在Font区域选中`Change Font`，然后找到`Meslo`字体。

选择`Meslo`字体。

### 安装配色方案

```
cd ~/.oh-my-zsh/custom/plugins
git clone https://github.com/altercation/solarized
cd solarized/iterm2-colors-solarized/
open .
```

在打开的finder窗口中，双击`Solarized Dark.itermcolors`和`Solarized Light.itermcolors`即可安装明暗两种配色：

打开`iTerm2 -> Preferences -> Profiles -> Colors -> Color Presets`中根据个人喜好选择。

### 安装代码高亮插件

```
cd ~/.oh-my-zsh/custom/plugins
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git
vim ~/.zshrc
```

这时我们再次打开zshrc文件进行编辑。找到plugins，此时plugins中应该已经有了git，我们需要把高亮插件也加上：

```
plugins=(git zsh-syntax-highlighting)
```

> 请务必保证插件顺序，`zsh-syntax-highlighting`必须在最后一个。

执行命令使刚才的修改生效:`source ~/.zshrc`

### 安装代码补全插件

跟代码高亮的安装方式一样，这也是一个zsh的插件，叫做`zsh-autosuggestion`，用于命令建议和补全。

```
cd ~/.oh-my-zsh/custom/plugins/
git clone https://github.com/zsh-users/zsh-autosuggestions
vim ~/.zshrc
```

同上找到plugins，把代码补全插件加上：

```
plugins=(git zsh-autosuggestions zsh-syntax-highlighting)
```

执行命令使刚才的修改生效:`source ~/.zshrc`

> 推荐一张背景图:`http://wx1.sinaimg.cn/large/81f8a509gy1fnjdvkkwgoj20zk0m8ak8.jpg`

![背景图](http://wx1.sinaimg.cn/large/81f8a509gy1fnjdvkkwgoj20zk0m8ak8.jpg)


## git

作为一名开发者怎么可能没有Git呢? 我们马上就来安装：

```
brew install git
```

好的，现在我们来测试一下Git是否安装完好：

```
git --version
```

接着，我们将定义你的Git帐号(与你在GitHub使用的用户名和邮箱一致)

```
git config --global user.name "Your Name Here"
git config --global user.email "your_email@youremail.com"  
```

这些配置信息将会添加进`~/.gitconfig`文件中。

推荐使用HTTPS方法(另一个是SSH)，将你的代码推送到Github上的仓库。
如果你不想每次都输入用户名和密码的话，可以按照此[描述](https://help.github.com/en/github/getting-started-with-github/set-up-git)说的那样，运行：

```
git config --global credential.helper osxkeychain
```

此外，如果你打算使用 SSH 方式，可以参考此[链接](https://help.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh).


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

curl https://file.io/pgiPc2
this is a secret pw

curl https://file.io/pgiPc2
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