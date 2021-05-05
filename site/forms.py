from flask_wtf import FlaskForm
from wtforms import PasswordField, StringField, SubmitField, TextAreaField, RecaptchaField
from wtforms.validators import Email
import wtforms.validators.InputRequired as DataRequired
from models import User
class LoginForm(FlaskForm):
    username = StringField("Username", validators=[DataRequired()])
    password = PasswordField("Password", validators=[DataRequired()])
    submit = SubmitField("Submit")
    def validate_username(self, form, field):
        if not Users.query.filter_by(username=form.username.data).first():
            raise ValidationError("There is no such user!")

class SignupForm(FlaskForm):
    username = StringField("Username", validators=[DataRequired()])
    email = StringField("Email", validators=[DataRequired(), Email()])
    password = PasswordField("Password", validators=[DataRequired()])
    submit = SubmitField("Submit")
    def validate_username(self, form, field):
        if Users.query.filter_by(username=form.username.data).first():
            raise ValidationError("A user with that name already exists")
class TrashBinForm(FlaskForm):
    line = TextAreaField("Code")
    submit = SubmitField("Submit")
