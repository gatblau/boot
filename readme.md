<a href="https://icon-icons.com/icon/boot-stomp/39342"><img src="boot.png" alt="By Lorc, Delapouite & contributors" width="100" height="100" align="right" /></a>
# Boot 

If an application uses static configuration files, boot can read environment variables and merge them into one or more configuration files.

# Use

given the following *test.cfg* config file:
```shell script
cfg1=${ENV_VAR_1}
cfg2=${ENV_VAR_2:ASD2}
cfg3=${ENV_VAR_3}
cfg4=${ENV_VAR_4}
```

```shell script
# export environment variables
export ENV_VAR_1=AAA
export ENV_VAR_2=BBB
export ENV_VAR_3=CCC
export ENV_VAR_3=DDD

# merge variables in the config file
./boot test.cfg
```

# Download binaries

Platform specific versions can be found [here](./build).

