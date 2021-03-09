# Build and push DB docker container
$DBBuild = $PSScriptRoot + "\db\build.ps1"
Invoke-Expression $DBBuild

# Build and push gateway docker container
$GatewayBuild = $PSScriptRoot + "\gateway\build.ps1"
Invoke-Expression $GatewayBuild

$MYSQL_ROOT_PASSWORD = 'rootpass'
$SESSIONKEY = "awernasckvznui"
$REDISADDR = 'rServe:6379'
$DSN = "'root:rootpass@tcp(mysql:3306)/db'"
$TLSCERT = '/etc/letsencrypt/live/planthelper.eguw.me/fullchain.pem'
$TLSKEY = '/etc/letsencrypt/live/planthelper.eguw.me/privkey.pem'
# ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me 'sudo docker rm -f $(sudo docker ps -a -q)'
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker network create planthelper"
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker pull eguw/planthelper_gateway"
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker pull eguw/planthelper_db"
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker run -d --network planthelper --name rServe redis"
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker run -d --network planthelper --name mysql -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -e MYSQL_DATABASE=db eguw/planthelper_db"
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker run -d --network planthelper --name gateway -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt:ro -e TLSCERT=$TLSCERT -e TLSKEY=$TLSKEY -e SESSIONKEY=$SESSIONKEY -e DSN=$DSN -e REDISADDR=$REDISADDR eguw/planthelper_gateway"