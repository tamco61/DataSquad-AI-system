from fastapi import FastAPI
from .router import router

app = FastAPI(
    title="AI Message Analyzer",
    version="1.0.0",
    description="Классификация и анализ входящих сообщений с помощью OpenRouter"
)

app.include_router(router)

@app.get("/")
async def root():
    return {"message": "AI Analyzer service is running"}