#prepare for dependencies from tweepy

import sys
import os
import csv
import datetime
import json
import time
import tweepy
from API_KEY import (ACCESS_TOKEN, ACCESS_TOKEN_SECRET, CONSUMER_KEY,
                     CONSUMER_SECRET)
from tweepy import OAuthHandler, Stream
from tweepy.streaming import StreamListener
import re




def deEmojify(text):
    regrex_pattern = re.compile(pattern = "["
        u"\U0001F600-\U0001F64F"  # emoticons
        u"\U0001F300-\U0001F5FF"  # symbols & pictographs
        u"\U0001F680-\U0001F6FF"  # transport & map symbols
        u"\U0001F1E0-\U0001F1FF"
        u"\U0001F700-\U0001F77F"  # alchemical symbols
        u"\U0001F780-\U0001F7FF"  # Geometric Shapes Extended
        u"\U0001F800-\U0001F8FF"  # Supplemental Arrows-C
        u"\U0001F900-\U0001F9FF"  # Supplemental Symbols and Pictographs
        u"\U0001FA00-\U0001FA6F"  # Chess Symbols
        u"\U0001FA70-\U0001FAFF"  # Symbols and Pictographs Extended-A
        u"\U00002702-\U000027B0"  # Dingbats
        u"\U000024C2-\U0001F251"# flags (iOS)
                           "]+", flags = re.UNICODE)
    return regrex_pattern.sub(r'',text)

def pre_process(text):
    # remove links
    text = re.sub("https://t.co/\S*","",text) 

    #remove @<blabla>
    text = re.sub("\@\S+","",text)

    #remove RT marker
    text = re.sub("RT","",text)

    # remove newline
    text = re.sub("\n","",text)

    #  remove digits
    text = re.sub("[0-9]","",text)

    # remove emojis
    text = deEmojify(text)

    return text

#setup listener
class Listener(StreamListener):

    #def on_status(self, status):
     #   print(status.text)


    def on_error(self, status_code):
        if status_code !=420:
            print(status_code)
            return False
        else:
            return True

    
    def on_data(self, data):
        json_tweet = json.loads(data)
        #print(json_tweet)
        id_user = json_tweet['user']['id']
        id_tweet = json_tweet['id']
        name = json_tweet['user']['screen_name']
        text_tweet = json_tweet['text']
        text_tweet = pre_process(text_tweet)
        print(text_tweet)
        with open(filesave,'a', newline='') as f:
            csvWriter = csv.writer(f)
            csvWriter.writerow([id_user,name,id_tweet,text_tweet])

##set boundary entries
language_setup = ['id']
filesave = 'twitter_fetch-text.csv'
keyword_track = ['kecemasan','lelah']
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


#creating the stream
customListener = Listener()
customStream = Stream(auth, listener = customListener)

#membuat filesave
#BEWARE karena disini pakai mode W potensi ilang, ubah modenya untuk siap PROD
with open (filesave, 'w', newline='') as csvFile:
    csvWriter = csv.writer(csvFile)
    csvWriter.writerow(['ID_USER', 'USERNAME','ID_TWEET','TWEET_TEXT'])

#start streaming
#customStream.sample(languages=["id"])
customStream.filter(track=['election','loser'],languages=["en","fr","es","id"])#batas bahasa indonesia wajib pakai track(APIKEYnya tidak dapat akses feed semua tweet)
#filter track dalam list
