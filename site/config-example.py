from app import app
app.secret_key = 'it\'s a secret to everyone'
app.config['RECAPTCHA_PUBLIC_KEY'] = 'yadda yadda'
app.config['RECAPTCHA_PRIVATE_KEY'] = 'yadda yadda'
