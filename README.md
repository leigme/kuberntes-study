# kuberntes-study
### 搭建环境
1. 备份源
```shell
cp /etc/apt/sources.list ~/sources.list.bak
```

2. 编辑源
```shell
sudo vim /etc/apt/sources.list
# 默认注释了源码仓库，如有需要可自行取消注释
deb https://mirrors.ustc.edu.cn/ubuntu/ jammy main restricted universe multiverse
# deb-src https://mirrors.ustc.edu.cn/ubuntu/ jammy main restricted universe multiverse

deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-security main restricted universe multiverse
# deb-src https://mirrors.ustc.edu.cn/ubuntu/ jammy-security main restricted universe multiverse

deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-updates main restricted universe multiverse
# deb-src https://mirrors.ustc.edu.cn/ubuntu/ jammy-updates main restricted universe multiverse

deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse
# deb-src https://mirrors.ustc.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse

# 预发布软件源，不建议启用
# deb https://mirrors.ustc.edu.cn/ubuntu/ jammy-proposed main restricted universe multiverse
# deb-src https://mirrors.ustc.edu.cn/ubuntu/ jammy-proposed main restricted universe multiverse
sudo apt update
```
3. 安装docker
```shell
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo service docker start
```
4. 将当前用户添加到docker用户组
```shell
sudo usermod -aG docker $USER
```
5. 修改docker访问镜像，可以使用自己的阿里云镜像
```shell
sudo vim /etc/docker/daemon.json
{
  "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]
}
sudo service docker stop
sudo service docker start
```
6. 安装kubectl
```shell
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
```
输出md5比对结果
```shell
kubectl: OK
```
输入安装命令
```shell
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```
7. 安装minikube
```shell
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```
8. 启动minikube容器
```shell
minikube start --driver=docker --image-mirror-country='cn' --kubernetes-version=v1.23.13
```
9. 安装golang开发环境
```shell
sudo apt install software-properties-common -y
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go -y
```
10. 配置golang环境变量
```shell
vim ~/.bash_profile
export GOPROXY=https://proxy.golang.com.cn,direct
export PATH=$GOPATH/bin:$PATH
```
11. windows环境wsl的添加自动启动
```shell
sudo vim /etc/init.wsl
service docker start
sudo chmod +x /etc/init.wsl
```
12. 关闭wsl，在windows添加启动任务
按win+r打开快速启动，输入
```shell
shell:startup
```
进入启动任务文件夹，创建linux-start.vbs
```shell
Set ws = WScript.CreateObject("WScript.Shell")
ws.run "wsl -d Ubuntu-22.04 -u root /etc/init.wsl"
```
13. 备份Ubuntu-22.04子系统
查看已经安装的版本
```shell
wsl -l
```
14. 在子系统关闭的状态下输入
```shell
wsl --export Ubuntu-22.04 ubuntu-22.04.tar
```
导入命令
```shell
wsl --import ubuntu-22.04.tar
```