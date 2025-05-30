FROM python:3.7-alpine
COPY ./ /usr/src/app/
WORKDIR /usr/src/app
RUN apk add curl openssl bash --no-cache
RUN curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" \
		&& chmod +x ./kubectl \
		&& mv ./kubectl /usr/local/bin/kubectl \
		&& curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/v3.17.3/scripts/get-helm-3 \
		&& chmod +x get_helm.sh && ./get_helm.sh
