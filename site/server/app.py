#!/usr/bin/python3
#-*- coding: utf-8 -*-
from flask import Flask
app = Flask(__name__)

@app.route('/')
def home():
  # @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
  # @                CAUTION:                    @ 
  # @  This page is meant for if the server is   @
  # @  well, unable to serve things, it's very   @
  # @  important that no other files are         @
  # @  included here and that there aren't       @
  # @  any things that may prohibit the page     @
  # @  from loading.                             @
  # @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
  # *phew* I finally got that ASCII warning sign right
  
  return '''
  <!DOCTYPE html>
  <html>
  <body>
  <h1>This is the hall of broken code that you will likely never see.</h1>
  <p>This is temporary!</p>
  <p>If you've stumbled upon this page, someone broke the server code lol</p>
  <p>Enjoy!</p>
  </body>
  </html>
  '''

if __name__ == '__main__':
  app.run()
