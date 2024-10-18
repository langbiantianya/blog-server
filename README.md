# blog-server

尽可能不去依赖外部数据库，不依赖外部服务，尽可能简单，不需要并发，技术栈越少越好,公开部分没有api请求，且具有SEO优化。

超自由的自定义!!!

## 替换关键字

1. ${{tag}}当前文章的标签
2. ${{title}}页面标题名字
3. ${{essay}}文章
4. ${{essayPage}}文章分页列表

## 计划

- [ ] 推送通知 例如微信、钉钉、电子邮箱
- [ ] baidu或者goole的访问统计
- [ ] 一个文章管理模块
  - [ ] 文章增删改查
  - [ ] 编译md为html
  - [ ] 分类标签
  - [ ] 图床(自己作为图传、考虑兼容其他图床服务)
  - [ ] ...
- [ ] 静态代理
  - [ ] 代理文章

## 技术选型

### 端口

* 8000 后台api
* 8001 文章静态资源代理

### 数据库

* sqlite

### webServer

- gin

### 配置

- env
- urfave/cli

### 前端框架

- [svelte.js](https://svelte.dev/)
- [fusejs](https://www.fusejs.io/)
- [vditor](https://github.com/Vanessa219/vditor)
- [fluent-ui](https://learn.microsoft.com/zh-cn/fluent-ui/web-components/)
