# atp update with W: Target Packages

```bash
$ sudo apt update              
[sudo] password for moxa: 
Hit:1 http://tw.archive.ubuntu.com/ubuntu focal InRelease
Hit:2 http://tw.archive.ubuntu.com/ubuntu focal-updates InRelease                                                                                       
Hit:3 http://tw.archive.ubuntu.com/ubuntu focal-backports InRelease                                                                                     
Hit:4 http://dl.google.com/linux/chrome/deb stable InRelease                                                                                            
Hit:5 https://download.docker.com/linux/ubuntu focal InRelease                                                                                          
Hit:6 https://packages.microsoft.com/repos/vscode stable InRelease                                                                                     
Get:7 https://typora.io/linux ./ InRelease [793 B]                                                                               
Hit:8 http://security.ubuntu.com/ubuntu focal-security InRelease           
Fetched 793 B in 1s (909 B/s)
Reading package lists... Done
Building dependency tree       
Reading state information... Done
95 packages can be upgraded. Run 'apt list --upgradable' to see them.
W: Target Packages (main/binary-amd64/Packages) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target Packages (main/binary-all/Packages) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target Translations (main/i18n/Translation-en_US) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target Translations (main/i18n/Translation-en) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11 (main/dep11/Components-amd64.yml) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11 (main/dep11/Components-all.yml) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons-small (main/dep11/icons-48x48.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons (main/dep11/icons-64x64.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons-hidpi (main/dep11/icons-64x64@2.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons-large (main/dep11/icons-128x128.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target CNF (main/cnf/Commands-amd64) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target CNF (main/cnf/Commands-all) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target Packages (main/binary-amd64/Packages) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target Packages (main/binary-all/Packages) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target Translations (main/i18n/Translation-en_US) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target Translations (main/i18n/Translation-en) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11 (main/dep11/Components-amd64.yml) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11 (main/dep11/Components-all.yml) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons-small (main/dep11/icons-48x48.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons (main/dep11/icons-64x64.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons-hidpi (main/dep11/icons-64x64@2.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target DEP-11-icons-large (main/dep11/icons-128x128.tar) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target CNF (main/cnf/Commands-amd64) is configured multiple times in /etc/apt/sources.list:60 and /etc/apt/sources.list.d/vscode.list:3
W: Target CNF (main/cnf/Commands-all) is configured multiple times in /etc/apt/sources.list:60 and  /etc/apt/sources.list.d/vscode.list:3
```



/etc/apt/sources.list: (line 60)

```bash
deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main
```

/etc/apt/sources.list.d/vscode.list: (line 3)

```
deb [arch=amd64] http://packages.microsoft.com/repos/vscode stable main
```

choose one to use.

