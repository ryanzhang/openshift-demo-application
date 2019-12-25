import os

from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return """<marquee bgcolor=purple><font color=white><h1><pre>
 	 _   _      _ _           ____ ____ ____  
	| | | | ___| | | ___     / ___/ ___/ ___| 
	| |_| |/ _ \ | |/ _ \   | |  | |  _\___ \ 
	|  _  |  __/ | | (_) |  | |__| |_| |___) |
	|_| |_|\___|_|_|\___/( ) \____\____|____/ 
	                     |/                    
    """

if __name__ == '__main__':
    # Bind to PORT if defined, otherwise default to 5000.
    port = int(os.environ.get('PORT', 5000))
    app.run(host='0.0.0.0', port=port)
