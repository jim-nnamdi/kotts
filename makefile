kotts_backend:
	docker build -t backend .

kotts_container:
	docker run -it kotts_backend -p 8080:7800 -d backend:latest

mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=M@etroboomin50 -d mysql:latest

bind_container_to_database:
	docker exec -it mysql mysql  -u root -p -h 0.0.0.0 kotts

.PHONY: mysql kotts_backend kotts_container