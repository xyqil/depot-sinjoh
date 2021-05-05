FROM python:3.9-alpine

RUN adduser -D server
WORKDIR /home/server
RUN mkdir /etc/db && chown server /etc/db # This is where the DB is stored in Docker
VOLUME /etc/db
COPY requirements.txt .

RUN apk add -U --no-cache --virtual python3 build-base libsass && \
	pip3 install -r requirements.txt && \
	pip3 install gunicorn 

USER server
copy . .

# Remember to change your config file accordingly!
# Note to my devs: please always use Gunicorn for Python!
ENV FLASK_APP app.py

EXPOSE 5000
ENTRYPOINT ["gunicorn", "-b", ":5000", "--access-logfile", "--keyfile key.pem", "--certfile cert.pem", "-", "--error-logfile", "-", "app:app"]
