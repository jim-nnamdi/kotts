first_step:
	docker-compose build

second_step:
	docker run  --network kotts_default --name kottsapi -p 8080:8080 -d kotts_api:latest

third_step:
	docker run --name mysql -p 3306:3306 -e MYSQL_DATABASE=kotts -e MYSQL_ROOT_PASSWORD=M@etroboomin50 -d mysql:latest

fourth_step:
	docker exec -it mysql mysql -u root -p -h 0.0.0.0 kotts

.PHONY: first_step second_step third_step fourth_step