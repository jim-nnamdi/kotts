first_step:
	docker build -t backend .

second_step:
	docker run -it -network kotts_default kotts_backend -p 8080:8080 -d backend:latest

third_step:
	docker run --name mysql -p 3306:3306 -e MYSQL_DATABASE=kotts -e MYSQL_ROOT_PASSWORD=M@etroboomin50 -d mysql:latest

fourth_step:
	docker exec -it mysql mysql -u root -p -h 0.0.0.0 kotts

.PHONY: first_step second_step third_step fourth_step