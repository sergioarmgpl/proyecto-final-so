from flask import Flask, request
from flask import jsonify

app = Flask(__name__)

@app.route("/")
def url_root():
    data = { 
        "message":"OK"
    }
    return jsonify(data)

@app.route("/bitacora/n1/<int:n1>/n2/<int:n2>/op/<op>", methods=["GET"])
def writeBitacora(n1,n2,op):
    print(n1,n2,op)
    data = { 
        "n1":n1,
        "n2":n2,
        "op":op
    }
    return jsonify(data)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
