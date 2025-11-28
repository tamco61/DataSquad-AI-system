from openai import OpenAI


class LLM:
    def __init__(self, model, base_url, api_key, folder_id):
        self.model = model
        self.base_url = base_url
        self.api_key = api_key
        self.client = OpenAI(base_url=base_url, api_key=api_key, project=folder_id)

    def request(self, text, instruction):
        resp = self.client.responses.create(model=self.model, instructions=instruction, input=text)
        return resp.output_text
