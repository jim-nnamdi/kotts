first_step:
	docker-compose build

second_step:
	docker run --name kottsapi -p 8080:8080 -d kotts_api:latest

third_step:
	docker commit kottsapi kotts_api:latest 

fourth_step: 
	docker run -d -p 8080:8080 --name kottsapi kotts_api:latest --mount source=/vol/kbe target=/app 

# [development_usage]
install_SQL:
	docker run --name kottdb -p 3306:3306 -h kott.czve4izeamxt.us-east-2.rds.amazonaws.com -e MYSQL_DATABASE=kottdb -e MYSQL_ROOT_PASSWORD=Metroboomin50 -d mysql:8

# [production_usage]
use_rds:
	docker run --name mysql -h kott.czve4izeamxt.us-east-2.rds.amazonaws.com -p 3306:3306 -e MYSQL_DATABASE=kottdb -e MYSQL_ROOT_PASSWORD=Metroboomin50 -d mysql:latest

fifth_step:
	docker exec -it mysql mysql -u root -p -h 0.0.0.0 kotts

remove_container:
	docker rm -f kottsapi

remove_image:
	docker rmi kotts_api

check_container_logs:
	docker logs kottsapi

cls: 
	clear

.PHONY: first_step second_step third_step install_SQL fourth_step fifth_step remove_container remove_image check_container_logs cls use_rds