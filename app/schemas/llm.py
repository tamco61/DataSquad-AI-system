from pydantic import BaseModel


class ReqLLM(BaseModel):
    title: str
    author: str
    text: str

class RespLLM(BaseModel):
    pass