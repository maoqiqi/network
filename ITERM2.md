# iTerm2 + Oh My Zsh 打造舒适终端体验

https://www.iterm2.com/

安装完成后，在/bin目录下会多出一个zsh的文件。

Mac系统默认使用dash作为终端，可以使用命令修改默认使用zsh：

```
chsh -s /bin/zsh
```

如果想修改回默认dash，同样使用chsh命令即可：

chsh -s /bin/bash

查看系统当前使用的shell

echo $SHELL

查看系统是否安装了zsh

```
cat /etc/shells
# List of acceptable shells for chpass(1).
# Ftpd will not allow users to connect who are not using
# one of these shells.

/bin/bash
/bin/csh
/bin/ksh
/bin/sh
/bin/tcsh
/bin/zsh
```


安装Oh my zsh

zsh的功能极其强大，只是配置过于复杂,通过Oh my zsh可以很快配置zsh。

安装方法有两种，可以使用curl或wget，看自己环境或喜好：

```
# curl 安装方式
sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
```

```
# wget 安装方式
sh -c "$(wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)"
```

卸载oh my zsh，在命令行输入如下命令，回车即可：uninstall_oh_my_zsh



安装Powerline

在官网有教程，我们只需要执行官网第一条安装指令就行

如果你的终端能够正常执行pip指令，那么直接执行下面的指令可以完成安装

```pip install powerline-status```

如果没有，则先执行安装pip指令

```sudo easy_install pip```


安装PowerFonts

```
git clone https://github.com/powerline/fonts.git --depth=1
cd fonts
./install.sh
```

安装好字体库之后，我们来设置iTerm2的字体，具体的操作是:

iTerm2 -> Preferences -> Profiles -> Text

在Font区域选中Change Font，然后找到Meslo LG字体。


背景图：http://wx1.sinaimg.cn/large/81f8a509gy1fnjdvkkwgoj20zk0m8ak8.jpg

安装配色方案(可跳过)

在常用的位置新建一个文件夹OpenSource.

```
在OpenSource目录下执行git clone命令:
git clone https://github.com/altercation/solarized
cd solarized/iterm2-colors-solarized/
open .
```

在打开的finder窗口中，双击Solarized Dark.itermcolors和Solarized Light.itermcolors即可安装明暗两种配色：

再次进入iTerm2 -> Preferences -> Profiles -> Colors -> Color Presets中根据个人喜好选择.


安装主题

```
在OpenSource目录下执行git clone命令:
git clone https://github.com/fcamblor/oh-my-zsh-agnoster-fcamblor.git
cd oh-my-zsh-agnoster-fcamblor/
./install
```

执行上面的命令会将主题拷贝到oh my zsh的themes.

执行命令打开zshrc配置文件，将ZSH_THEME后面的字段改为agnoster


安装高亮插件

```
这是oh my zsh的一个插件，安装方式与theme大同小异：
cd ~/.oh-my-zsh/custom/plugins/
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git
vi ~/.zshrc
```

这时我们再次打开zshrc文件进行编辑。找到plugins，此时plugins中应该已经有了git，我们需要把高亮插件也加上：

```
plugins=(git composer docker docker-compose osx zsh-autosuggestions zsh-syntax-highlighting)
```

请务必保证插件顺序，zsh-syntax-highlighting必须在最后一个。
然后在文件的最后一行添加：

```
source ~/OpenSource/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
```

```
执行命令使刚才的修改生效：
source ~/.zshrc
```

可选择、命令补全

跟代码高亮的安装方式一样，这也是一个zsh的插件，叫做zsh-autosuggestion，用于命令建议和补全。

```
cd ~/.oh-my-zsh/custom/plugins/
git clone https://github.com/zsh-users/zsh-autosuggestions
vi ~/.zshrc
```

```
plugins=(git composer docker docker-compose osx zsh-autosuggestions zsh-autosuggestion zsh-syntax-highlighting)
```

有同学说补全命令的字体不太清晰，与背景颜色太过相近，其实可以自己调整一下字体颜色。

Preferences -> Profiles -> Colors 中有Foreground是标准字体颜色，ANSI Colors中Bright的第一个是补全的字体颜色。




Zsh和Bash，究竟有何不同.坑很深