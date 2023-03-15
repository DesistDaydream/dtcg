# 有新包时，批量更新集换社的 card_version_id
#!/bin/bash

counter=1343
cid=4351

while [ $counter -le 1546 ]
do
  mysql -uroot -plch1382121 -e "UPDATE my_dtcg.card_prices SET card_version_id = $cid WHERE id = $counter;"
  counter=$((counter+1))
  cid=$((cid+1))
done
