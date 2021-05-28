import requests
import csv

url="localhost"#dummy load

with ('twitter-fetch','r') as f:
    csvReader= csv.reader(f)
    

payload="dummy"#dummyload
resp = requests.post(url,data=payload)



resp_dict = resp.json()
print(resp_dict["prediction"])

with open('prediction_res', 'a') as f:
    csv.writer#dummy load
