# beego开发用户管理系统
go语言和beego开发的一个RESTful API接口，一个学习demo

###软件准备工作
1、安装好beego框架，操作流程参考https://beego.me/

``go get github.com/astaxie/beego``

2、安装beego工具，操作流程参考https://beego.me/docs/install/

``go get github.com/astaxie/beego``

3、我是用MySQL数据库，所以需要对应的驱动https://github.com/go-sql-driver/mysql

``go get -u github.com/go-sql-driver/mysql``

4、安装MySQL软件，下载社区版

``https://dev.mysql.com/downloads/file/?id=484920``

5、安装时选择Server Only，如果出现安装失败，注意系统是否已经安装了Visual C++ Redistributable for Visual Studio 2015，8.0的mysql需要该软件支持，下载地址：

``https://www.microsoft.com/zh-cn/download/confirmation.aspx?id=48145``

6、安装数据库查看操作工具，Navicat（收费软件）
输入密码后测试连接，如果出现2059错误，打开MySQL 8.0 Command Line Client这个程序，输入密码，在mysql>后输入

``ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '这是新密码';``

使用新密码再次使用navicat登录mysql即可


###建库建表

1、新建连接名称testDB，输入mysql密码

2、创建表

````
CREATE TABLE `userinfo` (
    `autid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `password` VARCHAR(64) NULL DEFAULT NULL,
    `uid` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`autid`)
);

