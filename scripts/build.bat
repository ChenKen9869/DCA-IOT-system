@echo off

set localDir=G:\code\go-backend\
cd %localDir%
swag init
set GOOS=linux
go build -o ./build/go-backend
echo Down!
exit
