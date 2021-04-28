build:
	docker build -t coach:latest .
run-container:
	docker run -it -v $(PWD):/coach -p 3000:3000 -w /coach coach:latest