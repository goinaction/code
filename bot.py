from flask import Flask, request, jsonify
from bot import analyze_message
from vision import analyze_image

app = Flask(__name__)

@app.route("/message", methods=["POST"])
def message():
    data = request.get_json()
    reply = analyze_message(data["text"])
    return jsonify({"reply": reply})

@app.route("/vision", methods=["GET"])
def vision():
    result = analyze_image("assets/test.jpg")
    return jsonify({"vision": result})

if __name__ == "__main__":
    app.run(debug=True)
