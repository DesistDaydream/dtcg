SELECT
   c_set.pack_id,c_set.pack_prefix,
	card.card_id,
	card_version_id,
	card.card_id - card_version_id AS chazhi,serial,sc_name,rarity,min_price,avg_price
FROM
	card_desc_from_dtcg_dbs card
	LEFT JOIN card_prices price ON price.card_id=card.card_id
	LEFT JOIN card_group_from_dtcg_dbs c_set ON c_set.pack_id=card.card_pack