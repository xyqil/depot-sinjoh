from flask_wtf import FlaskForm
from wtforms import PasswordField, StringField, SubmitField, TextAreaField, RecaptchaField
from wtforms.validators import DataRequired

class SignupForm(FlaskForm):
    username = StringField("Username", validators=[DataRequired()])
    password = PasswordField("Password", validators=[DataRequired()])
    captcha = RecaptchaField()
    submit = SubmitField("Submit")
class TrashBinForm(FlaskForm):
    line = TextAreaField("Code")
    captcha = RecaptchaField()
    submit = SubmitField("Submit")
