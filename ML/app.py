import os
import pickle

from flask import Flask, request, jsonify
import joblib

app = Flask(__name__)

# Get the parent directory of the current working directory
parent_directory = os.path.dirname(os.getcwd())
# Load the machine learning model
# Load the machine learning model, vectorizer, and vulnerability names
# with open(parent_directory + '\\lvm-mnit\\ml\\vulnerability_detection_model.pkl', 'rb') as model_file:
#     model = pickle.load(model_file)
#
# with open(parent_directory + '\\lvm-mnit\\ml\\vectorizer.pkl', 'rb') as vectorizer_file:
#     vectorizer = pickle.load(vectorizer_file)
#
# with open(parent_directory + '\\lvm-mnit\\ml\\vulnerability_names.pkl', 'rb') as names_file:
#     vulnerability_names = pickle.load(names_file)


model = joblib.load(parent_directory + '\\lvm-mnit\\ml\\vulnerability_detection_model.joblib')
vectorizer = joblib.load(parent_directory + '\\lvm-mnit\\ml\\vectorizer.joblib')
vulnerability_names = joblib.load(parent_directory + '\\lvm-mnit\\ml\\vulnerability_names.joblib')

@app.route('/predict', methods=['POST'])
def predict_vulnerability():
    try:
        code = request.json['code']
        features = vectorizer.transform([code])
        prediction = model.predict(features)

        # Assuming 1 is 'vulnerable' and 0 is 'not vulnerable'
        result = 'vulnerable' if prediction[0] == 1 else 'not vulnerable'
        # Get the corresponding vulnerability name
        vulnerability_name = vulnerability_names[0] if prediction[0] == 1 else vulnerability_names[1]

        return jsonify({'result': result, 'vulnerability_name': vulnerability_name})

    except Exception as e:
        return jsonify({'error': str(e)}), 400


if __name__ == '__main__':
    app.run(debug=True)
