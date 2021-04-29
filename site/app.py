#!/usr/bin/python3
#-*- coding: utf-8 -*-
from flask import Flask, render_template, send_from_directory
app = Flask(__name__)

@app.route('/oh_dear_what_a_blunder_ive_made')
def uhohspeghettios():
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
debug = True
if debug:
  @app.route('/thirdparty/spectrum/vars/dist/<css>')
  def spectrum_vars(css):
    return send_from_directory('thirdparty/spectrum/vars/dist', css)
  @app.route('/thirdparty/spectrum/page/dist/<css>')
  def spectrum_page(css):
    return send_from_directory('thirdparty/spectrum/page/dist', css)
  @app.route('/thirdparty/spectrum/typography/dist/<css>')
  def spectrum_typography(css):
    return send_from_directory('thirdparty/spectrum/typography/dist', css)
  @app.route('/thirdparty/spectrum/icon/dist/<css>')
  def spectrum_icon(css):
    return send_from_directory('thirdparty/spectrum/icon/dist', css)
  @app.route('/thirdparty/spectrum/button/dist/<css>')
  def spectrum_button(css):
    return send_from_directory('thirdparty/spectrum/button/dist', css)
@app.route('/')
def home():
  return render_template("index.html")
if __name__ == '__main__':
  app.run(debug=debug)
