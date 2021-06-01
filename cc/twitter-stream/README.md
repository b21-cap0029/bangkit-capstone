dalam pekerjaan ini ada 3 program yang akan berjalan: aaida-access.py, request-tf.py, dan streamer.py  

streamer.py  
streamer.py berguna untuk menjalankan streaming tweet yang akan disimpan menjadi twitter_fetch.csv  
guna dari twitter_fetch.csv sebagai input untuk program request-tf.py  

Urutan Running Program  
streamer.py>>request.tf>>aaida-access.py 
how to run  
>>>persiapkan file API_KEY.py sebagai penyimpanan API token  
```
CONSUMER_KEY= "**"
CONSUMER_SECRET = "**"
BEARER_TOKEN = "**"
ACCESS_TOKEN = "**"
ACCESS_TOKEN_SECRET= "**"
```
>>>jalankan program streamer.py  
>>>untuk memastikan streamer berjalan, console akan menampilkan tweet yang masuk (streamed)  
>>>selama streamer berjalan, file twitter_fetch akan selalu dimodifikasi(hati-hati)  


request-tf.py  
request-tf.py berguna untuk request prediction ke tensorflow dan menyimpannya ke responses-record.csv  
apabila melebihi treshold tertentu maka akan dicatat  

how to run  
>>>pastikan sudah ada file twitter_fetch.csv (jalankan streamer.py jika tidak ada)  
>>>jalankan request-tf.py  
>>>untuk memastikan request-tf.py berjalan, console akan menampilkan notifikasi (per-entry twitter_fetch.csv)  

aaida-access.py  
digunakan untuk mengirimkan hasil prediction yang disimpan di responses-record.csv  

how to run  
>>>pastikan ada file responses-record.csv  
>>>jalankan program aaida-access.py  
>>>untuk memastikan program selesai dengan semua entry yang ada, akan diberikan notifikasi  
