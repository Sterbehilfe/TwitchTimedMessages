cd TwitchTimedMessages/
go env -w GOOS=linux GOARCH=arm
go build -o ../Build/linux/TwitchTimedMessages ./main.go
go env -w GOOS=windows GOARCH=amd64
go build -o ../Build/win64/TwitchTimedMessages.exe ./main.go
cd ..
cp Settings.json Build/linux/
cp Settings.json Build/win64/
