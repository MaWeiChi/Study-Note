# Docker Container


### Concept

Docker image : 可單獨運行的輕量級套件，包含所有執行程式所需要的函式庫、環境變數與設定檔等

Docker container : image 的執行實體，將影像（image）載入至記憶體中執行之後的環境。


VM 相較於 Docker 包含程式及函式庫外，還需使用作業系統（OS）。

VM架構
![image](https://github.com/EricMa19920728/study-note/blob/master/Picture/virtual-machine-20170625-1.png)

Docker架構
![image](https://github.com/EricMa19920728/study-note/blob/master/Picture/docker-container-20170625-1.png)

Docker container 只需要包含程式及函式庫，所有程式共用一個 Host OS 系統核心
