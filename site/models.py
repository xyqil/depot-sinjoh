from app import db
from werkzeug.security import generate_password_hash, check_password_hash
from flask_login import UserMixin


class User(db.Model, UserMixin):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(80), unique=True)
    password = db.Column(db.String(100))
    theme_preferenece = db.Column(db.String(100), default="dark")
    posts = db.relationship("Post", backref="author", lazy="dynamic")

    def set_pw_hash(self, password):
        self.password = generate_password_hash(password)

    def check_pw_hash(self, password):
        return check_password_hash(self.password, password)


class Post(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    body = db.Column(db.String())
    user_id = db.Column(db.Integer, db.ForeignKey("user.id"))
