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

from pydantic import BaseModel
from typing import Literal

class AnalyzeRequest(BaseModel):
    title: str | None = None
    text: str

class AnalyzeResponse(BaseModel):
    classification: Literal["spam", "complaint", "request", "incident", "question", "other"]
    summary: str
    risk_level: Literal["low", "medium", "high", "critical"]
    sla_deadline: str  # ISO 8601