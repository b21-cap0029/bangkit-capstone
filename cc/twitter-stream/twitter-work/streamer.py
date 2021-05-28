#prepare for dependencies from tweepy


from API_KEY import CONSUMER_KEY
from API_KEY import CONSUMER_SECRET
from API_KEY import ACCESS_TOKEN
from API_KEY import ACCESS_TOKEN_SECRET
import tweepy
from tweepy import Stream
from tweepy import OAuthHandler
from tweepy.streaming import StreamListener
import time
import json
import datetime
import csv

##set boundary entries
language = ['id']
#coordinates = []
#filesave ='twitter_fetch'+(datetime.datetime.now().strftime("%Y-%m-%d %H:%M"))+'.csv'
filesave = 'twitter_fetch.csv'
#authentication to twitter
auth = OAuthHandler(CONSUMER_KEY,CONSUMER_SECRET)
auth.set_access_token(ACCESS_TOKEN,ACCESS_TOKEN_SECRET)
api = tweepy.API(auth)

#check API credentials

try:   
    api.verify_credentials()
    print("Authentication Success")
except:
    print("Authentication error")

#setup listener
class Listener(StreamListener):

    #def on_status(self, status):
     #   print(status.text)


    def on_error(self, status_code):
        if status_code !=420:
            print(status_code)
            return False #reconnect stream with 3 tries
        else:
            return True

    
    def on_data(self, data):
        json_tweet = json.loads(data)
        tweet = {'id':json_tweet['user']['id_str'],'screenName':json_tweet['user']['screen_name'],'text:':json_tweet['text']}
        id = json_tweet['user']['id_str']
        name = json_tweet['user']['screen_name']
        text_tweet = json_tweet['text']
        print(text_tweet)
        with open(filesave,'a', newline='') as f:
            csvWriter = csv.writer(f)
            csvWriter.writerow([id,name,text_tweet])




#creating the stream
customListener = Listener()
customStream = Stream(auth, listener = customListener)

#membuat filesave
#BEWARE karena disini pakai mode W potensi ilang, ubah modenya untuk siap PROD
with open (filesave, 'w', newline='') as csvFile:
    csvWriter = csv.writer(csvFile)
    csvWriter.writerow(['ID_STR', 'USERNAME','TWEET_TEXT'])

#start streaming
#customStream.sample(languages=["id"])
customStream.filter(track=['Trump'],languages=["en","fr","es","id"])#batas bahasa indonesia wajib pakai track(APIKEYnya tidak dapat akses feed semua tweet)
#pengerjaan parameter filternya disini

#preparing for prediction request


#catch the prediction responses


#process the prediction