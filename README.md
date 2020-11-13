# pushnotification-service

pushnotification-service is a microservice to send push notifications. It's easy to install and use the microservice
on Windows and Unix platforms, since it's a single Go binary. 

*The features are:*
  - Supports both REST and gRPC calls as server
  - Support only Firebase Cloud Messaging to send push notifications
  - Sends a single push token to given token with cli
  - Could be configured easily with a yaml config file
  
### Usage

##### REST call

Use a HTTP client to make a push notification request

```shell script
TOKEN="token to send push notification"
$ curl -XPOST localhost:8080/push -d '{"token": "$TOKEN", "notification": {"body": "hello"}}'
```

Also, call async api which is implemented with go routines.

```shell script
TOKEN="token to send push notification"
$ curl -XPOST localhost:8080/pushAsync -d '{"token": "$TOKEN", "notification": {"body": "hello"}}'
```

##### gRPC call

proto files can be found in "proto/" directory to generate gRPC client.

Do a remote call to "push.PushNotificationService.Send" to with the following paylod

```shell script
TOKEN="token to send push notification"
{"token": "$TOKEN", "notification": {"body": "hello"}}
```

For example, a rpc request with [grpcurl](https://github.com/fullstorydev/grpcurl)
```shell script
TOKEN="token to send push notification"

$ grpcurl --d '{"token": "$TOKEN", "notification": {"body": "hello"}}' -plaintext 127.0.0.1:18080 push.PushNotificationService.Send
```
 

### Install

There a few options of installing pushnotification-service 

#### Docker images

You have two options when you work with docker images. Either you build your own image or you could
mount your config files to existing container.

##### Building your own image

Use existing "pns" image as your base image like below:  

```shell script
FROM bilalekremharmansa/pns:latest
COPY config.yaml /etc/pns.d/
COPY serviceAccount.json /etc/pns.d/

WORKDIR /app
ENTRYPOINT ["./pns"]
CMD ["--help"]
```

Then, all you need is locate your config file in your context,

```
.
├── config.yaml
└── serviceAccount.json
└── Dockerfile
```

Run the following command to build docker image,

```shell script
$ docker build -t your-pns-image .
```

##### Run prebuild docker image

To run REST and gRPC docker image as docker container with default configuration,
firebase service account json file must be mounted onto container. 

Run following command to run pns image, and publish default gRPC port (18080) to docker host network 
```shell script
$ docker run --rm -p 18080:18080 -v /path/service-account.json:/etc/pns.d/serviceAccount.json bilalekremharmansa/pns server
```

#### Build source code

Source code could be compiled with Makefiles easily, run the following command on root directory of the project.

```shell script
$ make -f scripts/Makefile build
```

make command will produce a single binary in build directory, which is named "pns'.

Run push notification service like below,
```shell script
$ build/pns server
```

### TODO

- Kafka support
- Tests

## License

Apache License 2.0, see [LICENSE](https://github.com/bilalekremharmansa/pushnotication-service/blob/master/LICENSE).