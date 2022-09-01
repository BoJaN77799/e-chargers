from http.server import BaseHTTPRequestHandler, HTTPServer

import json
import pickle
import numpy as np
from sklearn.feature_extraction.text import TfidfVectorizer

from StrikesAITrainModel import read_from_csv, predict

hostName = "localhost"
serverPort = 50006

def generate_handler():

    class MyHandler(BaseHTTPRequestHandler):

        svm = None
        train = None

        def load(self):
            with open('model_pkl', 'rb') as f:
                self.svm = pickle.load(f)

            self.train = read_from_csv('my_train.csv')

        def predict_proba(self, text):
            
            # train
            text_train = self.train['text']
                
            # making vectorizer (dictionary)
            vectorizer = TfidfVectorizer(lowercase = True, stop_words = {'english'} , token_pattern=r"(?u)\b\w\w+\b|!|\?|\"|\'")
            vectorizer.fit_transform(text_train)
            tfidf_vectorized_validate = vectorizer.transform(np.array([text]))
                
            return predict(self.svm, tfidf_vectorized_validate)

        def do_POST(self):
            if self.path == "/prediction":
                # getting post request body
                req_content_len = int(self.headers.get('Content-Length'))
                post_body = self.rfile.read(req_content_len)
                recension_text = post_body.decode('utf8').replace("'", '"')

                # call predict
                result = self.predict_proba(recension_text)
                result_list = result.tolist()[0]
                
                # adjust response
                self.send_response(200)
                self.send_header("Content-type", "application/json")
                self.end_headers()
                self.wfile.write(("[%s, %s]" % (str(result_list[0]), str(result_list[1]))).encode())
            else:
                self.send_error(404)
        
    return MyHandler 

if __name__ == "__main__":
    handler = generate_handler()
    handler.load(handler)      
    webServer = HTTPServer((hostName, serverPort), handler)
    print("Server started http://%s:%s" % (hostName, serverPort))

    try:
        webServer.serve_forever()
    except KeyboardInterrupt:
        pass

    webServer.server_close()
    print("Server stopped.")