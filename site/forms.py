from flask_wtf import FlaskForm
from wtforms import PasswordField, StringField, SubmitField

class SignupForm(FlaskForm):
    username = StringField("Username")
    password = PasswordField("Password")
    submit = SubmitField("Submit")