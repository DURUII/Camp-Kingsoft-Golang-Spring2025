## 玩具项目

1. 🎮 简易数字猜谜游戏 [[README]](./proj/numguess/README.md)

- 基于 Golang 和 Bubbletea 实现的现代终端 UI 猜数字游戏，采用 Elm 架构（MVU 模型）实现交互逻辑与视图分离，明确状态机跳转，视图更新由数据驱动。
- 支持实时输入合法性校验、交互式反馈等功能，具备状态管理与用户操作追踪。

<div align="center">
<img src="./proj/numguess/assets/demo.gif" style="width:50%;"/>
</div>

2. 🧩 简易 JSON 词汇数据解析并写入 SQLite [[README]](./proj/fileprocessing/README.md)

- 设计并实现一个高性能、可扩展的数据导入方案，将结构化 JSON 词汇数据解析后存入
  SQLite，数据来源于开源项目[english-vocabulary](https://github.com/KyleBing/english-vocabulary/)。
- 技术亮点：
    - **数据表建模**
      ：将词汇、翻译、短语抽象为独立表，参考[阿里云数据设计规范](https://developer.aliyun.com/article/709387)，支持软删除与多来源管理。
    - **写入性能调优**：采用批量插入、事务控制、WAL模式、关闭 GORM 日志与默认事务，SQLite 写入性能提升数量级。优化后导入
      CET4 + CET6 总数据耗时降至 **500ms**（原始实现单个文件 30s），性能提升数量级。

<!-- 3. 🤖 基于 LLM 的编程考核题生成服务 [[README]](./proj/ai-question-generator/README.md)

- 设计并实现了一个基于 Golang 的编程题生成服务，支持动态生成多种类型的编程题：支持灵活但可控的的参数配置、支持多模型切换（如通义、deepseek）以及清晰的格式输出（JSON）。
- 技术亮点
    - 采用分层架构（接入层、服务层、适配层、数据操作层）提高可维护性与扩展性；实现了通用接口抽象，支持多种第三方服务接入，增强了代码复用性。
    - 通过迭代优化提示词，并针对提示词注入问题（如下图）做输入校验，缓解了大语言模型返回内容不可控问题。

<div align="center">
<img src="./proj/ai-question-generator/doc/img.png" style="width:50%;"/>
</div> -->