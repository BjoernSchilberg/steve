# Steve

## Start

```shell
./steve -username steve -server localhost:64738 -key steve.key.pem -certificate steve.crt.pem
```

## Convert p12 to pem

```shell
openssl pkcs12 -info -in certificate.p12
openssl pkcs12 -in path.p12 -out steve.crt.pem -clcerts -nokeys
openssl pkcs12 -in path.p12 -out steve.key.pem -nocerts -nodes
```

## Check

```shell
openssl x509 -in steve.crt.pem -text -noout
openssl rsa -in steve.key.pem -check
```

## Tips & Tricks

```shell
mplayer -ao alsa:device=hw=0.0 kaffee.mp3
```
