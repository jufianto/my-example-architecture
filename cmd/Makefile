isAvailable=$(docker network ls | grep antrian-online | awk 'print $2')
locDockerFile=server/DockerDatabase
nameNetwork="antrian-online"
default:
	@echo "This tools make for antrian online zainab"
hello:
	@echo "hello this tools make by Jufi with a headache"

create-network:
	docker network create $(nameNetwork)

run-app:
	go run main.go

run-db:
	docker run --name dbantrianonline --restart always -d -p 5421:5432 --network $(nameNetwork) dbantrianonline
build-db:
	docker build -t dbantrianonline -f  .

run-adminer:
	docker run --name adminer-dbantrianonline -d -p 8787:8080 --restart always --network $(nameNetwork) adminer

stop-adminer:
	docker stop dbantrianonline