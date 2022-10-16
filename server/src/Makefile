
dep:
	sudo apt-get -y update
	sudo apt-get -y install \
	ca-certificates \
	curl \
	gnupg \
	lsb-release
	sudo mkdir -p /etc/apt/keyrings
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg -f --dearmor -o /etc/apt/keyrings/docker.gpg
	echo "deb [arch=$(shell dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(shell lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
	sudo apt-get -y update
	sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin docker-compose
	sudo service docker start


mon:
	cd ../monitoring
	sudo docker-compose up

run:
	go run server.go