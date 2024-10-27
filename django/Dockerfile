FROM python:3.11

SHELL ["/bin/bash", "-c"]

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONBUFFERED=1

RUN pip install --upgrade pip

# Обновим список пакетов и установим зависимости
RUN apt-get update && \
    apt-get install -y \
    pkg-config \
    python3-dev \
    default-libmysqlclient-dev \
    build-essential \
    tzdata

WORKDIR /app/backend

COPY ./requirements.txt .
RUN pip install -r requirements.txt

COPY . .

USER root

ENTRYPOINT ["sh", "./entrypoint.sh"]