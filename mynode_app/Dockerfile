ARG NODE_VERSION=18
FROM node:${NODE_VERSION}

RUN apt-get update && \
    apt-get install -y vim && \
    rm -rf /var/lib/apt/lists/*

ENV PORT=3001
ENV MESSAGE="Hello Docker!"

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# RUN useradd -m mynode
# USER mynode

HEALTHCHECK --interval=10s --timeout=5s --start-period=5s --retries=3 \
    CMD ["curl", "-f", "http://localhost:3001"]

VOLUME [ "/data" ]

EXPOSE 3001

CMD ["node", "index.js"]