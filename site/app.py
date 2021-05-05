#!/usr/bin/python3
#-*- coding: utf-8 -*-
from flask import Flask, render_template, send_from_directory, request, redirect
from os.path import join

from os import getcwd
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate
from flask_login import LoginManager, login_user, logout_user, current_user, login_required
from forms import *
from flask_bootstrap import Bootstrap
app = Flask(__name__)
# --CONFIG--
app.config['SQLALCHEMY_DATABASE_URI'] = f'sqlite:///{join(getcwd(), "app.db")}'
debug = True
theme = "dark"
import config
# --END CONFIG--

db = SQLAlchemy(app)
migrate = Migrate(app, db, compare_type=True)
login = LoginManager(app)
bootstrap = Bootstrap(app)
import models
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
  @app.route('/thirdparty/spectrum/actionbutton/dist/<css>')
  def spectrum_actionbutton(css):
    return send_from_directory('thirdparty/spectrum/actionbutton/dist', css)
  @app.route('/thirdparty/spectrum/table/dist/<css>')
  def spectrum_table(css):
    return send_from_directory('thirdparty/spectrum/table/dist', css)
  @app.route('/styles/css/<css>')
  def compiled_scss(css):
    return send_from_directory('styles/css', css)
  @app.route('/styles/<_>/<css>')
  def styles(_, css):
    return send_from_directory('styles', css)
  @app.route('/apis/<_>/<_2>/<css>')
  def apis(_, _2, css):
    return send_from_directory('styles', css)
  @app.route('/assets/images/<asset>')
  def assets(asset):
    return send_from_directory('assets/images',asset)
@app.route('/')
@app.route("/home")
def home():
  page_theme = request.args.get("theme", None)
  if current_user.is_authenticated:
    if not current_user.theme_preferenece == page_theme and page_theme != None or current_user.theme_preferenece == None:
      current_user.theme_preferenece = page_theme
      db.session.add(current_user)
      db.session.commit()
    if current_user.theme_preferenece:
      theme = current_user.theme_preferenece
    else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  return render_template("index.html", theme=theme)
@app.route('/services')
def services():
  page_theme = request.args.get("theme", None)
  if current_user.is_authenticated:
    if not current_user.theme_preferenece == page_theme and page_theme != None or current_user.theme_preferenece == None:
      current_user.theme_preferenece = page_theme
      db.session.add(current_user)
      db.session.commit()
    if current_user.theme_preferenece:
      theme = current_user.theme_preferenece
    else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  return render_template("services.html", theme=theme)
@app.route('/bin/post', methods=["GET","POST"])
@login_required
def bin():
  page_theme = request.args.get("theme", None)
  if current_user.is_authenticated:
    if not current_user.theme_preferenece == page_theme and page_theme != None or current_user.theme_preferenece == None:
      current_user.theme_preferenece = page_theme
      db.session.add(current_user)
      db.session.commit()
    if current_user.theme_preferenece:
      theme = current_user.theme_preferenece
    else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  form = TrashBinForm()
  if form.validate_on_submit():
    p = models.Post(body=form.line.data,author=current_user)
    db.session.add(p)
    db.session.commit()
    return redirect('home')
  return render_template("signup.html",form=form, theme=theme)
@app.route('/faq')
def faq():
  page_theme = request.args.get("theme", None)
  if current_user.is_authenticated:
    if not current_user.theme_preferenece == page_theme and page_theme != None or current_user.theme_preferenece == None:
      current_user.theme_preferenece = page_theme
      db.session.add(current_user)
      db.session.commit()
    if current_user.theme_preferenece:
      theme = current_user.theme_preferenece
    else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  return render_template("faq.html", theme=theme)
@app.route('/login', methods=['GET','POST'])
def loginuser():
  page_theme = request.args.get("theme", None)
  if current_user.is_authenticated:
    if not current_user.theme_preferenece == page_theme and page_theme != None or current_user.theme_preferenece == None:
      current_user.theme_preferenece = page_theme
      db.session.add(current_user)
      db.session.commit()
    if current_user.theme_preferenece:
      theme = current_user.theme_preferenece
    else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  form = SignupForm()
  if form.validate_on_submit():
    user = models.User.query.filter_by(username=form.username.data).first()
    if user.check_pw_hash(form.password.data):
      print("omg someone logged in poggers")
      login_user(user)
      return redirect('home')
  return render_template("signup.html",form=form, theme=theme)
@app.route('/signup', methods=['GET','POST'])
def signup():
  page_theme = request.args.get("theme", None)
  if current_user.is_authenticated:
    if not current_user.theme_preferenece == page_theme and page_theme != None or current_user.theme_preferenece == None:
      current_user.theme_preferenece = page_theme
      db.session.add(current_user)
      db.session.commit()
    if current_user.theme_preferenece:
      theme = current_user.theme_preferenece
    else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  else:
      if not page_theme:
        theme = "dark"
      else:
        theme = page_theme
  form = SignupForm()
  if form.validate_on_submit():
    u = models.User(username=form.username.data)
    u.set_pw_hash(form.password.data)
    db.session.add(u)
    db.session.commit()
    return redirect('/home')
  return render_template("signup.html",form=form, theme=theme)
@app.route('/logout')
def logoutuser():
  logout_user()
  return redirect('/home')
@login.user_loader
def load_user(id):
  return models.User.query.get(id)
if __name__ == '__main__':
  app.run(debug=debug)
