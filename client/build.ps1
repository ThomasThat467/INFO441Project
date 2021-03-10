$scriptpath = $MyInvocation.MyCommand.Path
$dir = Split-Path $scriptpath

# Build docker container
docker build -t eguw/planthelper_client $dir
docker push eguw/planthelper_client