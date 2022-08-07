first_step:
	docker-compose build

# this happens when you want to update an existing container with the volume for reflection -- docker commit {containerID} {imageID} -- docker run -d -p xxx:xxx --name {cont_name} image:tag --mount source=vol,target={container work_dir} 

second_step:
	docker run -v $(pwd)/vol/kbe:/app --network kotts_default --name kottsapi -p 8080:8080 -d kotts_api:latest

install_SQL:
	docker run --name mysql -p 3306:3306 -e MYSQL_DATABASE=kotts -e MYSQL_ROOT_PASSWORD=M@etroboomin50 -d mysql:latest

fourth_step:
	docker exec -it mysql mysql -u root -p -h 0.0.0.0 kotts

remove_container:
	docker rm -f kottsapi

remove_image:
	docker rmi kotts_api

check_container_logs:
	docker logs kottsapi

cls: 
	clear

.PHONY: first_step second_step install_SQL fourth_step remove_container remove_image check_container_logs cls