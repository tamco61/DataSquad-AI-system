from fastapi import APIRouter, Depends, HTTPException
from uuid import UUID
from .schemas import GenerateDraftRequest, GenerateDraftResponse
from .services.generator import generate_drafts
from .dependencies import get_task_context
from .services.llm import LLM

from dotenv import load_dotenv
import os
load_dotenv()

folder_id = os.getenv("folder_id")
api_key = os.getenv("api_key")

model = f"gpt://{folder_id}/qwen3-235b-a22b-fp8/latest"
router = APIRouter()
llm = LLM(model=model, base_url="https://rest-assistant.api.cloud.yandex.net/v1", api_key=api_key, folder_id=folder_id)

@router.post("/draft/generate", response_model=GenerateDraftResponse)
async def generate_draft(payload: GenerateDraftRequest):
    task_id = str(payload.task_id)
    context = get_task_context(task_id)

    if "не найден" in context:
        raise HTTPException(status_code=404, detail="Task not found")

    drafts = generate_drafts(llm, context)

    return GenerateDraftResponse(drafts=drafts)

from fastapi import APIRouter, Depends
from .schemas import AnalyzeRequest, AnalyzeResponse
from .services.analyzer import analyze_message

router = APIRouter(tags=["AI Analysis"])

@router.post("/ai/analyze", response_model=AnalyzeResponse)
async def analyze_endpoint(
    request: AnalyzeRequest):
    result = analyze_message(llm,
        title=request.title,
        text=request.text
    )
    return AnalyzeResponse(**result)