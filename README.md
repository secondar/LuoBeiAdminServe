<p align="center">
  <img width="320" src="https://www.bugquit.com/bg/ad/logo.png">
</p>

## 由来

- 一直想做一款后台管理系统，看了很多优秀的开源项目但是发现没有合适自己的。
- 因为自己不是很喜欢用已经完完整整的较为臃肿的后台管理系统，仅仅就想要一个做好权限管理就好了的后台管理系统
- 于是利用空闲休息时间开始自己写一套后台系统。 如此有了萝北管理系统。
- 它可以用于所有的 Web 应用程序，如网站管理后台，网站会员中心，CMS，CRM，OA 等等
- 当然，您也可以对它进行深度定制，以做出更强系统。
- 所有前端后台代码封装过后十分精简易上手，出错概率低。
- 同时支持移动客户端访问。系统会陆续更新一些实用功能。

前端地址：

[gitee](https://gitee.com/secondar/luo-bei-admin-view.git)

[github](https://github.com/secondar/LuoBeiAdminView.git)

## 现状

- 用户管理 > 系统账户管理
- 角色管理 > 控制角色内的账户所拥有的权限
- 菜单管理 > 控制前端页面路由及提供角色管理进行权限控制
- 系统配置 > 仅有配置系统的名称等，没有啥具体用处
- 文章管理 > 普通文章 CURD

 <img width="100%" src="https://www.bugquit.com/bg/ad/1.png">
 <img width="100%" src="https://www.bugquit.com/bg/ad/2.png">
 
## 准备

- golang
- bee
- beego
- git

## 开始

```bash
git clone https://gitee.com/secondar/luo-bei-admin-serve.git
# OR
git clone https://github.com/secondar/LuoBeiAdminServe.git

cd luo-bei-admin-serve
# OR
cd LuoBeiAdminServe

# 新建数据库导入sql下的sql文件，并在conf/app.conf配置相关设置

bee run

```
