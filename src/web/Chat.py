from fastapi import APIRouter
from model import Chat
from service import Nova

router = APIRouter(prefix="/chat")


@router.post("/")
async def echo(message: Chat.Message) -> str:
    return await Nova.nova_number(message)
