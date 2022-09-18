# Game-Currency

I use existing libs :

 - Chi Router
 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - Sqlx
 - PostgreSQL


# For Setup After Cloning The Repo:
> cd game-currency
> go mod tidy

# To Do Unit Test :
> go to the package you want to testing then run a command "go test"
> To see the coverage testing in each package, open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package and then you can see the coverage in OUTPUT TAB Console.

# Summary Of Unit Test 
I have done the unit test with 100% coverage, and here are the result :
>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-go1GEgku/go-code-cover game-currency/src/app/use_cases/conversion
>>ok  	game-currency/src/app/use_cases/conversion	0.880s	coverage: 100.0% of statements

>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-go1GEgku/go-code-cover game-currency/src/app/use_cases/currency
>>ok  	game-currency/src/app/use_cases/currency	0.641s	coverage: 100.0% of statements

>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-go1GEgku/go-code-cover game-currency/src/interface/rest/handlers/conversion
>>ok  	game-currency/src/interface/rest/handlers/conversion	0.706s	coverage: 100.0% of statements

>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-go1GEgku/go-code-cover game-currency/src/interface/rest/handlers/currency
>>ok  	game-currency/src/interface/rest/handlers/currency	0.696s	coverage: 100.0% of statements

# for db table :
> In folder db, there is a .sql file with the create table command and insert command. I use postgresql for this case. you can run the command in your sql editor page.
> If you running this project without docker, just make a connection in you local Postgre into your localhost
> If you running this project with docker, make a connection in your local Postgres into 0.0.0.0 and make a database, e.g I use "projek"
> then make new Schema in db "projek", named "game_currency", after that run all the command in .sql file

# the endpoint
> here is the curl for the endpoint :
curl --location --request POST 'http://localhost:8080/api/currency/' \
--header 'x-api-key: attn' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"disena"
}'

curl --location --request POST 'http://localhost:8080/api/conversion/' \
--header 'x-api-key: attn' \
--header 'Content-Type: application/json' \
--data-raw '{
    "currencyIdFrom" : 1,
    "currencyIdTo":1,
    "rate": 200
}'

curl --location --request GET 'http://localhost:8080/api/currency' \
--header 'x-api-key: attn'

curl --location --request GET 'http://localhost:8080/api/conversion/2/1/580' \
--header 'x-api-key: attn'


> i use x-api-key in header, you can fill it with anything in example I fill it with "attn". it just an example if in assumsion the endpoint using a header middleware

> here is the postman link if you want to use postman instead : 
> https://www.getpostman.com/collections/187c901796ba29d3a0d1

# TO RUN THE PROJECT
if you're not using docker, just set the .env file with yoyr database credential, then cd game-currency, do go run main.go

# TO RUN THE PROJECT WITH DOCKER
after clone and do some set up that explained before, do this following actions :
- set database credential in .env

in this part :
DB_HOST=database (recommend to literally use "database" according to the docker-compose.yaml)
DB_PORT=5432  
DB_NAME=projek/your_db_name
DB_USERNAME=your_db_user
DB_PASSWORD=your_db_password
DB_SCHEMA=warung_pintar
DB_SSL_MODE=disable

in this part :
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_db_password
POSTGRES_DB=projek

- cd game-currency, docker-compose up
- deactivate your local postgres which is outside docker
- go to you postgresql db editor (pgAdmin, etc)
- make a new connection to 0.0.0.0
- make a new database, in this project I make a db named "projek"
- in db "projek" make a new schema named "game_currency"
- do all command to make the table and insert, you can see the command in db/account.sql and db/customer.sql
- project ready to use

# MISC
in case there is any error when you try to set up this repo in your local, dont hesitate to call me
my WA +6287825913730
Thankyou
