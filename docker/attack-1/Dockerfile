# Use Ubuntu 22.04 as base image
FROM ubuntu:22.04

# Set non-interactive mode for apt
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && \
    apt-get install -y curl wget git build-essential python3 python3-venv python3-pip jq && \
    ln -sf /usr/bin/python3 /usr/bin/python3 && \
    rm -rf /var/lib/apt/lists/*

# Install Node.js 18.20.2 and npm 6.14.6
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get update && \
    apt-get install -y nodejs && \
    npm install -g npm@6.14.6 && \
    rm -rf /var/lib/apt/lists/*

# Install Go 1.18
RUN wget https://go.dev/dl/go1.18.10.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.18.10.linux-amd64.tar.gz && \
    rm go1.18.10.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install Foundry 
# RUN curl -L https://foundry.paradigm.xyz | bash && \
#     /root/.foundry/bin/foundryup
# ENV PATH="/root/.foundry/bin:${PATH}"
# too slow, use pre-installed foundry
COPY .foundry /root/.foundry 
RUN chmod +x /root/.foundry/bin/foundryup && \
    /root/.foundry/bin/foundryup
ENV PATH="/root/.foundry/bin:${PATH}"

# Install Poetry
RUN pip3 install poetry

# Set workdir
WORKDIR /app

# Copy attack-1-code.zip
COPY code/attack-1-code.zip /app/code/attack-1-code.zip
# Copy node-deploy.zip
COPY node-deploy.zip /app/node-deploy.zip
# Copy query directory
COPY query /app/query

# Unzip attack-1-code.zip, build geth binary
RUN apt-get update && apt-get install -y unzip make && rm -rf /var/lib/apt/lists/* && \
    unzip /app/code/attack-1-code.zip -d /app/code/ && \
    cd /app/code/attack-1-code && make geth

# Unzip node-deploy.zip, move geth to node-deploy/bin, set up Python venv, install requirements, run cluster script
RUN apt-get update && apt-get install -y unzip && rm -rf /var/lib/apt/lists/* && \
    unzip /app/node-deploy.zip -d /app/ && \
    mkdir -p /app/node-deploy/bin && \
    cp /app/code/attack-1-code/build/bin/geth /app/node-deploy/bin/ && \
    cd /app/node-deploy && \
    python3 -m venv /app/node-deploy/venv && \
    apt-get update && apt-get install -y python3.12-venv && rm -rf /var/lib/apt/lists/* && \
    . /app/node-deploy/venv/bin/activate && \
    pip3 install -r requirements.txt

# Default command: start cluster and run monitoring script
CMD bash -c 'cd /app/node-deploy && bash -x ./bsc_cluster.sh reset && cd /app/query && go run main.go --node=21'
