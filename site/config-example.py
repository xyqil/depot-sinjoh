from app import app
app.secret_key = 'it\'s a secret to everyone'
app.config['HCAPTCHA_ENABLED'] = not app.debug
app.config['HCAPTCHA_SITE_KEY'] = 'yadda yadda'
app.config['HCAPTCHA_SECRET_KEY'] = 'yadda yadda'
app.config['SQLALCHEMY_DATABASE_URI'] = f'sqlite:///{join(getcwd(), "app.db")}'