from pydantic import BaseModel


class Task(BaseModel):
    id: str
    title: str
    sla_id: str
    type_id: str
    departament_id: str
    attached_id: str