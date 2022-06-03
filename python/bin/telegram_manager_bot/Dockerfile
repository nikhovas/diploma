FROM python:3.9
WORKDIR /src
COPY . .
WORKDIR /src/bin/telegram_manager_bot
ENV PYTHONPATH=${PYTHONPATH}:/src/bin/telegram_manager_bot
RUN pip3 install poetry
RUN poetry config virtualenvs.create false
RUN poetry install --no-dev
ENTRYPOINT python3 ./telegram_manager_bot/__main__.py
