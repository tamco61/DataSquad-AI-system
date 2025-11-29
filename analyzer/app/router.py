from fastapi import APIRouter, Depends
from .schemas import AnalyzeRequest, AnalyzeResponse
from .services import analyze_message

router = APIRouter(tags=["AI Analysis"])

@router.post("/ai/analyze", response_model=AnalyzeResponse)
async def analyze_endpoint(
    request: AnalyzeRequest):
    result = analyze_message.analyze_message(
        title=request.title,
        text=request.text
    )
    return AnalyzeResponse(**result)
