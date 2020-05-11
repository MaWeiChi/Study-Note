

# Git

​																																									2020-02

## 分散式版本控制管理系統

Git可以把檔案的狀態作為更新歷史記錄保存起來。因此可以把編輯過的檔案復原到以前的狀態，也可以顯示編輯過內容的差異。
避免版本更新時，相同內容重複儲存，或是共同編輯時遭到覆蓋而遺失。

透過分散式的架構，讓檔案存取機制非限定於單一主機，當遠端數據庫斷線時，使用者可在本機端自行完成版本控制，待恢復連線後再進行同步。

#### 差分編碼

https://zh.wikipedia.org/wiki/%E5%B7%AE%E5%88%86%E7%B7%A8%E7%A2%BC

## 安裝設定（Mac）

### 安裝 Git

1. https://git-scm.com/download/mac

2. brew：

```
$ brew install git
```

Located & Version：

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

>#### Git 中的三個位置：
>
>工作目錄 （Work tree）：正在處理檔案的目錄，Git 相關的操作都會在這個目錄下完成。
>
>暫存區（Staging Area）：又名索引（index），位於工作目錄和數據庫之間，是為了向數據庫提交作準備的暫存區域。
>
>數據庫（Repository）：檔案最後儲存的位置。



----

### init：

```git init```：

```
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
因為這個目錄除了```.git``` 目錄外沒有其他內容，所以回傳 nothing to commit。

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



---

### add：

在修改的過程中，使用 add 將工作目錄裡已經完成修改動作的檔案放置暫存區的指令。

```git add```：
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



##### 衍生用法：

```
$ git add *.html
```

以萬用字元```*```將所有 html 檔加到暫存區。

```
$ git add --all
```

以```--all```參數將所有資料加到暫存區。

```
$ git add .
```

將當下目錄及其所有附屬目錄的檔案加入暫存區，不建議使用。



----

### commit

使用 ```git commit -m "message"``` 指令將已經修改後的檔案從 Staging Area 寫入 Repository 所需執行的動作。

```-m "message" ```說明這次 commit 的摘要讓自己之後回來查詢或是同事可以暸解。 commit 時，系統會要求務必輸入訊息，讓人了解每次執行動作的內容及理由，在訊息空白的狀態下執行 commit 會失敗。


```
$ git commit -m "first commit"
[master (root-commit) b0d97d7] first commit
 1 file changed, 1 insertion(+)
 create mode 100644 welcome.html
 
$ echo "hello ,world" > welcome.html

$ git add welcome

$ git commit -m ""
Aborting commit due to empty commit message.

$ git commit -m " "
Aborting commit due to empty commit message.
```

若執行  commit 時未加上 `-m` 的參數，會進入預設的 Vim 編輯器。

 commit 只會將暫存區中的內容加入數據庫中，尚未加入到暫存區的內容則不會。

每次 commit 後會產生不重複的 hash code（重複機率極低）。


Git 的標準提交訊息：

 第1行：提交時修改內容的摘要
 第2行：空行
 第3行以後：修改的理由



##### 未有任何改變commit ：

加上`--allow-empty`參數。

```
$ git commit --allow-empty -m "nothing changed"
[master 56e571a] nothing changed
```



##### 一鍵 add&commit：

加上 `-a` 參數。

```
$ echo "Hi, git" > welcome.html

$ git commit -a -m "add&commit"
[master 5c9c412] add&commit
 1 file changed, 1 insertion(+), 1 deletion(-)
```

`-a` 參數只針對已存在 Repository 的檔案有效。



##### 查詢 commit 紀錄

`git log` ：

``` 
$ git log
commit 5c9c412fc37de563e64a707cbacdc0bbd99291b3 (HEAD -> master)
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:27:18 2020 +0800

    add&commit

commit 56e571a35cf59f629e92560f1955ca5d757f0f10
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:26:55 2020 +0800

    nothing changed

commit b0d97d78eed6e42f9a506f6baf00a8d2400301f0
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 11:56:15 2020 +0800

    first commit
```

`git log` 提供的資訊：

1. commit 的 hash code。 
2. commit 的編輯者。
3. commit 的時間日期。
4. commit 的訊息。

使用 `git log` 後，按下Ｑ鍵後離開 `git log command` 狀態。


`git log --oneline --graph` ：

```
$git log --oneline --graph
* 5c9c412 (HEAD -> master) add&commit
* 56e571a nothing changed
* b0d97d7 first commit
```

`--oneline`參數呈現每個 commit 的 `hash code` 前七碼及訊息。



`git log --author`：

```
$ git log --author='Ma'
commit 5c9c412fc37de563e64a707cbacdc0bbd99291b3 (HEAD -> master)
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:27:18 2020 +0800

    add&commit

commit 56e571a35cf59f629e92560f1955ca5d757f0f10
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:26:55 2020 +0800

    nothing changed

commit b0d97d78eed6e42f9a506f6baf00a8d2400301f0
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 11:56:15 2020 +0800

    first commit
```

`--author` 參數篩選出特定編輯者。



`git log --grep`：

```
$ git log --grep 'commit'
commit 5c9c412fc37de563e64a707cbacdc0bbd99291b3 (HEAD -> master)
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:27:18 2020 +0800

    add&commit

commit b0d97d78eed6e42f9a506f6baf00a8d2400301f0
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 11:56:15 2020 +0800

    first commit
```

`--grep`參數篩選特定訊息。



`git log -S`：

```
$ echo '1234' > welcome.html

$ git add welcome.html

$ git commit -m 'second commit'
[master f95b055] second commit
 1 file changed, 1 insertion(+), 1 deletion(-)

$ git log -S "1234"
commit f95b055a91a15c689be87968489311d93dfe2040 (HEAD -> master)
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:31:05 2020 +0800

    second commit
```

`-S` 參數在所有 commit 中的檔案篩選出特定內容。



`git log --since --unitl --after --before  `：

 ```
$ git log --oneline --since='10am' --until='12pm'
bb4e0f0 (HEAD -> master) second commit

$ git log --oneline --since='10am' 
bb4e0f0 (HEAD -> master) second commit

$ git log --oneline --until='12pm'
bb4e0f0 (HEAD -> master) second commit
a6fa823 add&commit
5868443 nothing changed
55a48b5 first commit

$ git log --oneline  --until='12pm' --after='2020-02'
bb4e0f0 (HEAD -> master) second commit

$ git log --oneline  --until='12pm' --before='2020-02'
a6fa823 add&commit
5868443 nothing changed
55a48b5 first commit

$ git log --oneline --since='11am' --after='2020-01-30' --before='2020-01-31'
a6fa823 add&commit

$ git log --oneline --since='11am' --after='2020-01-30'
bb4e0f0 (HEAD -> master) second commit
a6fa823 add&commit
 ```

使用 `--since` ` --unitl` ` --after` `--before` 參數對特定時間點的 commit 篩選，`--since`及`--until` 指定時間區段， `--after`及`--before` 指定日期區段，若未使用 `--before` 參數指定結束日期，則預設為至今日。



`git log filename`

```
$ git log welcome.html
commit ddc1eac5c4c4213c9433188078ec8defc2d86417 (HEAD -> master)
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:31:05 2020 +0800

    second commit
```

`git log -p filename`

```
$ git log -p welcome
commit ddc1eac5c4c4213c9433188078ec8defc2d86417 (HEAD -> master)
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:31:05 2020 +0800

    second commit

diff --git a/welcome.html b/welcome.html
index 707b497..81c545e 100644
--- a/welcome.html
+++ b/welcome.html
@@ -1 +1 @@
-Hi, git
+1234

commit 5c9c412fc37de563e64a707cbacdc0bbd99291b3
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 13:27:18 2020 +0800

    add&commit

diff --git a/welcome.html b/welcome.html
index 94bab17..707b497 100644
--- a/welcome.html
+++ b/welcome.html
@@ -1 +1 @@
-hello, git
+Hi, git

commit b0d97d78eed6e42f9a506f6baf00a8d2400301f0
Author: Erik Ma <ErikWC.Ma@moxa.com>
Date:   Thu Feb 6 11:56:15 2020 +0800

    first commit

diff --git a/welcome.html b/welcome.html
new file mode 100644
index 0000000..94bab17
--- /dev/null
+++ b/welcome.html
@@ -0,0 +1 @@
+hello, git
```



`git blame filename`

```
$git blame welcome
ddc1eac5 (Erik Ma 2020-02-06 13:31:05 +0800 1) 1234
```



##### 修改 commit 訊息 & 追加檔案

(不要使用在已經 push 的 commit)

`git commit --amend -m`：

修改訊息：

```
$ git commit --amend -m '2nd commit' 
[master 0b24974] 2nd commit
 Date: Thu Feb 6 13:31:05 2020 +0800
 1 file changed, 1 insertion(+), 1 deletion(-)
 
$ git log --oneline
0b24974 (HEAD -> master) 2nd commit
5c9c412 add&commit
56e571a nothing changed
b0d97d7 first commit
```

`--amend` 參數可以修改上一次 commit 的訊息。

因為 commit 內容修改後，Git 會重新計算新的 hash code。

追加檔案：

```
$ echo 'bye bye' > goodbye.html

$ git add goodbye.html

$ git commit --amend --no-edit
[master e316235] 2nd commit
 Date: Thu Feb 6 13:31:05 2020 +0800
 2 files changed, 2 insertions(+), 1 deletion(-)
 create mode 100644 goodbye.html

$ git log --oneline
e316235 (HEAD -> master) 2nd commit
5c9c412 add&commit
56e571a nothing changed
b0d97d7 first commit
```

`--no-edit` 參數為不編輯訊息或是可用。
修改訊息 ＆ 追加檔案

```
$ echo 'see ya' > goodbye.html

$ git add goodbye.html

$ git commit --amend -m 'second commit'
[master ddc1eac] second commit
 Date: Thu Feb 6 13:31:05 2020 +0800
 2 files changed, 2 insertions(+), 1 deletion(-)
 create mode 100644 goodbye.html
 
$ git log --oneline
ddc1eac (HEAD -> master) second commit
5c9c412 add&commit
56e571a nothing changed
b0d97d7 first commit
```













---

### rm & mv

在 Git 中刪除檔案和修改檔名都是一種「修改」。



##### 使用系統指令刪除檔案或更改檔名

 `rm` 刪除檔案：

```
$ rm welcome.html 

$ git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    welcome.html

no changes added to commit (use "git add" and/or "git commit -a")
```

此時 `welcome.html` 檔案狀態為 deleted ，若需將完成這次刪除動作，還需使用 `git add` 將這次修改加到暫存區。

```
$ git add welcome.html

$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	deleted:    welcome.html
```



`$ git checkout filename`

```
$ rm welcome.html

$ ls -l
total 8
-rw-r--r--  1 erikwc_ma  staff  7 Feb  6 13:40 goodbye.html

$ git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    welcome.html

no changes added to commit (use "git add" and/or "git commit -a")

$ git checkout welcome.html
Updated 1 path from the index

$ ls -l
total 16
-rw-r--r--  1 erikwc_ma  staff  7 Feb  6 

$ git status
On branch master
nothing to commit, working tree clean:40 goodbye.html
-rw-r--r--  1 erikwc_ma  staff  5 Feb  7 14:00 welcome.html
```



`$ git checkout *`

```
$ rm *

$ ls -l

$ git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    goodbye.html
	deleted:    welcome.html

no changes added to commit (use "git add" and/or "git commit -a")

$ git checkout *
Updated 2 paths from the index

$ ls -l
total 16
-rw-r--r--  1 erikwc_ma  staff  7 Feb  7 14:30 goodbye.html
-rw-r--r--  1 erikwc_ma  staff  5 Feb  7 14:30 welcome.html

$ git status
On branch master
nothing to commit, working tree clean
```





 `mv` 更改檔名：

```
$ mv welcome.html hello.html

$ git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "restore <file>..." to discard changes in working directory)
	deleted:    welcome.html

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	hello.html

no changes added to commit (use "git add" and/or "git commit -a")
```

與刪除檔案一樣，需使用 `git add` 將修改加入暫存區，但需對新舊檔名各做一次。

```
$ git add hello.html

$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	new file:   hello.html

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    welcome.html

$ git add welcome.html

$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	renamed:    welcome.html -> hello.html
```



##### 使用 Git 指令刪除檔案或更改檔名

`git rm`：

```
$ git rm welcome.html
rm 'welcome.html'

$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	deleted:    welcome.html
```

`git rm` 指令達到 `rm` 及 `git add`兩個指令的效果。



`git rm --cached`：

```
$ git rm --cached welcome.html
rm 'welcome.html'

$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	deleted:    welcome.html

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	welcome.html

$ ls -l
total 8
-rw-r--r--  1 erikwc_ma  wheel  5 Feb  3 13:17 welcome.html
```

`--cached` 參數讓檔案脫離 git 的控管，並未刪除檔案。



`git mv`：

```
$ git mv welcome.html hello.html

$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	renamed:    welcome.html -> hello.html

```



---



### .keep

```
$ mkdir afolder

$ git status
On branch first_branch
nothing to commit, working tree clean

$ touch afolder/.keep

$ git status
On branch first_branch
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	afolder/

nothing added to commit but untracked files present (use "git add" to track)
```



---

### .gitignore









































