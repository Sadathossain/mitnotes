# MITnotes

This is a simple web note list app made for my project on Institute of Information Technology, Dhaka University. It's written in Golang + Javascript.

## Building the image

### dev

```bash
$ git clone https://github.com/Sadathossain/mitnotes.git
$ cd mitnotes
$ vagrant up --provision
$ vagrant ssh
$ cd /vagrant/mitnotes
$ sudo docker-compose up -d --build
```

### Compile the code

```bash
$ go get -u github.com/Sadathossain/mitnotes
```

#### On OSX

```bash
$ CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w -X main.appVersion=$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match)" -a -installsuffix cgo -o bin/mitnotes .
```

### Build the Container

```bash
$ docker build -t Sadathossain/mitnotes .
# Tag the image if you want
$ docker tag -f sadathossain/mitnotes sadathossain/mitnotes:latest
$ docker push sadathossain/mitnotes
```

## Testing

```bash
$ ./integration_test.sh
```

## Usage

```
Usage of bin/mitnotes:
  -health-check int
           Period to check all connections (default 15)
  -master string
           The connection string to the Redis master as <hostname/ip>:<port> (default "redis-master:6379")
  -master-password string
           The password used to connect to the master
  -slave string
           The connection string to the Redis slave as <hostname/ip>:<port> (default "redis-slave:6379")
  -slave-password string
           The password used to connect to the slave
  -version
           Shows the version
```
