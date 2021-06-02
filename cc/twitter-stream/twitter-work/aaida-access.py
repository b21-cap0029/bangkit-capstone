import json
import time
from datetime import datetime
import requests
import csv
import os
import os.path

print(os.getcwd())
filename = 'responses-record.csv'
url = 'localhost/cases/submit'#dummy load dengan HTTP-POST
payload = {} #structure dasar payload
treshold = 0.936 #defining treshold dari skor tensor 

with open ('responses-record.csv','r') as record:#tidak ada kewajiban menulis file csv
    record_counter = 0
    csvReader = csv.reader(record)
    for line in csvReader:
        print(line)
        now = datetime.now()#input waktu created
        calendar = now.strftime("%d/%m/%Y %H:%M:%S")
        print(calendar)
        payload["created_date"] = calendar
        payload["tweet_id"] = line[2]# index[2] menyimpan tweet_id
        if float(line[4]) >= treshold:#index[4] menyimpan score prediction
            payload["class"] = "Teridentifikasi"
        else:
            payload["class"] = "tidak teridentifikasi"
        payload["score"] = line[4]#index[4]menimpan score prediction
        payload["twitter_user_id"] = line[0]#menyimpan id twitter user
        payload["owner_id"] = 0 #dummy values
        payload["is_claimed"] = False
        payload["is_closed"] = False        
        #resp = requests.post(url,json=payload)
        print(dict(payload))#just test
"""   
ini kalau aaida-backend sudah siap
        if resp.status_code != 200:
            print("record {} failed to sent".format(record_counter))
        record_counter +=1
"""