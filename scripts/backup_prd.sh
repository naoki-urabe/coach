#!/bin/bash
DATE=`date +"%Y-%m-%d"`
docker exec -it coach_coach_db_prod_1 mysqldump coach_db > "./backend/db/backup/$DATE.sql"