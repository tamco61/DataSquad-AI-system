from pydantic import BaseModel


class Type(BaseModel):
    id: str
    name: str
    description: str