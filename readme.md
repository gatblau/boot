<a href="https://icon-icons.com/icon/boot-stomp/39342"><img src="boot.png"
     alt="By Lorc, Delapouite & contributors"
     style="float: right; margin-right: 10px; width:100px;" />
</a>

# Boot 

Merge environment variables in static configuration files before a container starts.

# Use

```shell script
# start.sh
# ========
# merges environment variables in the specified files
boot app_config_1.conf app_config_2 ...

# launch the app
exec "app"
```
```dockerfile
# dockerfile
CMD ["start.sh"]
```