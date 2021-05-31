import requests
import json
import csv
import os
import os.path

print(os.getcwd())
fileresponse = 'responses-record.csv'# file response
#workdir = os.chdir("..")#naik satu folder untuk ambil twitter-fetch
#print(workdir)


#Persiapan Requests API
url="http://35.247.137.111:8501/v1/models/model:predict"#dummy load untuk setting request tensor 
payload = {}
#REQUEST BUILDER



with open('twitter_fetch.csv','r') as f:
    csvReader= csv.reader(f)
    line_counter=0
    for row in csvReader:
        if line_counter == 0:#untuk skip row dan penanda
            line_counter =+ 1
            print("this is first entry")
        else:
            #CREATE REQUEST AND PAYLOAD FOR TENSOR HERE
            payload["instances"] = [row[2]]
            json_payload = json.dumps(payload)
            #create request and catch the responses prepare the requests
            resp = requests.post(url,json=json_payload)#atur parameter request disini(doc with requests)
            result_dict = resp.json()  #ambil hasilnya dalam bentuk json. cari key response(hasil test dengan curl
            #hasilnya result_dict["predictions"]
            result = result_dict["predictions"]  #untuk mempermudah akses nilai
            if result >= 0.5:#sementara aja ini kalau resultnya dipasang treshold 0.5 dari 1
                with open(fileresponse,'a',newline='') as responses:
                    csvWriter = csv.writer(responses)
                    print(row[0],row[1],row[2],result)#verbose all record(id_str,username,tweet_text dan prediction score)
                    #csvWriter.writerow([row[0],row[1],result]) #opsi dengan line ini hanya menyimpan id_str,username, dan prediction score
    

