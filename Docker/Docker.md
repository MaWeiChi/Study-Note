# Docker Container



## 1. Concept


Docker image : 映像檔，唯讀的模板，單獨運行的輕量級套件，包含所有執行程式所需要的函式庫、環境變數與設定檔等，可用來建立容器。

Docker container : image 的執行實體，將映像檔載入至記憶體中執行之後的環境，可視為一個簡易的 Linux ；可以被啟動、開始、停止、刪除。

Docker Repository：集中存放映像檔的位置，分為公開倉庫及私有倉庫。最大的公開倉庫為 [Docker Hub](https://hub.docker.com/) 。使用者可在本地端建立私有倉庫。

#### Docker vs VM

VM 相較於 Docker 包含程式及函式庫外，還需使用作業系統（OS）。

VM架構

![image](https://github.com/EricMa19920728/study-note/blob/master/Picture/virtual-machine.jpg)


Docker container 只需要包含程式及函式庫，所有程式共用一個 Host OS 系統核心

Docker架構

![image](https://github.com/EricMa19920728/study-note/blob/master/Picture/docker-container-20170625-1%20(1).jpg)

由上敘可知，對比於 VM 需要維持額外的系統資源，Docker 在使用實體硬體資源效率更好。

以 配置 ８ GB Ram為例，假設原生 OS 使用 2GB 維持，一台 VM 需要 1 GB 維持前提，在使用兩台 VM 後，可使用資源僅剩 4 GB 可用，若使用 Docker 則有 6 GB 資源可用。

而 Docker 在啟動耗費時間上僅需要秒級內，VM 則需先行完成系統啟動的部分。


#### Summary

| 特性       | 容器               | 虛擬機     |
| ---------- | ------------------ | ---------- |
| 啟動       | 秒級               | 分鐘級     |
| 硬碟容量   | 一般為 MB          | 一般為 GB  |
| 效能       | 接近原生           | 比較慢     |
| 系統支援量 | 單機支援上千個容器 | 一般幾十個 |


## 2. Install 

### On Mac
Docker 官網 : https://docs.docker.com/docker-for-mac/install/

Homebrew :

```
brew cask install docker
```



## 3. Image

### Get image：

```
$ docker pull jupyter/base-notebook
Using default tag: latest
latest: Pulling from jupyter/base-notebook
5c939e3a4d10: Downloading [=========>                                         ]  4.992MB/26.69MB
c63719cdbe7a: Download complete 
19a861ea6baf: Download complete 
651c9d2d6c4f: Download complete 
21b673dc817c: Downloading [=====================================>             ]  7.085MB/9.392MB
1594017be8ef: Download complete 
b392f2c5ed42: Download complete 
8e4f6538155b: Download complete 
0e7a6279e2b7: Download complete 
d36911bea109: Download complete 
94cddf08e6e9: Waiting 
3d6a3f4b6e39: Waiting 
bdb10ba740a5: Waiting 
558b32bec68b: Pulling fs layer 
7b7baa387b50: Waiting 
989d432af0d4: Waiting 
bf2a412948a6: Waiting 
156a8789d9d4
```

```
$ docker pull jupyter/base-notebook
Using default tag: latest
latest: Pulling from jupyter/base-notebook
5c939e3a4d10: Pull complete 
c63719cdbe7a: Pull complete 
19a861ea6baf: Pull complete 
651c9d2d6c4f: Pull complete 
21b673dc817c: Pull complete 
1594017be8ef: Pull complete 
b392f2c5ed42: Pull complete 
8e4f6538155b: Pull complete 
0e7a6279e2b7: Pull complete 
d36911bea109: Pull complete 
94cddf08e6e9: Pull complete 
3d6a3f4b6e39: Pull complete 
bdb10ba740a5: Pull complete 
558b32bec68b: Pull complete 
7b7baa387b50: Pull complete 
989d432af0d4: Pull complete 
bf2a412948a6: Pull complete 
156a8789d9d4: Pull complete 
Digest: sha256:a4b03739ac37a0ccc38d71779c236f82cae342936be81d5c9d677497a8a84dfe
Status: Downloaded newer image for jupyter/base-notebook:latest
docker.io/jupyter/base-notebook:latest
```

`docker pull`會從預設的 Docker Hub 下載映像檔。在下載完成後以 Pull complete 表示。



### List images：

```
$ docker images 
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
jupyter/base-notebook   latest              5aef07718da8        11 days ago         599MB
ubuntu                  latest              ccc6e87d482b        3 weeks ago        64.2MB
```

Options：https://docs.docker.com/engine/reference/commandline/images/

REPOSITORY：映像檔的倉庫來源，比如 ubuntu

TAG：映像檔的標記，比如 14.04、latest

IMAGE ID：映像檔的`ID` 號，使用 sha 256 編碼

CREATED：映像檔的建立時間

SIZE：映像檔的大小

映像檔的 `ID` 是唯一的，擁有相同 `ID` 表示是同一個映像檔。`TAG` 用來區分同一個倉庫內不同映像檔，例如 `ubuntu` 倉庫中有多個映像檔，通過 `TAG` 來區分發行版本，例如 `10.04`、`12.04`、`12.10`、`13.04`、`14.04` 等。



### Build Image：

1. `docker commit`

```
$ sudo docker run -t -i ubuntu /bin/bash
root@c59c0fc4d88c:/#
```

```
root@c59c0fc4d88c:/# python3 --version
bash: python3: command not found

root@c59c0fc4d88c:/# apt update
...

root@c59c0fc4d88c:/# apt install python3
...

root@c59c0fc4d88c:/# apt install python3-pip
...

root@c59c0fc4d88c:/# python3 --version
Python 3.6.9

root@c59c0fc4d88c:/# exit
exit

$ docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                     PORTS               NAMES
c59c0fc4d88c        ubuntu              "/bin/bash"         3 minutes ago       Exited (0) 3 seconds ago                       funny_tharp
```

```
$ sudo docker commit -m "Added python3" -a "Erik Ma" c59c0fc4d88c ubuntu_with_python
sha256:ce7a618e242a2427c13d99ad6719bcb3224030737c7c659c7e10bc1b1fd29297

$ docker images
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
ubuntu_with_python      latest              ce7a618e242a        14 seconds ago      471MB
jupyter/base-notebook   latest              5aef07718da8        11 days ago         599MB
ubuntu                  latest              ccc6e87d482b        3 weeks ago         64.2MB

$erikwc_ma$ sudo docker run -t -i ubuntu_with_python /bin/bash
root@c2ea9fba3954:/# python3 --version
Python 3.6.9
```



2. Dockerfile 

```
$ mkdir -p Docker/Dockerfile/MyJupyter

$ cd Docker/Dockerfile/MyJupyter

$ touch Dockerfile

$ nano DOckerfile
```

```
# Dockerfile build a Jupyter notebook image with python3 on 2020.02.11

FROM ubuntu_with_python

MAINTAINER ErikMa <erik24ma@gmail.com>

RUN apt-get -qq update

RUN pip3 install jupyter
```

```
$ sudo docker build -t="myjupyter" .
...
Removing intermediate container 9227202045f9
 ---> c6fe90b44045
Successfully built c6fe90b44045
Successfully tagged myjupyter:latest-

$ docker images
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
myjupyter               latest              c6fe90b44045        18 seconds ago      572MB
ubuntu_with_python      latest              ce7a618e242a        7 minutes ago       471MB
jupyter/base-notebook   latest              5aef07718da8        11 days ago         599MB
ubuntu                  latest              ccc6e87d482b        3 weeks ago         64.2MB
```



### Remove images

```
$ docker images
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
myjupyter               latest              3172525b81f5        About an hour ago   574MB
ubuntu_with_python      latest              ce7a618e242a        2 hours ago         471MB
<none>                  <none>              6baef3f018ce        3 hours ago         130MB
jupyter/base-notebook   latest              5aef07718da8        12 days ago         599MB
ubuntu                  latest              ccc6e87d482b        3 weeks ago         64.2MB
doc

$ docker rmi 6baef3f018ce
Deleted: sha256:6baef3f018cefc0fc0c914f965102dd9142fa319e762a5e26393c1724c031468
Deleted: sha256:c7e965a60ed3bb44c1bfed4839038d6a9754a6363bf1a69fc59b3c7be114bb1d
Deleted: sha256:ce4891da156090a9312bf5d972b78c5ddcbeafd59b8b84e06c25aec31e21c797

$ docker images
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
myjupyter               latest              3172525b81f5        About an hour ago   574MB
ubuntu_with_python      latest              ce7a618e242a        2 hours ago         471MB
jupyter/base-notebook   latest              5aef07718da8        12 days ago         599MB
ubuntu                  latest              ccc6e87d482b        3 weeks ago         64.2MB

```

*注意：在刪除映像檔之前要先用 `docker rm` 刪掉依賴於這個映像檔的所有容器。



## Container 

### Launch container :

1. Run a new container :

   ```
   $ docker run ubuntu /bin/echo 'Hello world'
   Hello world
   
   $ docker run -t -i  ubuntu /bin/bash
   root@636fe0e11482:/#
   
   root@636fe0e11482:/# pwd
   /
   root@636fe0e11482:/# ls
   bin boot dev etc home lib lib64 media mnt opt proc root run sbin srv sys tmp usr var
   
   root@636fe0e11482:/# exit
   exit
   ```

   - 檢查本地是否存在指定的映像檔，不存在就從公有倉庫下載
   - 利用映像檔建立並啟動一個容器
   - 分配一個檔案系統，並在唯讀的映像檔層外面掛載一層可讀寫層
   - 從宿主主機設定的網路橋界面中橋接一個虛擬埠到容器中去
   - 從位址池中設定一個 ip 位址給容器
   - 執行使用者指定的應用程式
   - 執行完畢後容器被終止

2. Start a stopped container :

   ```
   $ docker ps -a
   CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                     PORTS               NAMES
   a3df94a31a25        ubuntu              "/bin/bash"         2 minutes ago       Exited (0) 6 seconds ago                       gifted_rubin
   
   $ docker start a3df94a31a25
   a3df94a31a25
   
   $ docker ps
   CONTAINER ID        IMAGE               COMMAND             CREATED              STATUS              PORTS               NAMES
   a3df94a31a25        ubuntu              "/bin/bash"         About a minute ago   Up 5 seconds                            gifted_rubin
   
   $ docker attach a3df94a31a25
   root@a3df94a31a25:/#
   
   root@a3df94a31a25:/#exit
   exit
   
   $ docker start -a a3df94a31a25
   root@a3df94a31a25:/#
   ```

   

### Stop container :

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
a3df94a31a25        ubuntu              "/bin/bash"         8 minutes ago       Up 7seconds                            gifted_rubin

$ docker stop a3df94a31a25
a3df94a31a25

$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES

$ docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                      PORTS               NAMES
a3df94a31a25        ubuntu              "/bin/bash"         10 minutes ago      Exited (0) 40 seconds ago                       gifted_rubin
```



### Enter container :

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
1836236dc58f        ubuntu              "/bin/bash"         31 seconds ago      Up 8 seconds                            awesome_nobel

$ docker attach 1836236dc58f
root@1836236dc58f:/#
```

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
1836236dc58f        ubuntu              "/bin/bash"         31 seconds ago      Up 8 seconds                            awesome_nobel

$ docker exec -ti awesome_nobel bash
root@1836236dc58f:/#
```



### Remove container 

```
$ docker ps -a
CONTAINER ID        IMAGE               COMMAND              CREATED             STATUS                      PORTS               NAMES
1836236dc58f        ubuntu              "/bin/bash"          13 minutes ago      Exited (0) 45 seconds ago                       awesome_nobel
540b8a01532f        ubuntu              "/bin/bash"          14 minutes ago      Exited (0) 14 minutes ago                       vigorous_tu
7e0e8b7ebcf8        myjupyter           "jupyter notebook"   44 minutes ago      Exited (0) 14 minutes ago                       stoic_meninsky

$ docker rm vigorous_tu
vigorous_tu

$ docker ps -a
CONTAINER ID        IMAGE               COMMAND              CREATED             STATUS                    PORTS               NAMES
1836236dc58f        ubuntu              "/bin/bash"          19 hours ago        Exited (0) 19 hours ago                       awesome_nobel
7e0e8b7ebcf8        myjupyter           "jupyter notebook"   20 hours ago        Exited (0) 19 hours ago                       stoic_meninsky
```





## Repository

### Docker Hub

```
$ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: ericma19920728
Password:
Login Succeeded

$ docker login
Authenticating with existing credentials...
Login Succeeded
```



```
$ docker search ubuntu
NAME                                                      DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
ubuntu                                                    Ubuntu is a Debian-based Linux operating sys…   10482               [OK]
dorowu/ubuntu-desktop-lxde-vnc                            Docker image to provide HTML5 VNC interface …   389                                     [OK]
rastasheep/ubuntu-sshd                                    Dockerized SSH service, built on top of offi…   241                                     [OK]
consol/ubuntu-xfce-vnc                                    Ubuntu container with "headless" VNC session…   209                                     [OK]
ubuntu-upstart                                            Upstart is an event-based replacement for th…   105                 [OK]
ansible/ubuntu14.04-ansible                               Ubuntu 14.04 LTS with ansible                   98                                      [OK]
1and1internet/ubuntu-16-nginx-php-phpmyadmin-mysql-5      ubuntu-16-nginx-php-phpmyadmin-mysql-5          50                                      [OK]
ubuntu-debootstrap                                        debootstrap --variant=minbase --components=m…   42                  [OK]
nuagebec/ubuntu                                           Simple always updated Ubuntu docker images w…   24                                      [OK]
i386/ubuntu                                               Ubuntu is a Debian-based Linux operating sys…   18
1and1internet/ubuntu-16-apache-php-5.6                    ubuntu-16-apache-php-5.6                        14                                      [OK]
1and1internet/ubuntu-16-apache-php-7.0                    ubuntu-16-apache-php-7.0                        13                                      [OK]
ppc64le/ubuntu                                            Ubuntu is a Debian-based Linux operating sys…   13
1and1internet/ubuntu-16-nginx-php-phpmyadmin-mariadb-10   ubuntu-16-nginx-php-phpmyadmin-mariadb-10       11                                      [OK]
1and1internet/ubuntu-16-nginx-php-5.6                     ubuntu-16-nginx-php-5.6                         8                                       [OK]
1and1internet/ubuntu-16-nginx-php-5.6-wordpress-4         ubuntu-16-nginx-php-5.6-wordpress-4             7                                       [OK]
1and1internet/ubuntu-16-apache-php-7.1                    ubuntu-16-apache-php-7.1                        6                                       [OK]
darksheer/ubuntu                                          Base Ubuntu Image -- Updated hourly             5                                       [OK]
1and1internet/ubuntu-16-nginx-php-7.0                     ubuntu-16-nginx-php-7.0                         4                                       [OK]
pivotaldata/ubuntu                                        A quick freshening-up of the base Ubuntu doc…   3
pivotaldata/ubuntu16.04-build                             Ubuntu 16.04 image for GPDB compilation         2
smartentry/ubuntu                                         ubuntu with smartentry                          1                                       [OK]
pivotaldata/ubuntu-gpdb-dev                               Ubuntu images for GPDB development              1
1and1internet/ubuntu-16-php-7.1                           ubuntu-16-php-7.1                               1                                       [OK]
1and1internet/ubuntu-16-sshd                              ubuntu-16-sshd                                  1                                       [OK]
```





## Docker Volume

Docker 資料管理是由 Union File System (UFS) 執行，

Container 關閉後，對比於 image 檔，所有修改過的檔案皆會消失。且 Container 之間互不相連，無法共用資料。

為了保存修改後的檔案，以及讓各容器間能共用資料，

### Docker Volume (on MAC) :

```
$ docker run -it  -v /volume_forder ubuntu /bin/bash
root@928960161bf2:/# 

root@928960161bf2:/# ls -l
total 68
drwxr-xr-x   2 root root 4096 Jan 12 21:10 bin
drwxr-xr-x   2 root root 4096 Apr 24  2018 boot
drwxr-xr-x   5 root root  360 Feb 14 07:49 dev
drwxr-xr-x   1 root root 4096 Feb 14 07:49 etc
drwxr-xr-x   2 root root 4096 Apr 24  2018 home
drwxr-xr-x   8 root root 4096 May 23  2017 lib
drwxr-xr-x   2 root root 4096 Jan 12 21:10 lib64
drwxr-xr-x   2 root root 4096 Jan 12 21:09 media
drwxr-xr-x   2 root root 4096 Jan 12 21:09 mnt
drwxr-xr-x   2 root root 4096 Jan 12 21:09 opt
dr-xr-xr-x 202 root root    0 Feb 14 07:49 proc
drwx------   2 root root 4096 Jan 12 21:10 root
drwxr-xr-x   1 root root 4096 Jan 16 01:20 run
drwxr-xr-x   1 root root 4096 Jan 16 01:20 sbin
drwxr-xr-x   2 root root 4096 Jan 12 21:09 srv
dr-xr-xr-x  13 root root    0 Feb 13 06:31 sys
drwxrwxrwt   2 root root 4096 Jan 12 21:10 tmp
drwxr-xr-x   1 root root 4096 Jan 12 21:09 usr
drwxr-xr-x   1 root root 4096 Jan 12 21:10 var
drwxr-xr-x   2 root root 4096 Feb 14 07:49 volume_folder

root@928960161bf2:/# cd volume_folder
```

Another terminal

```
$ docker inspect 928960161bf2
[
    {...
          "Mounts": [
            {
                "Type": "volume",
                "Name": "9cc52aacc6d5b5b68b936ef18c47ec4ebc655029961f7f09fe8b311ef307ecff",
                "Source": "/var/lib/docker/volumes/9cc52aacc6d5b5b68b936ef18c47ec4ebc655029961f7f09fe8b311ef307ecff/_data",
                "Destination": "/volume_forder",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            }
        ],
     ...}
]

$ screen ~/Library/Containers/com.docker.docker/Data/vms/0/tty
docker-desktop:~#

docker-desktop:~# cd       /var/lib/docker/volumes/9cc52aacc6d5b5b68b936ef18c47ec4ebc655029961f7f09fe8b311ef307ecff/_data
```

```
docker-desktop:/var/lib/docker/volumes/.../_data# touch test-volume

docker-desktop:/var/lib/docker/volumes/.../_data#  ls -l 
total 0
-rw-r--r--    1 root     root             0 Feb 14 08:03 test-volume

root@928960161bf2:/volume_forder# ls -l
total 0
-rw-r--r-- 1 root root 0 Feb 14 08:03 test-volume

root@928960161bf2:/volume_forder# touch test-volume2

root@928960161bf2:/volume_forder# ls -l
total 0
-rw-r--r-- 1 root root 0 Feb 14 08:03 test-volume
-rw-r--r-- 1 root root 0 Feb 14 08:06 test-volume2

docker-desktop:/var/lib/docker/volumes/.../_data#  ls -l 
total 0
-rw-r--r--    1 root     root             0 Feb 14 08:03 test-volume
-rw-r--r--    1 root     root             0 Feb 14 08:06 test-volume2
```

Exit `screen` :`crtl+a` + `crtl+d`

Return `screen` : 

```
$ screen -list
There is a screen on:
	15425.ttys003.Erikde-MacBook-Pro	(Detached)
1 Socket in /var/folders/kg/hlj40p0d3wd837gr38cw57mh0000gn/T/.screen.

$ screen -r 15425
docker-desktop:~#
```

Exit `screen` without return : `crtl+a` + `crtl+k` 

### Local Host Volume :

```
$ docker run -it -v ~/Dokcer/DockerVolume:/Volume_WR -v ~/Doker/DockerVolume:/Volume_R:ro ubuntu /bin/bash
root@793766625ef8:/#

root@793766625ef8:/# ls -l
total 64
drwxr-xr-x   2 root root   64 Feb 14 08:44 Volume_R
drwxr-xr-x   2 root root   64 Feb 14 08:43 Volume_WR
drwxr-xr-x   2 root root 4096 Jan 12 21:10 bin
drwxr-xr-x   2 root root 4096 Apr 24  2018 boot
drwxr-xr-x   5 root root  360 Feb 14 08:44 dev
drwxr-xr-x   1 root root 4096 Feb 14 08:44 etc
drwxr-xr-x   2 root root 4096 Apr 24  2018 home
drwxr-xr-x   8 root root 4096 May 23  2017 lib
drwxr-xr-x   2 root root 4096 Jan 12 21:10 lib64
drwxr-xr-x   2 root root 4096 Jan 12 21:09 media
drwxr-xr-x   2 root root 4096 Jan 12 21:09 mnt
drwxr-xr-x   2 root root 4096 Jan 12 21:09 opt
dr-xr-xr-x 229 root root    0 Feb 14 08:44 proc
drwx------   2 root root 4096 Jan 12 21:10 root
drwxr-xr-x   1 root root 4096 Jan 16 01:20 run
drwxr-xr-x   1 root root 4096 Jan 16 01:20 sbin
drwxr-xr-x   2 root root 4096 Jan 12 21:09 srv
dr-xr-xr-x  13 root root    0 Feb 14 08:44 sys
drwxrwxrwt   2 root root 4096 Jan 12 21:10 tmp
drwxr-xr-x   1 root root 4096 Jan 12 21:09 usr
drwxr-xr-x   1 root root 4096 Jan 12 21:10 var

root@793766625ef8:/# touch Volume_R/test_volume
touch: cannot touch 'Volume_R/test_volume': Read-only file system

$ touch ~/Docker/DockerVolume/test_volume_host_created

$ ls -l ~/Docker/DockerVolume
total 0
-rw-r--r--  1 erikwc_ma  staff  0 Feb 14 16:47 test_volume_host_created

root@626d386cac3f:/# ls -l Volume_R
total 0
-rw-r--r-- 1 root root 0 Feb 14 08:47 test_volume_host_created

root@626d386cac3f:/# ls -l Volume_WR
total 0
-rw-r--r-- 1 root root 0 Feb 14 08:47 test_volume_host_created
```



```
root@626d386cac3f:/# touch Volume_WR/test_volume_container_created

root@626d386cac3f:/# ls -l Volume_R
total 0
-rw-r--r-- 1 root root 0 Feb 17 01:53 test_volume_container_created
-rw-r--r-- 1 root root 0 Feb 14 08:47 test_volume_host_created

$ ls -l ~/Docker/DockerVolume/
total 0
-rw-r--r--  1 erikwc_ma  staff  0 Feb 17 09:53 test_volume_container_created
-rw-r--r--  1 erikwc_ma  staff  0 Feb 14 16:47 test_volume_host_created
```



```
$  docker run -it --name another_container -v ~/Docker/DockerVolume:/Volume_Two ubuntu /bin/bash
root@182c21fce7b7:/#

root@182c21fce7b7:/# ls -l Volume_Two
total 0
-rw-r--r-- 1 root root 0 Feb 17 01:53 test_volume_container_created
-rw-r--r-- 1 root root 0 Feb 14 08:47 test_volume_host_created

root@182c21fce7b7:/# touch Volume_Two/test_volune_container_to_container

root@626d386cac3f:/# ls -l Volume_R
total 0
-rw-r--r-- 1 root root 0 Feb 17 01:53 test_volume_container_created
-rw-r--r-- 1 root root 0 Feb 14 08:47 test_volume_host_created
-rw-r--r-- 1 root root 0 Feb 17 01:58 test_volune_container_to_container
```



```
$ docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                      PORTS               NAMES
182c21fce7b7        ubuntu              "/bin/bash"         23 minutes ago      Exited (0) 8 seconds ago                        another_container
626d386cac3f        ubuntu              "/bin/bash"         2 days ago          Exited (0) 12 seconds ago                       silly_mahavira

$ docker rm -vf $(docker ps -a -q)
182c21fce7b7
626d386cac3f

$ docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES

$ ls -l ~/Docker/DockerVolume/
total 0
-rw-r--r--  1 erikwc_ma  staff  0 Feb 17 09:53 test_volume_container_created
-rw-r--r--  1 erikwc_ma  staff  0 Feb 14 16:47 test_volume_host_created
-rw-r--r--  1 erikwc_ma  staff  0 Feb 17 09:58 test_volune_container_to_container
```





### --volumes-from

```
$ docker run -it -v /opt/dbdata --name dbdata ubuntu /bin/bash
root@a3e9cb4ee2b7:/#

$ docker run -it --name a_container --volumes-from dbdata ubuntu /bin/bash
root@a35ed35617e3:/#

$ docker run -it --name b_container --volumes-from a_container ubuntu /bin/bash
root@f5fc0562e727:/#

root@a3e9cb4ee2b7:/opt/dbdata# touch test_volume_from
root@a3e9cb4ee2b7:/opt/dbdata# ls -l
total 0
-rw-r--r-- 1 root root 0 Feb 18 08:53 test_volume_from

root@f5fc0562e727:/# ls -l opt/dbdata/
total 0
-rw-r--r-- 1 root root 0 Feb 18 08:53 test_volume_from

root@a35ed35617e3:/# ls -l opt/dbdata/
total 0
-rw-r--r-- 1 root root 0 Feb 18 08:53 test_volume_from
```



```
$ docker inspect -f "{{.Mounts}}" a_container
[{volume e9f636536d4c51810e3c3c9aff9a8194e52d4f7e088f704140dc6684a7f0e9f2 /var/lib/docker/volumes/e9f636536d4c51810e3c3c9aff9a8194e52d4f7e088f704140dc6684a7f0e9f2/_data /opt/dbdata local  true }]


$ docker inspect -f "{{.Mounts}}" b_container
[{volume e9f636536d4c51810e3c3c9aff9a8194e52d4f7e088f704140dc6684a7f0e9f2 /var/lib/docker/volumes/e9f636536d4c51810e3c3c9aff9a8194e52d4f7e088f704140dc6684a7f0e9f2/_data /opt/dbdata local  true }]
```

Docker rm -v

```
$ docker rm a_container b_container
a_container 
b_container

docker-desktop:~# ls /var/lib/docker/volumes/
13539b67d0201c431f769404b77e01b6c23e36d7e6d90afd48ecb13b8a75bdb7

$ docker rm -v dbdata
dbdata

docker-desktop:~# ls /var/lib/docker/volumes/

```



## External access Container



```
$ docker run -P jupyter/base-notebook
$ docker ps
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS                     NAMES
dcebc01586ba        jupyter/base-notebook   "tini -g -- start-no…"   5 seconds ago       Up 4 seconds        0.0.0.0:32768->8888/tcp   lucid_albattani
```

Picture



`hostPort:containerPort`

```
$ docker run -p 10000:8888 jupyter/base-notebook
Executing the command: jupyter notebook

$ docker ps
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS                     NAMES
d7e26f386f35        jupyter/base-notebook   "tini -g -- start-no…"   6 minutes ago       Up 6 minutes        0.0.0.0:10000->8888/tcp   vigilant_haslett
```

picture



```
$ docker port 475e667ba4b4
8888/tcp -> 0.0.0.0:32769

$ docker port 475e667ba4b4 8888
0.0.0.0:32769
```

