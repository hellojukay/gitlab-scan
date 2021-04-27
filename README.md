# gitlab
一个 gitlab 分支扫描工具, 递归扫描 gitlab group 下面的所有项目。

### 安装
```
go get github.com/hellojukay/gitlab
```
### 使用
```
gitlab -api=<https://your gitlab/api/v4> -token=<you token> -group=<group id>
```
![demo](demo.png)
