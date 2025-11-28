from pydantic import BaseModel


class PositionDTO(BaseModel):
    id: str
    name: str
    description: str


class SLADTO(BaseModel):
    id: str
    name: str
    description: str


class DepartamentDTO(BaseModel):
    id: str
    name: str
    description: str


class TaskDTO(BaseModel):
    id: str
    title: str
    sla_id: str
    type_id: str
    departament_id: str
    attached_id: str


class TypeDTO(BaseModel):
    id: str
    name: str
    description: str


class UserDTO(BaseModel):
    id: str
    email: str
    telegram: str
    name: str
    departament_id: str
    position_id: str