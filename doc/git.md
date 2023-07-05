#### 免密码操作

```
git config --global user.name "用户名"  
git config --global user.email "邮箱"  
git config --global credential.helper store  
ssh-keygen -t rsa -C "git的邮箱"  
cat ~/.ssh/id_rsa.pub  

```

#### 撤销add

    git reset HEAD

#### 撤销刚刚的commit

    git reset --soft HEAD^

#### 拉取本地没有的远程分支

    git checkout -b 本地分支名 远程分支全称

#### 远程库同步到本地库

    git pull --rebase origin master

#### 推送所有tag

    git push --tags

#### 切换到master分支上，合并修改的分支到master

    git checkout master  
    git merge bugfix01  

#### GIT使用过程中出现，回退  (master|REBASE 1/10)

    git rebase --abort  

#### 查看提交次数

    git log | grep "^Author: " | awk '{print $2}' | sort | uniq -c | sort -k1,1nr

#### 多账户配置

*   <https://www.cnblogs.com/popfisher/p/5731232.html>

#### 只拉取某个文件

    git init projectName
    cd projectName
    git remote add origin 克隆地址
    git config core.sparsecheckout true
    echo 文件夹/文件 >> .git/info/sparse-checkout
    git pull --depth=1 origin master

#### git 清空所有commit记录方法

```
git checkout --orphan latest_branch
git add -A
git commit -am "commit message"
git branch -D master
git branch -m master
git push -f origin master

```

#### 强制提交

    git push -f origin master  

#### 硬核当前分支变为master

    git checkout dev
    git branch -D master
    git checkout -b master
    git push -u origin master --force

#### 强制覆盖本地

    git fetch --all  
    git reset --hard origin/master  

#### git删除远程分支

    git push origin --delete [branch_name]

#### 浅克隆

    git init <repo>
    cd <repo>
    git remote add origin <url>
    git config core.sparsecheckout true
    echo "finisht/*" >> .git/info/sparse-checkout
    git pull --depth=1 origin master

### 基础命令

    git add <文件名>	将工作区的内容放置暂存区
        git add * 		表示将工作区全部添加进暂存区

    git commit -m "注释内容"		将暂存区内文件存放到git仓库

    git commit --amend 		可以修改提交注释内容

    git log 	表示查看当前版本的历史版本快照
        （1）--decorate 		显示指向提交的所有引用（如分支、标签）
        （2）--oneline 		精简化显示一个快照的格式
        （3）--graph 		图形化方式显示
        （4）--all 			显示所有分支

    git reflog		可以看到版本前后所有快照

    git rm <文件名> 		删除工作区和暂存区文件
    	（1）需要把git区文件也删除，只需用git reset --soft HEAD~[num]/快照ID 改变头指针
    	
    	（2）当暂存区与工作区文件名相同，内容不同时，需要强制删除工作区和暂存区：git rm -f <文件名>
    	
    	（3）当暂存区与工作区文件名相同，内容不同时，需要删除暂存区，保留工作区：git rm --cached <文件名>
    	
    git mv <原文件名> <修改后文件名>		git重命名文件
    	
    	
    git branch <分支名> 		创建分支
    -a 						显示全部分支
    -d <分支名>				删除分支
    --delete <分支全名>		删除分支
    	
    git merge <分支名> 			将<分支名>与当前合并分支

### 版本操作

1.  git reset
    返回快照到暂存区。
    (1)git reset \[--mixed] HEAD\~\[num]

        (2)git reset [--soft] HEAD~[num]

        (3)git reset [--hard] HEAD~[num]

           --mixed表示将HEAD指向某个快照,并将暂存区内容更新为所指向的快照内容。
           --soft表示撤销上一次的提交git仓库操作，暂存区内容不变
           --hard表示将HEAD指向某个快照,并把指向暂存区的快照还原到工作区
               该参数位置为空时，默认--mixed
           ~表示前一个
           num表示~的数量，默认为1。

        (4)git reset [--mixed/--soft/--hard] <id值>		表示回滚指定快照

        (5)git reset 版本快照 文件名/路径		表示回滚个别文件

2.  git checkout <文件名>
    ```
    从暂存区恢复<文件名>文件到工作目录
    （1）git checkout HEAD~ <文件名>		<文件名>把git区上一个快照返回至工作区、暂存区

    （2）git checkout -- <文件名> 			从暂存区恢复<文件名>文件到工作目录，--预防恰好有一个分支叫做<文件名>

    （3）git checkout <分支名/id值>			切换分支，此时HEAD从原来 HEAD -> master 变为 HEAD -> <分支名> -> <快照>，并将快照返回到暂存区和工作区

    （4）git checkout -b <分支名> 			创建分支并切换到分支上
       
    ```

3.  git diff
    ```
    比较显示新（暂存区）旧（工作区）文件的区别
    操作：

        键盘按键 J 向下移动一行
        键盘按键 K 向上移动一行
        键盘按键 D 向下移动半页
        键盘按键 U 向上移动半页
        键盘按键 F 向下移动一页
        键盘按键 B 向上移动一页
        键盘按键 g 跳转第一行
            输入 3g 跳转第三行
        键盘按键 G 跳转最后一行
        键盘按键 /<内容> 从上向下搜索全部匹配内容
            按键盘 n 查找下一个
            按键盘 N 查找下一个
            
        键盘按键 Q 退出操作diff
        
    比较两个历史快照：
        git diff 快照ID1 快照ID2
        
    比较当前工作区与git区：
        git diff 快照ID1
        
    比较暂存区与git区：
        git diff --cached [快照ID]  不填ID，则比较的是git区最新快照与暂存区。填写则为指定的git区与暂存区进行比较

    ```

### 快照

    ```
    恢复文件：
        reset 比 checkout 更安全
        git reset --mixed <文件名> 		恢复<文件名>到暂存区
        git checkout <文件名>			恢复<文件名>到暂存区和工作区
        
    恢复快照：
        （1）checkout 比 reset 更安全
                git checkout <分支名>			在切换分支前会检查当前工作状态，如果不是clean，则无法操作
                git reset --hard HEAD~[num]		直接覆盖当前暂存区和工作区
                
        （2）更新HEAD指向
                reset 		移动HEAD所在分支指向			HEAD -> master -> 
                checkout 	移动HEAD自身指向另一个分支		HEAD -> 
    ```

### 更换git

    vi /Users/lcc/IdeaProjects/source_code/netty/.git/conf

    [core]
        repositoryformatversion = 0
        filemode = true
        bare = false
        logallrefupdates = true
        ignorecase = true
        precomposeunicode = true
    [remote "origin"]
        url = https://github.com/lccbiluox2/netty.git  # 这里修改为新的服务器
        fetch = +refs/heads/*:refs/remotes/origin/*
    [branch "4.1"]
        remote = origin
        merge = refs/heads/4.1

#### git代理

*   git config --global http.proxy <http://127.0.0.1:1080>
*   git config --global https.proxy <https://127.0.0.1:1080>

