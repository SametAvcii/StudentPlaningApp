docker:
	sudo snap install docker

docker-down:
	docker stop studentplan_mysql-dev_1
	docker rm studentplan_mysql-dev_1

migrateup:
	migrate -path database/migration -database "mysql://root:secret@tcp(localhost:3308)/StudentPlanApp?" -verbose up

migratedown:
	migrate -path database/migration -database "mysql://root:secret@tcp(localhost:3308)/StudentPlanApp?" -verbose down

mysql:
	docker-compose up

.PHONY: mysql docker