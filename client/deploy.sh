docker rm -f client

docker pull thomasthat467/client

docker run \
    -d \
    --name client \
    -p 80:80 -p 443:443 \
    -v /etc/letsencrypt:/etc/letsencrypt:ro \
    thomasthat467/client

exit