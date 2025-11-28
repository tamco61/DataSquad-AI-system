from fastapi import APIRouter, Depends, HTTPException
from uuid import UUID
from .schemas import GenerateDraftRequest, GenerateDraftResponse
from .services.draft_generator import generate_drafts
from .dependencies import get_task_context

router = APIRouter()


@router.post("/draft/generate", response_model=GenerateDraftResponse)
async def generate_draft(payload: GenerateDraftRequest):
    task_id = str(payload.task_id)
    context = get_task_context(task_id)

    if "не найден" in context:
        raise HTTPException(status_code=404, detail="Task not found")

    drafts = await generate_drafts(context)

    return GenerateDraftResponse(drafts=drafts)