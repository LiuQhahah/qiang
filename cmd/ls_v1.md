
# 功能

列出该文件下的文件信息

https://github.com/coreutils/coreutils/blob/master/src/ls.c

## 需求

- ls 列出基本的文件信息
- ls -l 列出详细的文件,包含大小,所属owner
- ls -s 排序,默认采用文件大小排序

### 需求分析

文件信息包含以下几个方面

- 文件名称
- 文件大小
- 文件所处owner
- 文件日期