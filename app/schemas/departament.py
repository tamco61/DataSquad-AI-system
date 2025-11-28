from pydantic import BaseModel


class Departament(BaseModel):
    id: str
    name: str
    description: str
