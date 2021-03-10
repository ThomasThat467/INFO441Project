GOOS=linux go build -o gateway
docker build -t thomasthat467/msqlconsole ../db/.
docker build -t thomasthat467/gateway .
go clean

docker push thomasthat467/msqlconsole
docker push thomasthat467/gateway

ssh -i "../../../../../../thomas.pem" thomas@planthelper.eguw.me < deploy.sh