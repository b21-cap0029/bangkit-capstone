FROM ubuntu:latest


RUN apt-get update
RUN apt-get install -y pip
RUN apt-get install -y python-dev libffi-dev libssl-dev 
RUN pip install tweepy
RUN pip install pyOpenSSL ndg-httpsclient pyasn1



COPY . /twitter

CMD python /twitter/twitter-work/streamer.py