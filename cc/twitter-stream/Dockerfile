FROM python:latest
run  mkdir /app
copy requirements.txt /app
workdir /app
run pip install -r requirements.txt 
copy . /app
cmd ["python","./streamer.py"]