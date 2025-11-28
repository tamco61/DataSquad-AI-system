from openai import OpenAI


class LLM:
    def __init__(self, model, base_url, api_key, folder_id):
        self.model = model
        self.base_url = base_url
        self.api_key = api_key
        self.client = OpenAI(base_url=base_url, api_key=api_key, project=folder_id)
        print(1)

    def request(self, text, instruction):
        print(2)
        resp = self.client.responses.create(model=self.model, instructions=instruction, input=text)
        print(3)
        return resp.output_text

from dotenv import load_dotenv
import os
load_dotenv()

folder_id = os.getenv("folder_id")
api_key = os.getenv("api_key")

model = f"gpt://{folder_id}/qwen3-235b-a22b-fp8/latest"
instruction = '''Ты — эксперт по классификации и приоритизации входящих обращений в службу поддержки / техподдержку / сервис-деск.

Твоя задача — строго по формату JSON вернуть анализ одного сообщения.

Правила классификации:
- spam — явная реклама, фишинг, бессмысленный набор символов, повторяющиеся сообщения
- complaint — жалоба, негатив, претензия, гнев клиента
- request — запрос на выполнение действия (добавить, изменить, предоставить доступ и т.д.)
- incident — сообщение о том, что что-то сломалось, не работает, ошибка
- question — обычный вопрос, просьба разъяснить
- other — всё, что не попадает под вышеуказанное

Уровень риска:
- low — обычный вопрос или запрос, не влияет на бизнес
- medium — жалоба, недовольство, может повлиять на репутацию
- high — инцидент, который уже влияет на работу клиента или системы
- critical — массовый сбой, утечка данных, юридические/финансовые риски, угрозы

SLA deadline — срок реакции. Учитывай риск и тип:
- critical → 15–30 минут
- high → 1–2 часа
- medium → 4–8 часов
- low → 24–48 часов

Ответ должен быть ТОЛЬКО валидным JSON без какого-либо другого текста, markdown или пояснений.
Формат:

{
  "classification": "spam|complaint|request|incident|question|other",
  "summary": "Краткое содержание обращения одной фразой (до 100 символов)",
  "risk_level": "low|medium|high|critical",
  "sla_deadline": "2025-11-28T15:30:00+03:00"   // ISO 8601 с таймзоной Москвы
}'''
input = "С вчерашнего вечера не могу оплатить заказ картой Visa. Постоянно ошибка 5003. Это уже третий раз за месяц!"
ai = LLM(model=model, base_url="https://rest-assistant.api.cloud.yandex.net/v1", api_key=api_key, folder_id=folder_id)
resp = ai.request(input, instruction)
print(resp)