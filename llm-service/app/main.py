from fastapi import FastAPI
from .router import router

app = FastAPI(
    title="LLM Service",
    version="1.0.0",
    description="Предоставляет функционал генерации ответа и анализа сообщений"
)

app.include_router(router)

@app.get("/")
async def root():
    return {
        "message": "LLM service up!"
    }
