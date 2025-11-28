from fastapi import FastAPI
from .router import router

app = FastAPI(title="Smart Reply Service")

app.include_router(router)

@app.get("/")
async def root():
    return {"message": "Smart Reply готов!"}