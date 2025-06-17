
# Install dependencies
apt-get update && \
    apt-get install -y curl wget git build-essential python3 python3-venv python3-pip jq unzip && \
    rm -rf /var/lib/apt/lists/*

# Install Node.js 18.20.2 and npm 6.14.6
curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g npm@6.14.6 && \
    rm -rf /var/lib/apt/lists/*

# Install Go 1.21
wget https://go.dev/dl/go1.21.10.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.10.linux-amd64.tar.gz && \
    rm go1.21.10.linux-amd64.tar.gz
export PATH="/usr/local/go/bin:${PATH}"

# Install Foundry 
curl -L https://foundry.paradigm.xyz | bash && \
    /root/.foundry/bin/foundryup
export PATH="/root/.foundry/bin:${PATH}"

# Install Poetry (Ubuntu 24.04+ needs --break-system-packages)
pip3 install --break-system-packages -i https://pypi.tuna.tsinghua.edu.cn/simple poetry
