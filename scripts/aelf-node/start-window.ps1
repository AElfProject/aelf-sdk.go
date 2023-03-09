wget https://github.com/microsoftarchive/redis/releases/download/win-3.2.100/Redis-x64-3.2.100.zip -OutFile  redis.zip
Expand-Archive -Path redis.zip -DestinationPath redis ;
$job1 = Start-Job -ScriptBlock { cd D:\a\1\s\redis; .\redis-server.exe; }
sleep 30
mkdir -p C:\Users\VssAdministrator\AppData\Local\aelf\keys
cp -r scripts\aelf-node\keys\* C:\Users\VssAdministrator\AppData\Local\aelf\keys;
wget https://github.com/AElfProject/AElf/releases/download/v1.2.3/aelf.zip -OutFile  aelf.zip ;
Expand-Archive -Path aelf.zip -DestinationPath aelf ;
cp scripts\aelf-node\appsettings.json  aelf\aelf\appsettings.json ;
cp scripts\aelf-node\appsettings.MainChain.TestNet.json  aelf\aelf\appsettings.MainChain.TestNet.json ;
cd aelf/aelf
$job2 = Start-Job -ScriptBlock { cd D:\a\1\s\aelf\aelf; dotnet AElf.Launcher.dll; } 
sleep 60
cd D:\a\1\s
go mod download
cd client/ 
go build
cd ../dto/
go build
cd ../model/
go build
cd ../utils/
go build
cd ../protobuf/generated/
go build
cd ../../test
go test
