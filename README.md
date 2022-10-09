# 概述

获取卡片列表，其中包括卡牌的详情
https://dtcgweb-api.digimoncard.cn/gamecard/gamecardmanager/weblist?
page=1&
limit=40&
name=&
state=0&
cardGroup=&
rareDegree=&
belongsType=&
cardLevel=&
form=&
attribute=&
type=&
color=&
envolutionEffect=&
safeEffect=&
parallCard=&
keyEffect=

获取可用过滤条件的接口。

- https://dtcgweb-api.digimoncard.cn/game/gamecard/weblist # 卡包列表
- https://dtcgweb-api.digimoncard.cn/card/cardraredegree/cachelist # 稀有度
- https://dtcgweb-api.digimoncard.cn/card/cardbelongstype/cachelist # 类别(数码蛋、数码宝贝、驯兽师、选项)
- https://dtcgweb-api.digimoncard.cn/card/cardlevels/cachelist # 等级
- https://dtcgweb-api.digimoncard.cn/card/cardcolor/cachelist # 颜色
- https://dtcgweb-api.digimoncard.cn/card/cardform/cachelist # 形态
- https://dtcgweb-api.digimoncard.cn/card/cardattribute/cachelist # 属性
- https://dtcgweb-api.digimoncard.cn/card/cardtype/cachelist # 类型
- https://dtcgweb-api.digimoncard.cn/card/cardgetway/cachelist # 关键词效果

上面这些接口获取到信息可以填充到 https://dtcgweb-api.digimoncard.cn/gamecard/gamecardmanager/weblist 中的 URL 参数中

# 构建

docker build . -t lchdzh/jihuanshe-exporter -f build/Dockerfile
