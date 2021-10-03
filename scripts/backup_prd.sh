#!/bin/bash
DATE=`date +"%Y-%m-%d"`
USER=`whoami`
docker exec coach_coach_db_prod_1 mysqldump coach_db > "/home/$USER/app/backups/coach/$DATE.sql"