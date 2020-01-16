

# Git
## 分散式版本控制管理系統
Git可以把檔案的狀態作為更新歷史記錄保存起來。因此可以把編輯過的檔案復原到以前的狀態，也可以顯示編輯過內容的差異。
避免版本更新時，相同內容重複儲存，或是共同編輯時遭到覆蓋而遺失。

透過分散式的架構，讓檔案存取機制非限定於單一主機，當遠端數據庫斷線時，使用者可在本機端自行完成版本控制，待恢復連線後再進行同步。
#### 差分編碼
https://zh.wikipedia.org/wiki/%E5%B7%AE%E5%88%86%E7%B7%A8%E7%A2%BC

## 安裝設定（Mac）
### 安裝 Git
1. 網站：https://git-scm.com/download/mac

2. brew：
```
$ brew install git
```

檢查位置及版本：
```
$ which git
/usr/local/bin/git
$ git --version
git version 2.25.0
```
### 使用者設定
使用前，先設定使用者名稱及信箱：
```
$ git config --global user.name "EricMa"
$ git config --global user.email "Ericma19920728@gmail.com"
```
檢視目前設定：
```
$ git config --list 
user.name=EricMa 
user.email=Ericma19920728@gmail.com
```
設定區域使用者：
```
$ git config --local user.name "EriK"
$ git config --local user.email "EriK@gmail.com"
```
### 其他設定
建立縮寫：
```
$ git config --global alias.co checkout 
$ git config --global alias.br branch
$ git config --global alias.st status
$ git config --global alias.l "log --oneline --graph"    #log查詢
$ git config --global alias.ls 'log --graph --pretty=format:"%h <%an> %ar %s"'    #log、使用者、時間查詢
```

### Repository (數據庫) 
Repository 是 Git 存放檔案、目錄及修改記錄的地方。可以追蹤狀態和版本。
分為本機和遠端數據庫。遠端數據庫可多人共享資料。

>#### Git 使用中會有三個位置：
>
>工作目錄 （Work tree）：正在處理檔案的目錄，Git 相關的操作都會在這個目錄下完成。
>
>暫存區（Staging Area）：又名索引（index），位於工作目錄和數據庫之間，是為了向數據庫提交作準備的暫存區域。
>
>數據庫（Repository）：檔案最後儲存的位置。

### init：
```git init```：
```
$ cd /tmp    
$ mkdir git-practice    
$ cd git-practice    
$ git init    # 初始化這個目錄，讓 Git 對這個目錄開始進行版控
Initialized empty Git repository in /private/tmp/git-practice/.git/
```
使用 ```git init``` 指令初始化這個目錄，讓 Git 開始對這個目錄進行版本控制。

這個指令會在這個目錄裡建立了一個 ```.git``` 檔，整個 Git 最重要的部分在這個 ```.git``` 檔裡。

```git status```：
```
$ git status
On branch master
Initial commit
nothing to commit (create/copy files and use "git add" to track)
```
使用```git status``` 查看這個目錄的狀態。
因為這個目錄除了```.git``` 目錄外沒有其他內容，所以回傳 ```nothing to commit```。

```
$echo "hello, git" > welcome.html

$ git status
On branch master
Initial commit
Untracked files:
(use "git add <file>..." to include in what will be committed)
welcome.html
nothing added to commit but untracked files present (use "git add" to track)
```

 ```welcome.html``` 目前狀態是 Untracked files ，意思是檔案尚未被加到 Git 版控系統裡，還沒開始正式被 Git「追蹤」，它只是剛剛才加入這個目錄而已。
目錄在初始化後，初始化前已經存在的檔案亦是 Untracked files 。

### add：
在修改的過程中，使用 add 將工作目錄裡已經完成修改動作的檔案放置暫存區的指令。

```git add```
使用```git add```指令讓檔案納入 Git 管控。

``` 
$ git add welcome.html

$ git status 
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
	new file:   welcome.html
```

檔案由 Untracked files 變成 new file ，表示檔案從工作目錄安置到暫存區。

##### ```git add```後，又修改檔案：

````
$ echo "world" >  welcome.html

$ git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
	new file:   welcome.html

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   welcome.html
````
此時 new file 的```welcome.html```為原本內容為 "hello, git" 的檔案，而修改過（modified）的 ```welcome.html``` 則尚未加入暫存區，需再次使用```git add```。 


##### Git 回收機制：



##### 衍生用法

```
$ git add *.html
```
以萬用字元```*```將所有 html 檔加到暫存區。

```
$ git add --all
```
以```--alll```參數將所有資料加到暫存區。

### commit
使用```git commit -m "message"```指令將已經修改後的檔案從 Staging Area 寫入 Repository 所需執行的動作。

```-m "message"```說明這次```commit```的摘要讓自己之後回來查詢或是同事可以暸解。```commit```時，系統會要求務必輸入訊息，讓人了解每次執行動作的內容及理由，在訊息空白的狀態下執行```commit```會失敗。


```
$ git commit -m "init commit"
[master (root-commit) ac36770] init commit
 1 file changed, 1 insertion(+)
 create mode 100644 welcome.html
 
$ echo "hello ,world" > welcome.html
$ git add welcome
$ git commit -m ""
Aborting commit due to empty commit message.
$ git commit -m " "
Aborting commit due to empty commit message.
```






```commit```只會將暫存區中的內容加入數據庫中，尚未加入到暫存區的內容則不會。


每次 commit 後會產生不重複的 hash code（重複機率極低，比中樂透還低）
 

Git的標準提交訊息：

 第1行：提交時修改內容的摘要
 第2行：空行
 第3行以後：修改的理由
 
 


