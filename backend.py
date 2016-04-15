from flask import Flask, request, send_from_directory
app = Flask(__name__)

@app.route("/scripts/<path:path>")
def send_js(path):
    return send_from_directory('./static/scripts', path)

@app.route("/")
def index():
    with open('static/index.html', 'r') as html:
        return "".join(html.readlines())

if __name__ == '__main__':
    app.run()
