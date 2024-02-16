import json
from openai import OpenAI

client = OpenAI()

# completion = client.chat.completions.create(
#     model="gpt-3.5-turbo",
#     messages=[
#         {"role":"system", "content": "You are a poetic assistant, skilled in explaining complex programming concepts with creative flair."},
#         {"role":"user", "content": "Compose a poem that explains the concept of recursion in programming."}
#     ]
# )

# file-2DjA1WUWHeGH6ynYa6zjv8ug

# print(completion.choices[0].message)
data_path = "./training-files/file1.jsonl"
with open(data_path, 'r', encoding='utf-8') as f:
    dataset = [json.loads(line) for line in f]

res = client.files.create(
    file=open('./training-files/file1.jsonl', "rb"),
    purpose="fine-tune"
)

# res = client.fine_tuning.jobs.create(
#     training_file='file-2DjA1WUWHeGH6ynYa6zjv8ug',
#     model="gpt-3.5-turbo"
# )

print(res)

