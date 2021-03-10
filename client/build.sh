docker build -t thomasthat467/client .
docker push thomasthat467/client

ssh -i "../../../../../thomas.pem" thomas@planthelperclient.eguw.me < deploy.sh