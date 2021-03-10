# Build and push client docker container
$ClientBuild = $PSScriptRoot + ".\build.ps1"
Invoke-Expression $ClientBuild

$TLSCERT = '/etc/letsencrypt/live/planthelperclient.eguw.me/fullchain.pem'
$TLSKEY = '/etc/letsencrypt/live/planthelperclient.eguw.me/privkey.pem'
# ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me 'sudo docker rm -f $(sudo docker ps -a -q)'
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker network create planthelper"
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker pull eguw/planthelper_client"
ssh -i C:\Users\Eric\INFO441_PlantHelper.pem ec2-user@planthelper.eguw.me "sudo docker run -d --network planthelper --name client -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt:ro -e TLSCERT=$TLSCERT -e TLSKEY=$TLSKEY eguw/planthelper_client"