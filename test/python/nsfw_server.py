import os
os.environ['TF_CUDA_CLANG'] = '0'
os.environ['CUDA_VISIBLE_DEVICES'] = '-1'  # 告诉 TensorFlow 你不想使用 GPU

# 违规图片识别

from flask import Flask, request, jsonify
from nsfw_detector import predict

model = predict.load_model('./nsfw_model.h5')  # 预加载模型
app = Flask(__name__)

@app.route('/predict', methods=['POST'])
def predict_nsfw():
    file = request.files['image']
    predictions = predict.classify(model, file)
    return jsonify(predictions)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)