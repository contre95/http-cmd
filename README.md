Description

Small webserver to receive http calls and execute commands in the host. I use this solution to execute commands from Home Assistant into the different linux machines I own.

### Home Assistant config

```yaml
rest_command:
  reset_shader:
    url: "http://<your-ip>:8080/reset"
    content_type: "application/x-www-form-urlencoded"
    method: "post"
  set_shader_red:
    url: "http://<your-ip>:8080/shade"
    content_type: "application/x-www-form-urlencoded"
    method: "post"
    payload: "shader=red.glsl"
  set_shader_blfilter:
    url: "http://<your-ip>:8080/shade"
    content_type: "application/x-www-form-urlencoded"
    method: "post"
    payload: "shader=blue-light-filter.glsl"
  set_shader_orange:
    url: "http://<your-ip>:8080/shade"
    content_type: "application/x-www-form-urlencoded"
    method: "post"
    payload: "shader=screenShader.frag"
```

### Systemd unit file 
```toml
[Unit]
Description=My Web Server
After=network.target

[Service]
Type=simple
ExecStart=/path/to/your/webserver/executable

[Install]
WantedBy=multi-user.target
```
