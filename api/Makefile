all: build

clean:
	rm -rf pkg bin
	rm nwhacks_api*

deploy: dockerbuild down
	docker run -d --name nwhacks_backend -p 8080:80 nwhacks_api	# run docker image as "nwhacks_backend"

up: build
	./nwhacks_api_local 						# run binary

dockerbuild:
	GOOS=linux GOARCH=amd64 go build -a -o nwhacks_api_linux 	# build a golang binary for linux
	docker build -t nwhacks_api . 					# build docker image from this dir

build:
	go build -o nwhacks_api_local 					# build a golang binary for your host machine (e.g. laptop)

down:
	(docker stop nwhacks_backend || true) && (docker rm nwhacks_backend || true)
