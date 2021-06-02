dalam pekerjaan ini ada 3 program yang akan berjalan: aaida-access.py, request-tf.py, dan streamer.py  
# Urutan Running Program  
1. streamer.py  
1. request.tf  
1. aaida-access.py 

# streamer.py  
streamer.py berguna untuk menjalankan streaming tweet yang akan disimpan menjadi twitter_fetch.csv  
guna dari twitter_fetch.csv sebagai input untuk program request-tf.py  


### how to run  
persiapkan file API_KEY.py sebagai penyimpanan API token  
```
CONSUMER_KEY= "**"
CONSUMER_SECRET = "**"
BEARER_TOKEN = "**"
ACCESS_TOKEN = "**"
ACCESS_TOKEN_SECRET= "**"
```
1. jalankan program streamer.py  
1. untuk memastikan streamer berjalan, console akan menampilkan tweet yang masuk (streamed)  
1. selama streamer berjalan, file twitter_fetch akan selalu dimodifikasi(hati-hati)  


# request-tf.py  
request-tf.py berguna untuk request prediction ke tensorflow dan menyimpannya ke responses-record.csv  
apabila melebihi treshold tertentu maka akan dicatat  

## how to run  
1. pastikan sudah ada file twitter_fetch.csv (jalankan streamer.py jika tidak ada)  
1. jalankan request-tf.py  
1. untuk memastikan request-tf.py berjalan, console akan menampilkan notifikasi (per-entry twitter_fetch.csv)  

# aaida-access.py  
digunakan untuk mengirimkan hasil prediction yang disimpan di responses-record.csv  

## how to run  
1. pastikan ada file responses-record.csv  
1. jalankan program aaida-access.py  
1. untuk memastikan program selesai dengan semua entry yang ada, akan diberikan notifikasi  
