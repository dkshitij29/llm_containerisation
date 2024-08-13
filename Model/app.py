from flask import Flask, request, jsonify
from transformers import LlamaForCausalLM, LlamaTokenizer

app = Flask(__name__)

model = LlamaForCausalLM.from_pretrained("path/to/llama/model")
tokenizer = LlamaTokenizer.from_pretrained("path/to/llama/tokenizer")

@app.route('/generate', methods=['POST'])
def generate():
    data = request.json
    input_text = data['text']
    inputs = tokenizer(input_text, return_tensors='pt')
    outputs = model.generate(**inputs)
    response = tokenizer.decode(outputs[0])
    return jsonify({"response": response})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)



'''
// used for checking if the request is beeing processed
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/generate', methods=['POST'])
def generate():
    data = request.json
    text = data.get('text', '')
    response = f"Processed text: {text}"  # Simulate Llama response
    return jsonify({"response": response})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
'''