[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

# 安装
1 点此下载最新版本：https://github.com/Huweicai/goto/releases 

2 打开—安装

# 使用
### 添加 
输入（or 自定义快捷键） ：gadd example  http://example.com/example

路径为 example 的这样一个链接就添加成功了。支持多级目录，用空格分隔，格式如：
dir1 dir2 key url，也可以用来存一个值，当备忘录使用，当你存入的值不是一个URL时会将其复制到剪贴板中
 
 ![add](https://anonymous-1253692322.cos.ap-beijing.myqcloud.com/github/goto/goto_add_example.png)
 
 
### 打开
输入（or 自定义快捷键）：goto excample

然后就会根据这个路径在浏览器中打开对应的网址 http://example.com/example

输入的时候是支持自动补全的，可以在极快的时间内打开指定的地址

![open](https://anonymous-1253692322.cos.ap-beijing.myqcloud.com/github/goto/goto_open_example.png)


### 实现原理
所以的信息都是储存在 workFlow 目录下的 config.yaml文件中的，@表示它的上一级对应的打开路径，所以其实可以通过直接修改这个yaml
文件来进行对URL的管理

![config.yaml]( https://anonymous-1253692322.cos.ap-beijing.myqcloud.com/github/goto/goto_data.png)
