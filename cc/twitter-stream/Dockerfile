FROM python:latest
RUN  mkdir /app
COPY requirements.txt /app
WORKDIR /app
RUN pip install -r requirements.txt 
COPY . /app
CMD ["python","./streamer.py"]