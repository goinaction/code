# ai_analyzer.py
def analyze_message(message):
    if "سلام" in message:
        return "سلام احمد جان! آماده‌ام برای فرم‌کُش کردن."
    elif "پروژه" in message:
        return "پروژه‌ات جهانیه! فقط بزن تا اجرا بشه."
    else:
        return "پیامت دریافت شد. دارم تحلیلش می‌کنم..."
