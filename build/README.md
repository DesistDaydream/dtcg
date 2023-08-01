# 概述

使用 gin 实现，还想自己写个 web 页面，实现很多功能

已实现的：

- 通过卡组信息获取该卡组的价格
- 获取卡片描述

还想实现很多功能：

- 获取集换社信息并在 web 展示
- 自己写一个销售系统，可以在 web 端操作
- 等等

# 构建

nerdctl build . -t lchdzh/jhs-exporter:v1.1.2 -f build/jhs_exporter/Dockerfile

nerdctl build . -t lchdzh/dtcg:v2.3.2 -f build/dtcg/Dockerfile

# 运行

nerdctl run -it --rm --name jhs-exporter --network host -v ~/projects/DesistDaydream/dtcg/internal/database:/dtcg/internal/database lchdzh/jhs-exporter:v1.1.2

nerdctl run -it --rm --name dtcg --network host -v ~/projects/DesistDaydream/dtcg/config_file:/etc/dtcg lchdzh/dtcg:v2.3.2

# ChangeLog

## 1.3.0

添加数个逻辑

1. 根据 DTCG DB 卡组广场中卡组的 URL 中最后的 HID，获取纯字符串格式的卡组所有卡牌的 ID
2. 根据 DTCG DB 卡组广场中卡组的 URL 中最后的 HID，直接获取卡组价格

## 1.4.0

1. 添加 /card/price 接口，对应的数据库操作添加获取卡牌价格带分页逻辑。

## 1.5.0

1. 添加 /set/desc 接口，对应的数据库操作添加获取卡牌集合描述带分页逻辑。

## 1.6.0

1. 添加 /deck/price/cdid/:cdid 接口，通过云卡组 ID 获取卡组价格

## 1.7.0

1. 添加 /card/price GET 接口
2. 将 /card/price POST 接口改为可以通过条件筛选结果的逻辑

## 1.8.0

1. 添加 /card/pricewithimg 接口，用以获取带有 dtcgdb 网站中图片的卡牌价格详情。同步添加相关逻辑和数据库逻辑。
2. 添加了一些注释
3. 修改了一些结构体名字

## 2.0.0

集换社和卡查的 TOKEN 放到数据库中保存

## 2.1.0

1. 添加 /user/info/:userid 接口，以获取存到数据库中的用户信息，包括各种 TOKEN。

## 2.2.0

1. 添加周期性更新数据库中集换社的 Token 的逻辑，以避免过期。

## 2.3.0

1. 添加了认证中间件、登录逻辑、登录验证逻辑。
2. 配置文件中添加两个字段: tokenExpiresAt 和 jsh.autoUpdateTokenDuration
3. 将全局 Flags 相关逻辑放到单独的目录中

TODO: 需要将数据库中的密码加密
