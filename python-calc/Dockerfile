FROM python:3.11-slim AS builder
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir --target=/app/deps -r requirements.txt

FROM python:3.11-slim
RUN useradd -u 10001 appuser
USER appuser

WORKDIR /app
COPY --from=builder /app/deps /usr/local/lib/python3.11/site-packages
COPY main.py .

EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]