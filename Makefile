build-api:
	docker build -t coach-api:latest -f ./docker/api/Dockerfile.api .
build-front:
	docker build -t coach-front:latest -f ./docker/front/Dockerfile.front .
run-api-container:
	docker run -it -u `id -u $(USER)` -v $(PWD)/backend:/go/src/coach -p 8080:8080 -w /coach coach-api:latest
run-front-container:
	docker run -it -u `id -u $(USER)` -v $(PWD)/front:/coach -p 3000:3000 -w /coach coach-front:latest
run-container:
	docker run -it -v $(PWD):/coach -p 3000:3000 -w /coach coach:latest
