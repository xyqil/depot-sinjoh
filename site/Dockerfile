FROM python:3.9-alpine

RUN adduser -D server
WORKDIR /home/server
COPY key.pem key.pem
COPY cert.pem cert.pem
RUN chown server key.pem 
RUN chown server cert.pem
RUN chmod +r key.pem
RUN mkdir /etc/db && chown server /etc/db # This is where the DB is stored in Docker
VOLUME /etc/db
COPY requirements.txt .

RUN apk add -U --no-cache --virtual python3 build-base libsass && \
	pip3 install -r requirements.txt && \
	pip3 install gunicorn 

USER server
copy . .

# Comment if not using SSL

# Remember to change your config file accordingly!
# Note to my devs: please always use Gunicorn for Python!
ENV FLASK_APP app.py

EXPOSE 5000

ENTRYPOINT ["gunicorn", "-b", ":5000", "--access-logfile", "-", "--certfile", "cert.pem", "--keyfile", "key.pem", "--error-logfile", "-", "app:app"]
# Command if not using SSL
#ENTRYPOINT ["gunicorn", "-b", ":5000", "--access-logfile", "-", "--error-logfile", "-", "app:app"]
# Uncomment if not using SSL
