build-api:
	docker build -t coach-api:latest -f ./docker/api/Dockerfile.api .
build-front:
	docker build -t coach-front:latest -f ./docker/front/Dockerfile.front .
run-api-container:
	docker run -it -u "1000:1000" -v /etc/group:/etc/group:ro -v /etc/passwd:/etc/passwd:ro -v $(PWD)/backend:/usr/local/go/src/coach -p 8080:8080 coach-api:latest
run-front-container:
	docker run -it -v $(PWD)/front_temp:/usr/local/coach -p 3000:3000 coach-front:latest /bin/sh
run-container:
	docker run -it -v $(PWD):/coach -p 3000:3000 -w /coach coach:latest

run-mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=fg47gh62 -p 3306:3306 -d mysql:8.0

set-mysql-dev-login-info:
	docker exec -it coach_coach_db_dev_1 mysql_config_editor set -u root -p

set-mysql-dev-dump-info:
	docker exec -it coach_coach_db_dev_1 mysql_config_editor set --login-path=mysqldump -u root -p

mysql-backup-dev:
	docker exec -it coach_coach_db_dev_1 mysqldump coach_db > ./backend/db/backup/backup.sql

mysql-restore-dev:
	docker exec -i coach_coach_db_dev_1 mysql coach_db < ./backend/db/backup/backup.sql

set-mysql-prod-login-info:
	docker exec -it coach_coach_db_prod_1 mysql_config_editor set -u root -p

set-mysql-prod-dump-info:
	docker exec -it coach_coach_db_prod_1 mysql_config_editor set --login-path=mysqldump -u root -p

mysql-backup-prod:
	docker exec -it coach_coach_db_prod_1 mysqldump coach_db > ./backend/db/backup/backup.sql

mysql-restore-prod:
	docker exec -i coach_coach_db_prod_1 mysql coach_db < ./backend/db/backup/backup.sql