from pydantic import BaseModel

class ReqTask(BaseModel):
    task_id: str
