d:\
cd D:\MAI\computing_systems\comp_systems
chcp 65001
go build -o  res.exe  -ldflags "-s -w" .
cls
@ECHO OFF
echo Проект успешно собран
timeout 5