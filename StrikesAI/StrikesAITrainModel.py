import pandas as pd

from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.svm import SVC
from sklearn.metrics import f1_score

from sklearn.model_selection import train_test_split

import pickle

def read_from_csv(path):
    return pd.read_csv(path)


def mapping_isToxic(dataset):
    mapper = {"Toxic": 1, "Not Toxic": 0}
    dataset["is_toxic"] = dataset["is_toxic"].replace(mapper)


def calculate_micro_f_score(y_true, y_pred):
    return f1_score(y_true, y_pred, average='micro')

def predict(svm, x):
    return svm.predict_proba(x)


def train_svm(svm, x_train, y_train):
    return svm.fit(x_train, y_train)


def train():
    df = read_from_csv('toxicity_en.csv')
    mapping_isToxic(df)


    train, test = train_test_split(df, test_size=0.2)

    # train
    text_train = train['text']
    is_toxic_train = train['is_toxic']
        
    # test
    text_validate = test['text']
    is_toxic_validate = test['is_toxic']

    # print(text_validate)

    vectorizer = TfidfVectorizer(lowercase = True, stop_words = {'english'} , token_pattern=r"(?u)\b\w\w+\b|!|\?|\"|\'")
    #     vectorizer = TfidfVectorizer(lowercase = False)
    tfidf_vectorized_train = vectorizer.fit_transform(text_train)
    tfidf_vectorized_validate = vectorizer.transform(text_validate)

    svm = SVC(kernel = 'linear', C = 1.8)
    svm.probability = True
    svm = train_svm(svm, tfidf_vectorized_train, is_toxic_train)
        
    predicted = predict(svm, tfidf_vectorized_validate)

    y_pred = [0 if pred[0] > pred[1] else 1 for pred in predicted]

    score = calculate_micro_f_score(is_toxic_validate, y_pred)

    # saving
    train.to_csv(r'my_train.csv', index=False)

    test.to_csv(r'my_test.csv', index=False)

    with open('model_pkl', 'wb') as f:
        pickle.dump(svm, f)

    print(score)


if __name__ == '__main__':
    train()

