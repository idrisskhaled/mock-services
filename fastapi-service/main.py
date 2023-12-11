from fastapi import FastAPI

app = FastAPI()

@app.get("/api/provision-resources")
def provision_resources():
    return {"message": "Resources provisioned"}
