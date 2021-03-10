docker rm -f rServe
docker network rm sessNetwork
docker rm -f msqlconsole
docker rm -f gateway

docker pull thomasthat467/msqlconsole
docker pull thomasthat467/gateway

export TLSCERT=/etc/letsencrypt/live/planthelper.eguw.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/planthelper.eguw.me/privkey.pem
export SESSIONKEY=sessionkey
export REDISADDR=rServe:6379
export MYSQL_ROOT_PASSWORD=sqlpassword
export DB_NAME=UserDB
export DSN=root:$MYSQL_ROOT_PASSWORD@tcp\(msqlconsole:3306\)/$DB_NAME

docker network create 441ProjectNetwork
docker run -d --name rServe --network 441ProjectNetwork redis

docker run -d \
    -p 3306:3306 \
    --name msqlconsole \
    --network 441ProjectNetwork \
    -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
    -e MYSQL_DATABASE=$DB_NAME \
    thomasthat467/msqlconsole

docker run \
    -d \
    --name gateway \
    --network 441ProjectNetwork \
    -p 443:443 \
    -v /etc/letsencrypt:/etc/letsencrypt:ro \
    -e TLSCERT=$TLSCERT \
    -e TLSKEY=$TLSKEY \
    -e SESSIONKEY=$SESSIONKEY \
    -e REDISADDR=$REDISADDR \
    -e DSN=$DSN \
    thomasthat467/gateway

exit