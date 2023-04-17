# 有新包时，批量更新集换社的 card_version_id
#!/bin/bash

read -p "enter password: " PASSWORD

card_id=1579 # 数据库中的 ID，从这个 ID 的卡开始更新 card_version_id
stop_id=1586 # 更新到该 ID 的卡为止
cid=4601 # card_version_id 的值，从这个值开始更新

while [ $card_id -le $stop_id ]; do
  mysql -uroot -p"$PASSWORD" -e "UPDATE my_dtcg.card_prices SET card_version_id = $cid WHERE id = $card_id;"
  card_id=$((card_id+1))
  cid=$((cid+1))
done
