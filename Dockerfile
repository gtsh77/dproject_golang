FROM golang:1.8

# Install node prereqs, nodejs and yarn
# Ref: https://deb.nodesource.com/setup_8.x
# Ref: https://yarnpkg.com/en/docs/install
RUN \
  apt-get update && \
  apt-get install -yqq apt-transport-https
RUN \
  echo "deb https://deb.nodesource.com/node_8.x jessie main" > /etc/apt/sources.list.d/nodesource.list && \
  wget -qO- https://deb.nodesource.com/gpgkey/nodesource.gpg.key | apt-key add - && \
  echo "deb https://dl.yarnpkg.com/debian/ stable main" > /etc/apt/sources.list.d/yarn.list && \
  wget -qO- https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
  apt-get update && \
  apt-get install -yqq nodejs yarn && \
  rm -rf /var/lib/apt/lists/*

#env
RUN mkdir -p /app/gorest
WORKDIR /app/gorest

#copy and build golang app
COPY main.go /app/gorest
RUN go get github.com/gorilla/mux
RUN go build

#add vue app
ADD frontend /app/gorest/frontend

#build front
WORKDIR /app/gorest/frontend
RUN npm install
RUN npm run build

WORKDIR /app/gorest

#bind port
EXPOSE 8000:8000