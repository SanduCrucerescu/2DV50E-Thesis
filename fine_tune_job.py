from openai import OpenAI

client = OpenAI()

res = client.fine_tuning.jobs.create(
    training_file='file-drwKU8DWSQ8iL7MXUVioBO6r',
    model="gpt-3.5-turbo"
)

print(res)