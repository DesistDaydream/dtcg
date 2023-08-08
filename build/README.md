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

nerdctl build . -t lchdzh/dtcg:v2.8.0 -f build/dtcg/Dockerfile

# 运行

nerdctl run -it --rm --name jhs-exporter --network host -v ~/projects/DesistDaydream/dtcg/internal/database:/dtcg/internal/database lchdzh/jhs-exporter:v1.1.2

nerdctl run -it --rm --name dtcg --network host -v ~/projects/DesistDaydream/dtcg/config_file:/etc/dtcg lchdzh/dtcg:v2.8.0

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

## 2.4.0

1. 添加 /deck/price/wlid/:wlid 接口，可以通过集换社心愿单 ID 获取整个清单的价格。

## 2.5.0

1. 添加 Login 逻辑，添加列出所有用户接口
2. 完善中间件逻辑，允许跨域等
3. config 文件与 NewHandler 参数中关于第三方的信息移到数据库的 user 表中。user 表添加对应的字段
4. 将 dtcg 的 NewHandler 逻辑放到中间件中
5. database 添加列出用户信息逻辑
6. 删除 dtcg api models 中的重复内容，将各种借口所需的 reqQuery 和 reqBody 统一到一个结构体中

## 2.6.0

1. StructToMapStr 功能的 Tag 改为 form，与 Gin 的 ShouldBindQuery 中的转换 map 逻辑中使用的 Tag 保持一致；并且将集换社和 dtcgdb 中的 StructToMapStr 功能提出来合并到一起。
2. 添加用于集换社的列出商品和更新商品两个接口，让前端直接从本程序获取商品信息，避免跨域问题。

2.6.1

- handler 中的 database.GetUser() 逻辑删除，需要在调用 NewHandler 之前从数据库中获取用户信息

2.6.2

- sdk moecard 的 client 添加通用的响应体结构体 CommonResp，同步修改请求逻辑中的最后部分，以处理响应体中的 interface{} 类型的字段。
- 另外，moecard 的 client 中，将 database.GetUser("1") 写死为 1 了，这个需要修改，但是应该从哪里获取用户 ID 是个问题。

## 2.7.0

1. 添加 /api/v1/me 接口用以返回当前登录的用户信息
2. 添加一个 /ping 接口

## 2.7.1

1. 修复 /deck/price/wlid/:wlid 接口 BUG，由于每页最大数据是 15，需要添加分页处理，避免缺少商品

## 2.8.0

1. sdk jhs 添加 card-versions 下的部分接口
2. ProductsListReqQuery.Page 类型改为 int，让 int 到 string 的转换在 sdk 内部实现，避免外部传递异常的字串符