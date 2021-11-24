FROM gitpod/workspace-mysql

# cli socket client
RUN npm install -g wscat
# redis in background
RUN sudo apt-get update  && sudo apt-get install -y redis-server && sudo rm -rf /var/lib/apt/lists/* && redis-server --daemonize yes 
# git workaround for private repo
RUN echo 'export GITHUB_TOKEN=mKkLNrkKormlvZuouW7zzY9r8Modw70bNbkc' && \
    echo 'export GOPRIVATE=github.com/aghchan && export GOPROXY=direct' && \
    git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf https://github.com/
