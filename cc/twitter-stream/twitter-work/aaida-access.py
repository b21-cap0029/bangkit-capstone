import json
import time
from datetime import datetime
import requests
import csv
import os
import os.path
from urllib.parse import urljoin


AAIDA_BACKEND_BASE_URL= os.getenv("AAIDA_BACKEND_BASE_URL")

print(os.getcwd())
filename = 'responses-record.csv'
url = urljoin(AAIDA_BACKEND_BASE_URL, 'cases/submit')  #dummy load dengan HTTP-POST
treshold = 0.936 #defining treshold dari skor tensor 

with open ('responses-record.csv','r') as record:#tidak ada kewajiban menulis file csv
    record_counter = 0
    payload = {} #structure dasar payload
    csvReader = csv.reader(record)
    for line in csvReader:
        print(line)
        payload["tweet_id"] = int(line[2])# index[2] menyimpan tweet_id
        if float(line[4]) >= treshold:#index[4] menyimpan score prediction
            payload["class"] = "Teridentifikasi"
        else:
            payload["class"] = "tidak teridentifikasi"
        payload["score"] = float(line[4])#index[4]menimpan score prediction
        payload["twitter_user_id"] = int(line[0])#menyimpan id twitter user
        payload["is_claimed"] = False
        payload["is_closed"] = False        
        print(dict(payload)) #just test
        resp = requests.post(url,json=payload)
        print(f'{resp.status_code=} {resp.text=}')
"""   
ini kalau aaida-backend sudah siap
        if resp.status_code != 200:
            print("record {} failed to sent".format(record_counter))
        record_counter +=1
"""