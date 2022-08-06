mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_USER=root -e MYSQL_ROOT_PASSWORD=M@etroboomin50 -d mysql:tag

.PHONY: mysql