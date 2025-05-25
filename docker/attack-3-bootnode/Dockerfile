# Use Ubuntu 22.04 as base image
FROM ubuntu:24.04

# Set non-interactive mode for apt
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && \
    apt-get install -y curl wget git build-essential python3 python3-venv python3-pip jq unzip && \
    rm -rf /var/lib/apt/lists/*

# Install Node.js 18.20.2 and npm 6.14.6
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g npm@6.14.6 && \
    rm -rf /var/lib/apt/lists/*

# Install Go 1.21
RUN wget https://go.dev/dl/go1.21.10.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.10.linux-amd64.tar.gz && \
    rm go1.21.10.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install Foundry 
# RUN curl -L https://foundry.paradigm.xyz | bash && \
#     /root/.foundry/bin/foundryup
# ENV PATH="/root/.foundry/bin:${PATH}"
# too slow, can use pre-installed foundry
COPY .foundry /root/.foundry 
ENV PATH="/root/.foundry/bin:${PATH}"

# Install Poetry (Ubuntu 24.04+ needs --break-system-packages)
RUN pip3 install --break-system-packages -i https://pypi.tuna.tsinghua.edu.cn/simple poetry

# Set workdir
WORKDIR /app

# Copy attack-3-bootnode-code.zip
COPY code/attack-3-bootnode-code.zip /app/code/attack-3-bootnode-code.zip
# Copy node-deploy.zip
COPY node-deploy.zip /app/node-deploy.zip
# Copy query directory
# can use pre-built query binary
COPY query /app/query

# Unzip attack-3-bootnode-code.zip, build geth binary
# RUN unzip /app/code/attack-3-bootnode-code.zip -d /app/code/ && \
#     cd /app/code/attack-3-bootnode-code && make geth
# can use pre-built geth binary
COPY code/attack-3-bootnode-code/build/bin/geth /app/code/attack-3-bootnode-code/build/bin/geth

# Unzip node-deploy.zip, move geth to node-deploy/bin, set up Python venv, install requirements, run cluster script
# use Tsinghua mirror
RUN unzip /app/node-deploy.zip -d /app/ && \
    mkdir -p /app/node-deploy/bin && \
    cp /app/code/attack-3-bootnode-code/build/bin/geth /app/node-deploy/bin/ && \
    cd /app/node-deploy && \
    python3 -m venv /app/node-deploy/venv && \
    . /app/node-deploy/venv/bin/activate && \
    pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple -r requirements.txt && \
    . /app/node-deploy/venv/bin/activate 

RUN cd /app/node-deploy && \
    bash -x ./bsc_cluster.sh prepare

# Default command: start cluster and run monitoring script
CMD bash -c 'cd /app/node-deploy && bash -x ./bsc_cluster.sh start-all && cd /app/query && ./getblock --node=21'
