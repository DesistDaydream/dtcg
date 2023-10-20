--- 删除记录后将 ID 重新排序，以便后面的添加可以连贯
SET @i=0;
UPDATE `card_descs` SET `id`=(@i:=@i+1);
ALTER TABLE `card_descs` AUTO_INCREMENT=0

--- 获取所有原画卡牌价格的总和
SELECT
    SUM(min_price) AS all_min_price,
    SUM(avg_price) AS all_avg_price
FROM
    card_prices
WHERE alternative_art = '否'

--- 获取指定卡包中的所有卡牌的价格的总和
SELECT
    ROUND(SUM(min_price),2) AS all_min_price,
    ROUND(SUM(avg_price),2) AS all_avg_price
FROM
    card_prices
WHERE set_prefix = 'SPC-1A';

--- 获取多个卡包中的所有卡牌的价格的总和
SELECT
    set_prefix,
    ROUND(SUM(min_price),2) AS all_min_price,
    ROUND(SUM(avg_price),2) AS all_avg_price
FROM
    card_prices
WHERE
    set_prefix IN ('BTC-04','BTC-05', 'BTC-06', 'BTC-07')
GROUP BY
    set_prefix;

--- 查询价格低于 500 的所有卡牌价格的总和
SELECT
    SUM(if(min_price < 500, min_price, 0)) AS all_min_price,
    SUM(if(avg_price < 500, avg_price, 0))
	  AS all_avg_price
FROM
    card_prices

--- 查询带有 card_desc 表中的 image 字段的 card_prices 表中的数据
SELECT
    card_prices.card_id_from_db,
    card_prices.set_prefix,
    card_prices.sc_name,
    card_prices.alternative_art,
    card_prices.rarity,
    card_prices.min_price,
    card_prices.avg_price,
    card_descs.image
FROM
    card_prices
    INNER JOIN card_descs ON card_prices.card_id_from_db = card_descs.card_id_from_db

