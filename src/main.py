from fastapi import FastAPI
from web import Chat

app = FastAPI()

app.include_router(Chat.router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", reload=True)