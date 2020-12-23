# 弱口令扫描器

## 概述

使用Go写了一款弱口令扫描器，支持MYSQL/MSSQL/SMB/SSH。使用扫描器可以扫描指定IP段。

##### 扫描模块的输入内容为为IP段，指定端口协议，线程数。

<img src="/Users/xyt/Library/Application Support/typora-user-images/image-20201223093335687.png" alt="image-20201223093335687" style="zoom:50%;" />

- `-h`帮助

- `-d string` 可指定协议对应端口的文件，默认为`port.txt`，可修改。文件内容默认：

  ```html
  SMB:446
  MYSQL:3306
  MSSQL:2245
  SSH:23
  ```

- `-ip string`  ip网段。可以输入单个ip，也可输入ip段。支持任意网段扫描。

- `-p string` 默认四种协议全部扫描，若不全部扫描，可指定。

- `-t int`扫描线程数，默认为10

- `-u string`用户名字典，默认为user.dic 

- `-m string`密码字典，默认为pass.dic

- `-r string`结果保存的字典，默认res.txt

## 运行示例

![image-20201223084937081](/Users/xyt/Library/Application Support/typora-user-images/image-20201223084937081.png)

<img src="/Users/xyt/Library/Application Support/typora-user-images/image-20201223092733102.png" alt="image-20201223092733102" style="zoom:50%;" />