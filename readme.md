<a href="https://icon-icons.com/icon/boot-stomp/39342"><img src="boot.png" alt="By Lorc, Delapouite & contributors" width="100" height="100" align="right" /></a>

# BOOT

*BOOT* is a utility that merges environment variables into (configuration) files. It is designed to pass environment variables to applications with static configuration files running inside Docker containers.

It is an alternative to overlaying the configuration file on the container on start up.

Simply:

- add environment variable placeholders into the files
- export environment variables with the required values
- call *boot* passing the file names to be updated
- launch your app

## Getting Started

1. Download the *boot* utility platform specific versions from [here](./build).

2. Create a bootstrap file to merge the variables and launch the app like follows:

```shell script
# start.sh
/app/boot /app/conf1 /app/conf2 ...
exec "$@"
```

3. Copy boot and the configuration with placeholders to the docker image

4. Call the bootstrap script from the CMD line (mosquitto mqtt broker example):
   
```dockerfile   
CMD ["sh", "/app/start.sh", "/app/mosquitto", "-c", "/app/mosquitto.conf"]
```

## Merging Vars

given the following *test.cfg* config file:

```shell script
cfg1=${ENV_VAR_1}
# ADS2 is a default value
cfg2=${ENV_VAR_2:ASD2}
cfg3=${ENV_VAR_3}
cfg4=${ENV_VAR_4}
```

on the command line run:

```shell script
# export environment variables
# note no variable ENV_VAR_2 is defined 
export ENV_VAR_1=AAA
export ENV_VAR_3=CCC
export ENV_VAR_4=DDD

# merge variables in the config file
./boot test.cfg
```

check the contents of the file:

```shell script
cfg1=AAA
# merged default
cfg2=ASD2
cfg3=CCC
cfg4=DDD
```

**NOTE**: using "PWD" as an environment variable is not allowed as it can retrieve the path of the current process in some OS.

If no variables are exported the execution will fail.
