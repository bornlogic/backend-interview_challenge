
build:
	docker build -t goapp .

run:
	docker run -it goapp

test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

is_matrix_squared:
	docker run -it goapp is_matrix_squared $1