-- 获取所有卡牌的价格
SELECT
    SUM(min_price) AS all_min_price,
    SUM(avg_price) AS all_avg_price
FROM
    card_prices