# 概述

主要是管理集换社的一些功能，比如针对自己的商品，可以

- 批量更新
- 批量上架
- 批量上架

将所有价格在 0-2.99 之间卡牌的价格不增加，以集换价售卖。
go run cmd/jhs_cli/main.go products update -s BTC-03 -r 0,2.99 -u

将所有价格在 3-9.99 之间的非异画卡牌的价格增加 0.5 元。
go run cmd/jhs_cli/main.go products update -s BTC-03 -r 3,9.99 -c 0.5 --art="否"

将所有价格在 0-1000 之间的异画卡牌的价格增加 5 元。
go run cmd/jhs_cli/main.go products update -s BTC-03 -r 10,50 -c 5 --art="是"

将所有价格在 50.01-1000 之间的卡牌的价格增加 10 元。
go run cmd/jhs_cli/main.go products update -s BTC-03 -r 50.01,1000 -c 10

# 构建

go build -o jhs_cli cmd/jhs_cli/main.go
