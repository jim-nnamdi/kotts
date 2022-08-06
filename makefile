kotts_backend:
	docker build -t backend .

kotts_container:
	docker run -it kotts_backend -p 8080:4500 -d backend:latest
	
mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=M@etroboomin50 -d mysql:latest

.PHONY: mysql