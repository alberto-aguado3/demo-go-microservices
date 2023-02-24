## How to run ##
* (if not already started) service docker start
* (Note: you will need at least go 1.17 (suggested go 1.20.1))
* make up
* docker exec -it demo-microservicio_demo-sql_1 mysql -u root -p 
    * Inside, run each command in the terminal, from file "schema.sql"
* go mod vendor
* If you want to debug, download VScode and copy the file "launch.json" to a directory ".vscode/"
* Start a debug session, add breakpoints anywhere.
* Import Postman collection provided, and start trying requests.
* go run test ./... (command to run all tests)