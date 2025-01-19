# petstay
A small software to connect pet care and owners who need the service in GO.

# Build project
docker-compose build --no-cache

# Execute docker
docker-compose up

# Accessing the docker
docker exec -it petstay-db-1 bash  
docker run -it petstay-app sh  

# Run app
docker start petstay-app-1

# Run in interactive mod
docker run -it petstay-app sh

# See dockers logs
docker logs petstay-app-1  
docker logs petstay-db-1