.PHONY: run indexUsers getRawUser updateUser addUser addUser2 getNewUser

HOST=127.0.0.1:8080
API_VER=api/v1

run:
	go run ../.

indexUsers:
	curl -i http://${HOST}/${API_VER}/users

getRawUser:
	curl -i http://${HOST}/${API_VER}/users/1

updateUser:
	curl -i -X PUT -d "name=rick" http://${HOST}/${API_VER}/users/1

addUser: 
	curl -i -X POST -d "name=rachel" -d "email=nueip@staff.com" http://${HOST}/${API_VER}/users

addUser2:
	curl -i -X POST \
		--header "Content-type: application/json" \
		--data '{"email":"joy@staff.com","name":"Joy"}' \
		http://${HOST}/${API_VER}/users

getNewUser:
	curl -i http://${HOST}/${API_VER}/users/2