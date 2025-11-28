from .services.analyzer import analyze_message
from fastapi import APIRouter
from .services.generator import generate_drafts
from .services.llm import LLM
from .schemas import SmartReplyResponse, SmartReplyRequest
from dotenv import load_dotenv
import os
load_dotenv()
import json

folder_id = os.getenv("folder_id")
api_key = os.getenv("api_key")

model = f"gpt://{folder_id}/qwen3-235b-a22b-fp8/latest"
router = APIRouter()
llm = LLM(model=model, base_url="https://rest-assistant.api.cloud.yandex.net/v1", api_key=api_key, folder_id=folder_id)

@router.post("/smart/reply", response_model=SmartReplyResponse)
async def smart_reply(request: SmartReplyRequest):
    analysis = analyze_message(llm, title=request.title, text=request.text)
    analy = json.loads(analysis)
    context = f"""Классификация: {analy['classification']}
Краткое содержание: {analy['summary']}
Уровень риска: {analy['risk_level']}
SLA: {analy['sla_deadline']}

Оригинальное сообщение: {request.text}

Сгенерируй вежливый профессиональный ответ."""

    recommended = generate_drafts(llm, context)

    return SmartReplyResponse(
        classification=analy["classification"],
        summary=analy["summary"],
        risk_level=analy["risk_level"],
        sla_deadline=analy["sla_deadline"],
        recommended_reply=recommended
    )