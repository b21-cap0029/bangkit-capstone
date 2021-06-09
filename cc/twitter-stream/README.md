# Cara menggunakan Streamer 
1.  persiapkan environment variable untuk menyimpan API_KEY
    ```
    CONSUMER_KEY = "***"
    CONSUMER_SECRET= "***"
    BEARER_TOKEN = "***"

    ACCESS_TOKEN= "***"
    ACCESS_TOKEN_SECRET = "***"
    ```

## build images
```
docker build -t streamer .  
```

## run the images
```
docker run -p 127.0.0.1:9000:9000/tcp --env-file <envfile> <images:tag>
```

## starting the stream
1. open browser
1. input url 127.0.0.1:9000/main
1. in around 120 seconds they will return message "work done" 
