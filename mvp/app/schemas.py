from pydantic import BaseModel, Field
from typing import List, Optional

class SmartReplyRequest(BaseModel):
    title: Optional[str] = Field(None, description="Заголовок сообщения (опционально)")
    text: str = Field(..., description="Текст сообщения")

class SmartReplyResponse(BaseModel):
    classification: str
    summary: str
    risk_level: str
    sla_deadline: str
    recommended_reply: str = Field(..., description="Лучший предложенный ответ")