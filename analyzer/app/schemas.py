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