FROM ubuntu:20.04

RUN apt update \
    && apt install -y ca-certificates openssh-client \
    wget curl iptables supervisor \
    && rm -rf /var/lib/apt/list/*

RUN apt install -y apt-transport-https \
    && curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list \
    && apt update 


ENV DOCKER_CHANNEL=stable \
	DOCKER_VERSION=20.10.9 \
	DOCKER_COMPOSE_VERSION=1.29.2 \
	DEBUG=false \
  DEBIAN_FRONTEND=noninteractive


RUN apt update

RUN apt install -y ca-certificates curl gnupg lsb-release

RUN mkdir -p /etc/apt/keyrings

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg

RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

RUN apt update

RUN apt install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin



RUN curl -L "https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
	&& chmod +x /usr/local/bin/docker-compose && docker-compose version


RUN apt update
RUN apt install -y kubectl
RUN apt -y install python3-pip
RUN apt -y install git
RUN apt -y install vim
RUN pip3 install flask
RUN pip3 install pandas
RUN pip3 install python-dotenv
RUN pip3 install requests
RUN pip3 install pyyaml


RUN mkdir /root/.kube

COPY . /home/NPIA_server
WORKDIR /home/NPIA_server

VOLUME /var/lib/docker
EXPOSE 7331

CMD ["python3","NPIA_remote.py"]
