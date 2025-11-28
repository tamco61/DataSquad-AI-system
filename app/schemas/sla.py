from pydantic import BaseModel

class SLA(BaseModel):
    id: str
    name: str
    description: str
