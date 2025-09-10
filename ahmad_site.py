from flask import Flask, render_template, request, redirect, url_for

app = Flask(__name__)

# صفحه اصلی
@app.route('/')
def index():
    return "<h1>به وب‌سایت من خوش آمدید!</h1><p>این یک صفحه ساده است.</p>"

# صفحه درباره ما
@app.route('/about')
def about():
    return "<h1>درباره ما</h1><p>اینجا اطلاعاتی درباره ما قرار می‌گیرد.</p>"

# صفحه تماس با ما با فرم ساده
@app.route('/contact', methods=['GET', 'POST'])
def contact():
    if request.method == 'POST':
        name = request.form['name']
        email = request.form['email']
        message = request.form['message']
        print(f"نام: {name}, ایمیل: {email}, پیام: {message}")
        return redirect(url_for('success'))

    return """
        <h1>تماس با ما</h1>
        <form method="post">
            نام: <input type="text" name="name"><br>
            ایمیل: <input type="email" name="email"><br>
            پیام: <textarea name="message"></textarea><br>
            <input type="submit" value="ارسال">
        </form>
    """

# صفحه موفقیت
@app.route('/success')
def success():
    return "<h1>پیام شما با موفقیت ارسال شد!</h1><p><a href='/'>بازگشت به صفحه اصلی</a></p>"

if __name__ == '__main__':
    app.run(debug=True)
