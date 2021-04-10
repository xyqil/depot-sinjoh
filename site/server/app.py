from flask import Flask
app = Flask(__name__)

@app.route('/')
def home():
  return '''
  <!DOCTYPE html>
  <html>
  <head>
  <title>The Hall of Broken Code</title>
  </head>
  <body>
  <style>
  body {
      font-family: sans-serif;
  }
  </style>
  <p>This is the hall of broken code that you will likely never see.</p>
  <h1>This is temporary!</h1>
  <p>If you've stumbled upon this page, JG has broken the server code</p>
  <p>Enjoy!</p>
  '''

if __name__ == '__main__':
  app.run()
