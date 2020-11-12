# TODO

- Build binary for different platforms
    - GOOS=windows go build -o cli.exe
    - GOOS=linux go build -o cli
    - GOARCH=armv7 GOOS=linux go build -o cli-rpi
    
 - implement GetDefaultConfigBaseDirPath for windows