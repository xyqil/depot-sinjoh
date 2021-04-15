from flask import Flask
app = Flask(__name__)

@app.route('/')
def home():
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

if __name__ == '__main__':
  app.run()
