-- 获取所有卡牌的价格
SELECT
    SUM(min_price) AS all_min_price,
    SUM(avg_price) AS all_avg_price
FROM
    card_prices

---查询带有 card_desc 表中的 image 字段的 card_prices 表中的数据
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

