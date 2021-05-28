import requests
import csv
import os
import os.path

print(os.getcwd())
#workdir = os.chdir("..")#naik satu folder untuk ambil twitter-fetch
#print(workdir)
url="localhost"#dummy load untuk setting request tensor


#path = "E:\BANGKIT\Project\bangkit-capstone\cc\twitter-stream" ##belum bisa akses relative path sama belum bisa buat workpathnya(aulah)
with open('twitter_fetch.csv','r') as f:
    csvReader= csv.reader(f)
    line_counter=0
    for row in csvReader:
        if line_counter == 0:#untuk skip row dan penanda
            print("this is first entry")
        else:
            #create request and catch the responses
            payload = row[2]
            resp = requests.post(url,data=payload)#atur parameter request disini(doc with requests)
            result = resp.responses
            if result >= 0.5:#sementara aja ini kalau resultnya dipasang treshold 0.5 dari 1
                with open('responses-record.csv','a') as responses:
                    csvWriter = csv.writer(responses)
                    csvWriter.writerow([row[0],row[1],result])
    

