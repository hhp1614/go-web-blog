# go-web-blog

本项目是基于beego开发的一个博客系统，采用传统MVC架构开发。

数据库使用MySQL，由beego自带的orm驱动。

## 功能

- [x] 注册
- [x] 登录
- [x] 退出登录
- [x] 写文章
- [x] 显示文章内容
- [x] 更新文章内容
- [x] 删除文章
- [x] 标签页
- [x] 相册页
- [x] 关于我页面

## 文件结构

```text
go-web-blog
│  main.go // 项目入口
│          
├─conf // 项目配置文件
│      
├─controllers // 控制器
│      
├─models // 模型
│      
├─routers // 路由
│      
├─static // 静态资源
│      
├─utils // 工具方法
│      
└─views // 视图
```

## 效果预览

- 注册

![注册](./readme/register.png)

- 登录

![登录](./readme/login.png)

- 首页-未登录

![首页-未登录](./readme/home-unlogin.png)

- 首页-已登录

![首页-已登录](./readme/home-login.png)

- 写文章

![写文章](./readme/add-article.png)

- 显示文章内容

![显示文章内容](./readme/article-detail.png)

- 更新文章内容

![更新文章内容](./readme/edit-article.png)

- 标签页

![标签页](./readme/tags.png)

- 标签页详情

![标签页详情](./readme/tags-detail.png)

- 相册

![相册](./readme/album.png)

- 关于我

![关于我](./readme/about.png)
