sampai saat ini ada 3 program yang disiapka
streamer.py:untuk menangani stream twitter sehingga menghasilkan file twitter_fetch.csv
request-tf.py:untuk menangani hasil stream(twitter_fetch.csv) dan mengirimkan request prediction ke tensorflow-serving
hasil akhir request-tf.py menghasilkan file responses-record.csv 
aaida-access.py:untuk menangani hasil yang sudah diprediksi (responses-record.csv) agar dikirim ke aaida-backend