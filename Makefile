
build:
	docker build -t goapp .

run:
	docker run -it goapp

test:
	go test -v

is_matrix_squared:
	docker run -it goapp is_matrix_squared $1