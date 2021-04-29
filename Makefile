build-api:
	docker build -t coach-api:latest -f ./docker/api/Dockerfile.api .
build-front:
	docker build -t coach-front:latest -f ./docker/front/Dockerfile.front .
run-api-container:
	docker run -it -u "1000:1000" -v /etc/group:/etc/group:ro -v /etc/passwd:/etc/passwd:ro -v $(PWD)/backend:/usr/local/go/src/coach -p 8080:8080 -w /usr/local/go/src/coach coach-api:latest
run-front-container:
	docker run -it -u "1000:1000" -v /etc/group:/etc/group:ro -v /etc/passwd:/etc/passwd:ro -v $(PWD)/front:/usr/local/coach -p 3000:3000 -w /usr/local/coach coach-front:latest /bin/sh
run-container:
	docker run -it -v $(PWD):/coach -p 3000:3000 -w /coach coach:latest
