from fastapi import APIRouter

router = APIRouter()


# crud tasks
@router.get("/task", response_model=None)
async def get_task(payload):
    pass


@router.post("/task", response_model=None)
async def add_task(payload):
    pass


@router.get("/tasks", response_model=None)
async def get_tasks(payload):
    pass


@router.post("/user", response_model=None)
async def add_user(payload):
    pass

@router.get("/user", response_model=None)
async def get_user(payload):
    pass
