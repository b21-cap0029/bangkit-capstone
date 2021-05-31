import json
import time
import datetime
import requests
import csv

filename = 'responses-record.csv'
url = 'localhost/cases/submit'#dummy load dengan HTTP-POST
payload = {} #structure dasar payload

with open (filename,'r') as record:#tidak ada kewajiban menulis file csv
    record_counter = 0
    csvReader = csv.reader(record)
    record_lenght = len(list(csvReader))
    for line in csvReader:
        payload["created_date"] = datetime.datetime.now()
        payload["tweet_id"] = 1## di twitter_fetch dan responses-record tidak ada tweet id DUMMYLOAD
        payload["class"] = "DUMMY"#maksudnya class apa?
        payload["score"] = line.row[2]
        payload["owner_id"] = line.row[0]
        payload["is_claimed"] = False
        payload["is_closed"] = False        
        resp = requests.post(url,json=payload)
        if resp.status_code != 200:
            print("record {} failed to sent".format(record_counter))
        if record_counter == record_lenght:##penanda end of record
            print("end of record")
        record_counter +=1
