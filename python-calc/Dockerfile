FROM python:3.12-alpine AS builder
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir --target=/app/deps -r requirements.txt

FROM python:3.12-alpine
# 1. PERBAIKAN: Gunakan adduser (sintaksis khusus Alpine Linux) dengan ID 10001
RUN adduser -D -u 10001 appuser

WORKDIR /app

# 2. Salin dependensi ke path internal python bawaan
COPY --from=builder /app/deps /usr/local/lib/python3.12/site-packages

# 3. PERBAIKAN: Ubah hak milik file main.py ke appuser agar bisa diakses dengan aman
COPY --chown=appuser:appuser main.py .

# 4. Pindahkan instruksi USER ke bagian paling bawah sebelum port dibuka
USER appuser

EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
