from pydantic import BaseModel

class User(BaseModel):
    id: str
    email: str
    telegram: str
    name: str
    departament_id: str
    position_id: str