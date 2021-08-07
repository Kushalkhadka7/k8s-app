# This will run rancher in docker container locally which is accessible in port 80.
run-rancher:
	@sudo docker run --privileged -d --restart=unless-stopped -p 9000:80 -p 443:443 rancher/rancher