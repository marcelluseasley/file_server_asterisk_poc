stop_fs_poc:
	-docker rm -f fs_poc

create_test_file_server_app: stop_fs_poc
	CGO_ENABLED=0 go build -a fileserver/file_server.go
	CGO_ENABLED=0 go build -o client

create_test_file_server_container: create_test_file_server_app
	docker build --no-cache -f Dockerfile -t fserver-fs-poc ./

launch_test_file_server_container: create_test_file_server_container
	docker run -d -p 8088:8088 \
	--name fs_poc \
	fserver-fs-poc