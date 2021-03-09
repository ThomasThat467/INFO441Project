$scriptpath = $MyInvocation.MyCommand.Path
$dir = Split-Path $scriptpath

# Compile for linux
Set-Location $dir
$env:GOOS = "linux"
go build -o gateway
Write-Host "gateway successfully compiled"

# Compile for windows
# Remove-Item Env:\GOOS
# $env:GOOS="windows"
# go build -o gateway.exe

# Build docker container
docker build -t eguw/planthelper_gateway $dir
docker push eguw/planthelper_gateway
go clean