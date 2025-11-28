
instruction = '''
    Ты должен сгенерировать ответ на сообщение используя следующий контекст, ответ должен быть официальным письмом, 
    которое можно будет сразу же отправить

'''

def generate_drafts(llm, context=None):
    if context == None:
        return

    resp = llm.request(context, instruction)
    return resp