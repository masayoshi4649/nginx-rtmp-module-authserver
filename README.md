# nginx-rtmp-module-authserver
Add authentication mechanism to nginx-rtmp-module

## How to use ?

1. Modify NGINX `.conf` file

   - Add `on_publish` `notify_method` to `rtmp > server` block

```

rtmp {
    server {
        listen 1935;
        chunk_size 100;
            
        application live {
            live on;
            record off;

            # ADD HERE !!!
            on_publish http://localhost:1934/;
            notify_method get;
        }
    }
}

```

2. Register `nginx-rtmp-module-authserver.service` to systemctl
```
cd <nginx-rtmp-module-authserver.tar.gz exist directory>
tar -zxvf nginx-rtmp-module-authserver.tar.gz -C /
chmod +x /nginx-rtmp-module-authserver/main
chmod +r /nginx-rtmp-module-authserver/AllowedKey.csv

cd /nginx-rtmp-module-authserver/
mv nginx-rtmp-module-authserver.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable nginx-rtmp-module-authserver
systemctl start nginx-rtmp-module-authserver
```
