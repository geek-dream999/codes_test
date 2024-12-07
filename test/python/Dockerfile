# 使用 Python 3.8 作为基础镜像
FROM python:3.8

# 配置pip使用国内源
RUN pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/

# 安装protobuf
RUN pip install protobuf==3.20.0

# 安装TensorFlow CPU版本和其他依赖
RUN pip install tensorflow-cpu==2.8 tensorflow-hub==0.7.0 nsfw-detector flask

# 设置环境变量
ENV TF_CUDA_CLANG=0
ENV PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION=python
ENV CUDA_VISIBLE_DEVICES=-1

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到工作目录中
COPY . /app

# 检查模型文件是否存在
RUN if [ ! -f /app/nsfw_model.h5 ]; then echo "nsfw_model.h5 not found"; exit 1; fi

# 启动 Flask 应用
CMD ["python", "nsfw_server.py"]