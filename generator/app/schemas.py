from pydantic import BaseModel, Field
from uuid import UUID
from typing import List

class DraftResponse(BaseModel):
    style: str = Field(..., description="Стиль ответа")
    content: str = Field(..., description="Сгенерированный текст черновика")

class GenerateDraftRequest(BaseModel):
    task_id: UUID

class GenerateDraftResponse(BaseModel):
    drafts: List[DraftResponse]