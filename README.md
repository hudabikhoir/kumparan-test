# KUMPARAN TECHNICAL TEST

Sample REST API build using echo server. 

The code implementation was inspired by port and adapter pattern or known as [hexagonal](blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example):

-   **Business**<br/>Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
-   **Modules**<br/>Contains implementation of interfaces that defined at the business (also called as server-side adapters in hexagonal's term)
-   **API**<br/>API http handler or controller (also called user-side adapters in hexagonal's term)

## TECH STACK

Tech stack that I use:
- Golang (`go1.16.2`)
- Echo Framework (`v3.3.10-dev`)
- Mariadb 
- Mongodb
- Redis
- Docker
- Docker-compose
- Git

## Preinstallation
- `git clone https://github.com/hudabikhoir/kumparan-test.git`
- `cd kumparan-test`
- `docker-compose up` to install all environment dependencies
- Adjust config database in `config/config.yaml`
```
database:
  driver: "mysql" #possible value are mongodb or mysql
  address: "127.0.0.1"
  port: 3306
  username: "root"
  password: "toor"
  name: "kumparan-test"
cache:
  driver: "redis"
  address: "127.0.0.1"
  port: 6379
  password: ""
  dbnumber: 1
```

## Installation
- docker build -t kumparango:1.0 .
- docker run -d --name kumparango -p 1323:1323 kumparango:1.0
- open `http://0.0.0.0:1323/` on your browser

or you can run manually
- `./run.sh`
- open `http://0.0.0.0:1323/` on your browser

## How To Consume The API

There are 2 availables API that ready to use:

-   GET `/v1/articles`
-   POST `/v1/articles`

To make it easier please download [Postman](https://www.postman.com/downloads/) app and import [this collection](https://raw.githubusercontent.com/hudabikhoir/kumparan-test/master/other/kumparan.postman_collection.json).
