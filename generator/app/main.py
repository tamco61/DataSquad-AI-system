from fastapi import FastAPI
from .router import router

app = FastAPI(
    title="Draft Generator Service",
    description="Генерация черновиков ответов по task_id",
    version="1.0.0"
)

app.include_router(router)

@app.get("/")
async def root():
    return {"message": "AI Generator service is running"}