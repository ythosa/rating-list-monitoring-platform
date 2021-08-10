FROM python:3.8

WORKDIR /usr/src/app

COPY requirements.txt requirements.txt

RUN apt-get install -y libpq-dev
RUN pip3 install -r requirements.txt

COPY . .
