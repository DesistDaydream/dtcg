# 概述

> [!Warning]
> 卒于 2023 年 11 月27 日，集换社 app 修改了加密逻辑，且 app 使用了爱加密企业版套壳，小程序的调用变成了云函数。个人能力有限，暂时无法继续更新集换社数据了。

dcgo # DTCG 电子版相关处理程序

download_images 下载 DTCG 图片。可选从 日版、美版、中版网站。

dtcg # tcg 前端程序中关于 dtcg 的后端处理程序

dtcg_cli # dtcg 相关命令行工具。向数据库添加卡牌信息、价格信息、卡集信息、etc.

jhs_cli # 集换社相关命令行工具。主要是与商家上相关的，商品管理（批量上架下架）、愿望单管理、订单管理、etc.

jhs_exporter # 用于监控卡牌价格的集换社 exporter

statisitcs # 简单的统计相关工具。统计每个卡盒中各种稀有度的数量以及异画数量、统计每个级别的各种 DP 的数量

# internal

database # 内部数据库的处理程序。

# pkg

database # 弃用，老的数据库处理

sdk # 集换社、bandai tcg plus、dtcg 卡查 的 SDK

