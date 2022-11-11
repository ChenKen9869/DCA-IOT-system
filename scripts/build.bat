@echo off

set localDir=G:\code\go-backend\
cd %localDir%
set GOOS=linux
go build -o ./build/go-backend
swag init
echo Down!
exit
