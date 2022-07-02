# nginx-rtmp-module-authserver
Add authentication mechanism to nginx-rtmp-module

# How to use ?
Modify .conf
Add [on_publish] and [notify_method]

rtmp {
    server {
        listen 1935;
        chunk_size 100;
            
        application live {
            live on;
            record off;

            # ADD HERE !
            on_publish http://localhost:1934/;
            notify_method get;
        }
    }
}