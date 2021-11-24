FROM gitpod/workspace-mysql

RUN npm install -g wscat && \
    export GITHUB_TOKEN=mKkLNrkKormlvZuouW7zzY9r8Modw70bNbkc && \
    export GOPRIVATE=github.com/aghchan && export GOPROXY=direct && \
    git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf https://github.com/ && \
    sudo apt-get update  && sudo apt-get install -y   redis-server  && sudo rm -rf /var/lib/apt/lists/* && \
    redis-server --daemonize yes