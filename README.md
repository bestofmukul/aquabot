# Aquabot

Aquabot is a raspberry pi based water sprinkler system written in go. Its primary purpose is to water indoor plants at specified schedule. This project was started as a learning path towards handling electronic components via raspberry pi.
Wiki here provides some insights into the journey.

![](images/Aquabot-assembled.jpeg?raw=true "Aquabot assembled image")

# Install
Pre-req: You should have a go environment setup. 

All you need to do is
```
go get github.com/bestofmukul/aquabot/...
```

# Usage
Build script provided here takes task name as parameter and create a binary compatible to raspberry pi platform. You can then scp the binary to pi and run it.

Circuit schematics and other details are in [wiki](https://github.com/bestofmukul/aquabot/wiki).

* Building aquabot
```
$ ./build.sh cmd/aquabot.go
```

* Scp to pi
```
$ scp aquabot pi@pi.local:~/
```

* Execute binary in pi
```
pi@pi:~ $ ./aquabot -help
  -days uint
        Frequency at which sprinklers will start (default 1)
  -sprinkle duration
        Run sprinklers for this much time in seconds (default 15s)
  -time string
        Time(format hh:mm) at which sprinklers will run (default "10:00")
pi@pi:~ $ ./aquabot --time=20:03 --sprinkle=30s
2017/11/17 20:01:11 Initializing connections...
2017/11/17 20:01:11 Initializing connection RaspberryPi-71B85F9F ...
2017/11/17 20:01:11 Initializing devices...
2017/11/17 20:01:11 Initializing device Relay-2C65B6B5 ...
2017/11/17 20:01:11 Robot AquaBot initialized.
2017/11/17 20:01:11 Starting Robot AquaBot ...
2017/11/17 20:01:11 Starting connections...
2017/11/17 20:01:11 Starting connection RaspberryPi-71B85F9F...
2017/11/17 20:01:11 Starting devices...
2017/11/17 20:01:11 Starting device Relay-2C65B6B5 on pin 11...
2017/11/17 20:01:11 Starting work...
2017/11/17 20:03:00 Starting sprinklers
2017/11/17 20:03:30 Stopping sprinklers
...
```