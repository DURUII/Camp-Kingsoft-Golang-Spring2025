# Bootcamp-Kingsoft-Server-Spring2025

参与这门训练营配合互联网资源自学，用到的学习记录和作业实现汇总在此仓库。

- `./basics` 包含 Go 语言的基本语法样例代码，配合详细的注释并与其他主流编程语言进行对比（如
  Python/Java），对避免陷入其他语言思考模式的固有陷阱有很大的帮助，包括：
    - ⚙️ 基础结构：常量、变量、条件、循环、字符串、slice、map
    - 🌆 抽象：结构体、接口
    - 🤕 错误处理与测试：单元测试、benchmark、BDD、json 性能调优
    - 🏗️ 架构设计：微内核、管道过滤器
    - ⚡️ 并发编程：mutex、waitgroup、channel 等
- `./leetcode` 包含用力扣练习巩固 Golang 基础语法的代码。
- `./proj` 包含我对训练营入门作业的设计文档和实现代码，考核作业基础需求不高因此没有skeleton_code，主要考察基础，有庞大的自我发挥空间。

## 作业简介

1. 🎮 简易数字猜谜游戏 [[README]](./proj/numguess/README.md)

- 基于 Golang 和 Bubbletea 实现的现代终端 UI 猜数字游戏，采用 Elm 架构（MVU 模型）实现交互逻辑与视图分离，明确状态机跳转，视图更新由数据驱动。
- 支持实时输入合法性校验、交互式反馈等功能，具备状态管理与用户操作追踪。

<div align="center">
<img src="./proj/numguess/assets/demo.gif" style="width:50%;"/>
</div>

2. 🧩 简易 JSON 词汇数据解析并写入 SQLite

- 设计并实现一个高性能、可扩展的数据导入方案，将结构化 JSON 词汇数据解析后存入
  SQLite，数据来源于开源项目[english-vocabulary](https://github.com/KyleBing/english-vocabulary/)。
- 技术亮点：
    - **数据表建模**
      ：将词汇、翻译、短语抽象为独立表，参考[阿里云数据设计规范](https://developer.aliyun.com/article/709387)，支持软删除与多来源管理。
    - **写入性能调优**：采用批量插入、事务控制、WAL模式、关闭 GORM 日志与默认事务，SQLite 写入性能提升数量级。优化后导入
      CET4 + CET6 总数据耗时降至 **500ms**（原始实现单个文件 30s），性能提升数量级。

## 优质在线资源

- The Uber Go Style Guide：https://github.com/uber-go/guide/blob/master/style.md
- 深入理解Go并发编程：https://github.com/smallnest/concurrency-programming-via-go-code
- Go语言从入门到实战：https://gitee.com/geektime-geekbang/go_learning