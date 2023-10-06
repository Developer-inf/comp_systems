@ECHO OFF
chcp 65001
cls
go build -o  res.exe  -ldflags "-s -w" .
echo Проект успешно собран