# java
export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk-10.0.2.jdk/Contents/Home
export PATH=$JAVA_HOME/bin:$PATH
export CLASSPATH=$JAVA_HOME/lib/tools.jar:$JAVA_HOME/lib/dt.jar:.

# android
export ANDROID_HOME=$HOME/Library/Android/sdk
export PATH=$ANDROID_HOME/tools:$ANDROID_HOME/platform-tool:$PATH

# export /Applications/Android\ Studio.app/Contents/gradle/gradle-4.6
# export PATH=$GRADLE_HOME/bin:$PATH

# flutter
export PUB_HOSTED_URL=https://pub.flutter-io.cn
export FLUTTER_STORAGE_BASE_URL=https://storage.flutter-io.cn
export PATH=$HOME/flutter/flutter/bin:$PATH

# go
export PATH=/usr/local/go/bin:$PATH
export GOPROXY=direct
export GOPRIVATE=git.verystar.cn
export GO111MODULE=on

gop(){
	currpath=`pwd`/
    gopath=${currpath%src/*}

	if [[ ${gopath} = "" ]];then
		echo "path not found src"
	else
        export GOPATH=${currpath%src/*}
		echo current ${GOPATH}
    fi
}

# proxy
proxy () {
    export http_proxy=http://127.0.0.1:1087
    export https_proxy=http://127.0.0.1:1087
    export all_proxy=http://127.0.0.1:1087
    export NO_PROXY="git.verystar.cn"
    echo "Proxy on"
}

unproxy () {
    unset http_proxy
    unset https_proxy
    unset all_proxy
    echo "Proxy off"
}

ungoproxy () {
    unset GOPROXY
    unset GONOPROXY
    echo "GOPROXY OFF"
}