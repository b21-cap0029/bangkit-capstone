FROM python:latest

RUN mkdir /app
COPY requirements.txt /app
WORKDIR /app

RUN pip install -r requirements.txt 
COPY twitter-work/ /app
EXPOSE 9000/tcp 

CMD ["python", "streamer.py"]
