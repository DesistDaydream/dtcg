SELECT
    card.set_id,card.set_prefix,
    card.card_id_from_db,card_version_id,card.card_id_from_db - card_version_id AS chazhi,
    card.serial,card.sc_name,card.rarity,
    min_price,avg_price
FROM
    card_descs card
    LEFT JOIN card_prices price ON price.card_id=card.card_id_from_db
