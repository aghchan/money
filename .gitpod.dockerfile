FROM gitpod/workspace-mysql

# cli socket client
RUN npm install -g wscat
# install redis
RUN sudo apt-get update  && sudo apt-get install -y redis-server && sudo rm -rf /var/lib/apt/lists/*

