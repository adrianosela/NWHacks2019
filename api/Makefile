all: build

clean:
	rm -rf pkg bin
	rm nwhacks_api*

deploy: push
	az webapp create --resource-group NWHacks2019 --plan nwhackssp --name ezpill --deployment-container-image-name adrianosela/nwhacks_api:laptop_push # deploy to azure

push: tag
	docker push adrianosela/nwhacks_api:laptop_push			# push image to docker hub

tag: dockerbuild
	docker tag nwhacks_api:latest adrianosela/nwhacks_api:laptop_push # tag image with the docker hub repo name

up: dockerbuild down
	docker run -d --name nwhacks_backend -p 8080:80 nwhacks_api	# run docker image as "nwhacks_backend"

run: build
	./nwhacks_api_local 						# run binary

dockerbuild:
	GOOS=linux GOARCH=amd64 go build -a -o nwhacks_api_linux 	# build a golang binary for linux
	docker build -t nwhacks_api:latest . 				# build docker image from this dir

build:
	go build -o nwhacks_api_local 					# build a golang binary for your host machine (e.g. laptop)

down:
	(docker stop nwhacks_backend || true) && (docker rm nwhacks_backend || true)
