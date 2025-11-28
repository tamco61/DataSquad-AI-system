from pydantic import BaseModel


class Position(BaseModel):
    id: str
    name: str
    description: str