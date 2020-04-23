# Steve

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Build](https://github.com/BjoernSchilberg/steve/workflows/Build/badge.svg)



- [Steve](#steve)
  - [Build](#build)
  - [Start](#start)
  - [Certificate handling](#certificate-handling)
    - [Convert p12 to pem](#convert-p12-to-pem)
    - [Check](#check)
  - [Tips & Tricks](#tips--tricks)
    - [Using murmur docker instance for developing](#using-murmur-docker-instance-for-developing)
      - [Getting and building docker image](#getting-and-building-docker-image)
      - [Running](#running)
    - [Test sound file](#test-sound-file)

## Build

```shell
go build
```

## Start

```shell
./steve -username steve -server localhost:64738 -key steve.key.pem -certificate steve.crt.pem
```

## Certificate handling

### Convert p12 to pem

```shell
openssl pkcs12 -info -in certificate.p12
openssl pkcs12 -in path.p12 -out steve.crt.pem -clcerts -nokeys
openssl pkcs12 -in path.p12 -out steve.key.pem -nocerts -nodes
```

### Check

```shell
openssl x509 -in steve.crt.pem -text -noout
openssl rsa -in steve.key.pem -check
```

## Tips & Tricks

### Using murmur docker instance for developing

#### Getting and building docker image

```shell
git clone https://github.com/mumble-voip/mumble.git
cd mumble/
docker build -t mumble-voip/murmur .
```

#### Running

```shell
docker run \
-v $HOME/.murmur:/data \
-p 64738:64738 \
-p 64738:64738/udp \
mumble-voip/murmur
```

### Test sound file

```shell
mplayer -ao alsa:device=hw=0.0 kaffee.mp3
```
