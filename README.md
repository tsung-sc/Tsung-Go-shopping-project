![Tsung聪慧商城项目](https://img-blog.csdnimg.cn/20201030110321665.png#pic_center)

# **Tsung聪慧商城项目（基于Go语言开发）**
![GitHub](https://img.shields.io/github/license/tsung-sc/Tsung-Go-shopping-project)
</font>

@[TOC](项目目录)


<hr style=" border:solid; width:100px; height:1px;" color=#000000 size=1">

# 前言

<font color=#999AAA >此项目涉及内容：大型企业级项目架构设计、MVC前后端API接口功能分组、用户RBAC权限管理（不同部门用户登陆后台显示不同菜单，设计部门、权限、用户的增删改查以及关联）、轮播图管理（GOLANG动态生成缩略图）、商品分类管理（多级分类关联）、商品管理（商品类型、商品属性、商品图库、商品颜色、商品关联商品、商品关联分类、商品搜索、商品异步ajax排序、商品ajax异步修改数量、商品详情wysiwyg-editor的使用、商品管理中动态生成商品属性表单、批量上传图片）、多协程、会员管理（登录、注册、发送短信、发送语音）、购物车、收货地址管理、订单管理、Golang生成支付二维码、Pc端微信支付、Pc端支付宝支付、事务处理、并发锁、高并发分布式架构、分布式Session、多域名共享Cookie、Redis的使用、Redis发布订阅、Linux部署golang项目、Win部署golang项目、Nginx负载均衡、SSL证书Https配置、前后端分离 RESTful API Api接口设计、Cookie Session跨域、Elasticsearch大数据全文搜索、海量数据查询优化、分布式Oss云存储、阿里云Oss、Jwt+OAuth2.0权限验证、Vue/Angular/react结合Golang实现Jwt权限验证等。</font>

<hr style=" border:solid; width:100px; height:1px;" color=#000000 size=1">

<font color=#999AAA >提示：以下是本篇文章正文内容，下面说明可供参考

# 一、项目规划图


<font color=#999AAA >此项目是由Go语言开发，并使用Beego框架搭建项目结构，以下为此项目的规划路线图，其中大部分功能已经实现，不过仍有些功能正在开发中。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201030105926654.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDQ4MTEyMw==,size_16,color_FFFFFF,t_70#pic_center)


# 二、项目截图
## 1.首页
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201030114401767.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDQ4MTEyMw==,size_16,color_FFFFFF,t_70#pic_center)




## 2.注册界面
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201030114554402.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDQ4MTEyMw==,size_16,color_FFFFFF,t_70#pic_center)

## 3.支付界面
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201030114735122.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDQ4MTEyMw==,size_16,color_FFFFFF,t_70#pic_center)
## 3.后台管理界面
![在这里插入图片描述](https://img-blog.csdnimg.cn/20201030114920441.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDQ4MTEyMw==,size_16,color_FFFFFF,t_70#pic_center)
## 4.更多功能
[探索更多功能请点击这访问聪慧商城官网](https://www.tsung.top)
<hr style=" border:solid; width:100px; height:1px;" color=#000000 size=1">

# 最后
<font color=#999AAA >此项目仍有许多需要完善的地方，如果正在访问的你有意同我一起完善此项目，请联系邮箱：genjutsu2010@gmail.com。


