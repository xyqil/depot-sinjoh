from app import db
from werkzeug.security import generate_password_hash, check_password_hash
class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(80), unique=True)
    password = db.Column(db.String(100))
    def set_pw_hash(self, password):
        self.password = generate_password_hash(password)
    def check_pw_hash(self, password):
        return check_password_hash(password, self.password)