import os

from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return """<marquee bgcolor=purple><font color=white><h1><pre>
 _          _ _               _     _               _       _
| |__   ___| | | ___      ___| |__ (_)_ __   __ _  | |_ ___| | ___ ___  _ __ ___
| '_ \ / _ \ | |/ _ \    / __| '_ \| | _ \ / _ | | __/ _ \ |/ __/ _ \| _  _      | 
| | | |  __/ | | (_) |  | (__| | | | | | | | (_| | | ||  __/ | (_| (_) | | | | | |
|_| |_|\___|_|_|\___( )  \___|_| |_|_|_| |_|\__,_|  \__\___|_|\___\___/|_| |_| |_|
                    |/

    """

if __name__ == '__main__':
    # Bind to PORT if defined, otherwise default to 5000.
    port = int(os.environ.get('PORT', 5000))
    app.run(host='0.0.0.0', port=port)
