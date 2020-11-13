# pushnotification-service

pushnotification-service is a microservice to send push notifications. It's easy to install and use the microservice
on Windows and Unix platforms, since it's a single Go binary. 

The features are:
  - Supports both REST and gRPC calls as server
  - Support only Firebase Cloud Messaging to send push notifications
  - Sends a single push token to given token with cli
  - Could be configured easily with a yaml config file
  
  
### Install

There a few options of installing pushnotification-service 

#### Docker images

You have two options when you work with docker images. Either you build your own image or you could
mount your config files to existing container.

##### Building your own image

Use existing "pns" image as your base image like below:  

```
FROM bilalekremharmansa/pns:latest
COPY config.yaml /etc/pushnotification-service/
COPY serviceAccount.json /etc/pushnotification-service/

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

```
$ docker build -t your-pns-image .
```

##### Run prebuild docker image

To run REST and gRPC docker image as docker container with default configuration,
firebase service account json file must be mounted onto container. 

Run following command to run pns image, and publish default gRPC port (18080) to docker host network 
```bash
$ docker run --rm -p 18080:18080 -v /path/service-account.json:/etc/pushnotication-service/serviceAccount.json bilalekremharmansa/pns server
```

#### Build source code

Source code could be compiled with Makefiles easily, run the following command on root directory of the project.

```bash
$ make -f scripts/Makefile build
```

make command will produce a single binary in build directory, which is named "pns'.

Run push notification service like below,
```bash
$ build/pns server
```