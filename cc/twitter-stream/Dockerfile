FROM python:latest


RUN apt-get update && apt-get install -y pip3 && \
    apt-get install -y python-dev libffi-dev libssl-dev && \
    pip install tweepy && \
    pip install pyOpenSSL ndg-httpsclient pyasn1



COPY . /twitter

CMD python /twitter/twitter-work/streamer.py