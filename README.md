# simple basket app with laravel

## Installation
1. Clone Project
> `git clone git@github.com:tmlsergen/simple_basket.git`
2. Starting Docker Containers
> `docker-compose up`
3. You can see the working containers with `docker ps`
```
 
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS        PORTS                               NAMES
42710cbe5be2   app-dev        "/bin/sh -c 'reflex …"   6 minutes ago   Up 1 second   0.0.0.0:4000->4000/tcp              app
6601a58afa67   mysql:latest   "docker-entrypoint.s…"   6 minutes ago   Up 1 second   0.0.0.0:3306->3306/tcp, 33060/tcp   product-api-go_db_1
a82b4a173291   redis:latest   "docker-entrypoint.s…"   6 minutes ago   Up 1 second   0.0.0.0:6379->6379/tcp              product-api-go_redis_1

```
4. You can access container bash with
> `docker exec -ti app_container_id bash`

# Environments
you can see your development enviroment in .env file
```
MYSQL_HOST=db
MYSQL_USER=root
MYSQL_PASSWORD=root
MYSQL_DATABASE=product_api

JWT_STRING=my_super_string

REDIS_HOST=redis
REDIS_PORT=6379
```
