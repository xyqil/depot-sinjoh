FROM python:3.9-alpine

RUN adduser -D server
WORKDIR /home/server

COPY ../requirements.txt .

RUN apk add -U --no-cache --virtual python3 && \
	pip3 install -r requirements.txt && \
	pip3 install gunicorn 

USER server
copy . .
# Note to my devs: please always use Gunicorn for Python!
ENV FLASK_APP app.py

EXPOSE 5000
ENTRYPOINT ["gunicorn", "-b", ":5000", "--access-logfile", "-", "--error-logfile", "-", "room:app"]
