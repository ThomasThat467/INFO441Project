$scriptpath = $MyInvocation.MyCommand.Path
$dir = Split-Path $scriptpath
docker build -t eguw/planthelper_db $dir
docker push eguw/planthelper_db