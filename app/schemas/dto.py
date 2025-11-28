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
    point_id: str
    text_id: str
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


class PointDTO(BaseModel):
    id: str
    point: str
    text_id: str


class TextDTO(BaseModel):
    id: str
    text: str
    author_id: str


class AuthorDTO(BaseModel):
    id: str
    email: str
    name: str
    description: str
    contact: str

