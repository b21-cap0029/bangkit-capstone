import requests
import json
import csv
import os
import os.path
from itertools import chain

print(os.getcwd())
fileresponse = 'responses-record.csv'# file response
input_fetch = 'twitter_fetch.csv'
#workdir = os.chdir("..")#naik satu folder untuk ambil twitter-fetch
#print(workdir)


#Persiapan Requests API
url="https://tensorflow-serving-4tl56tjpnq-as.a.run.app/v1/models/model:predict"#dummy load untuk setting request tensor 
payload = {}
#REQUEST BUILDER



with open(input_fetch,'r') as f:
    csvReader= csv.reader(f)
    line_counter=0
    for row in csvReader:
        if line_counter == 0:#untuk skip row dan penanda
            line_counter =+ 1
            print("this is first entry")
        else:
            #CREATE REQUEST AND PAYLOAD FOR TENSOR HERE
            print(row[3])
            data_input = row[3]
            request_input = {'instances':[[data_input]]}
            print(request_input)
            json_dumps = json.dumps(request_input)
            json_payload = json.loads(json_dumps)
            print(json_payload)
            #create request and catch the responses prepare the requests
            resp = requests.post(url,json=json_payload)#atur parameter request disini(doc with requests)
            print(resp.text)
            result_dict = resp.json()  #ambil hasilnya dalam bentuk json. cari key response(hasil test dengan curl
            print(result_dict)
            #hasilnya result_dict["predictions"]
            result = result_dict["predictions"]  #untuk mempermudah akses nilai
            result_list = list(chain.from_iterable(result))
            result = result_list[0]
            print(result_list)
            print(result)
            if result >= 0.936:#sementara aja ini kalau resultnya dipasang treshold 0.5 dari 1
                print("ada masukan ke aaida-backend")#verbose just test in text
                with open(fileresponse,'a',newline='') as responses:
                    csvWriter = csv.writer(responses)
                    print(row[0],row[1],row[2],row[3],result)#verbose all record(id_str,username,id_tweet,tweet_text, dan prediction score)(print only)
                    csvWriter.writerow([row[0],row[1],row[2],row[3],result]) 
            else:
                print("tidak ada isinya")#verbose mode
    

