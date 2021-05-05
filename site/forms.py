from flask_wtf import FlaskForm
from wtforms import PasswordField, StringField, SubmitField, TextAreaField
from wtforms.validators import Email
from wtforms.validators import InputRequired as DataRequired
from models import User
from app import hcaptcha
class LoginForm(FlaskForm):
    username = StringField("Username", validators=[DataRequired()])
    password = PasswordField("Password", validators=[DataRequired()])
    submit = SubmitField("Submit")
    def validate_username(self, form, field):
        if not Users.query.filter_by(username=form.username.data).first():
            raise ValidationError("There is no such user!")
        if not hcaptcha.verify():
            raise ValidationError("Please do the hCaptcha")

class SignupForm(FlaskForm):
    username = StringField("Username", validators=[DataRequired()])
    email = StringField("Email", validators=[DataRequired(), Email()])
    password = PasswordField("Password", validators=[DataRequired()])
    submit = SubmitField("Submit")
    def validate_username(self, form, field):
        if Users.query.filter_by(username=form.username.data).first():
            raise ValidationError("A user with that name already exists")
        if not hcaptcha.verify():
            raise ValidationError("Please do the hCaptcha")
class TrashBinForm(FlaskForm):
    line = TextAreaField("Code")
    submit = SubmitField("Submit")
    def validate_line(self, form, field):
        if not hcaptcha.verify():
            raise ValidationError("Please do the hCaptcha")