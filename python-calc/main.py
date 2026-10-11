from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def add(a: int, b: int):
    return {"result": a + b}