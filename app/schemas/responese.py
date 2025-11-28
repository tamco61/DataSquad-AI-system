from pydantic import BaseModel

class Response(BaseModel):
    pass

class TaskResponse(Response):
    id: str
    title: str
    point_id: str
    sla_id: str
    type_id: str
    departament_id: str
    attached_id: str