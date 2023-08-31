@echo off

set localDir=G:\code\DCA-IOT-system\
cd %localDir%
swag init
set GOOS=linux
go build -o ./build/dca
echo Down!
exit
