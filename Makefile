build:
	docker build -t coach:latest .
run-container:
	docker run -it -v $(PWD):/coach -w /coach coach:latest