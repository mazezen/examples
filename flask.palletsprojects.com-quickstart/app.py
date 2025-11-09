"""
Flask Quick Start
"""
import os

from flask import Flask, render_template, redirect, abort
from flask import request
from markupsafe import escape, Markup
from flask import url_for

app = Flask(__name__)

@app.route("/")
def index() -> Markup:
    return escape("index!")

@app.route('/about')
def about():
    return escape("About")

@app.route('/')
def hello_world():
    return "<p>Hello, World!</p>"

@app.route('/hello')
def hello():
    name = request.args.get('name')
    return f"hello, {escape(name)}"

@app.route('/user/<username>')
def show_user_profile(username):
    # show the user profile for that user
    return f"User {escape(username)}"

@app.route('/post/<int:post_id>')
def show_post(post_id):
    # show the post with the given id, the id is an integer
    return f"Post {post_id}"

@app.route('/path/<path:subpath>')
def show_posts(subpath):
    # show the subpath after /path/
    return f"Subpath {escape(subpath)}"

@app.route('/login')
def login():
    return "login"

"""
/
/login
/login?next=/
/user/John%20Doe
"""
with app.test_request_context():
    print(url_for('index'))
    print(url_for('login'))
    print(url_for('login', next=''))
    print(url_for('show_user_profile', username='John Doe'))

@app.route('/projects/')
def projects():
    return "The project page"

@app.route('/ab')
def ab():
    return "The about page"

@app.route('/login2', methods=['GET', 'POST'])
def login2():
    if request.method == 'POST':
        return do_the_login()
    else:
        return show_the_login_form()

def do_the_login():
    return "do_the_login"

def show_the_login_form():
    return "show_the_login_form"

@app.route('/hello2')
@app.route('/hello2/<name>')
def hello2(name=None):
    return render_template('hello.html', person=name)

@app.route('/login3', methods=['POST'])
def login3():
    if request.method == 'POST':
        username = request.form['username']
        password = request.form['password']
        return login3_do(username, password)

def login3_do(username, password):
    return username+password

@app.route('/upload', methods=['POST'])
def upload_file():
    if request.method == 'POST':
        f = request.files['file']
        f.save(f"/Users/bz/coding/examples/flask.palletsprojects.com-quickstart/{f.filename}", )
        return f.filename


@app.route('/re')
def re():
    return redirect(url_for('login4'))

@app.route('/login4')
def login4():
    abort(401)



if __name__ == '__main__':
    app.run(debug=True)
