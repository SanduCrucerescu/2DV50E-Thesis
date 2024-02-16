import os
# from openai import OpenAI

# client = OpenAI()

# res = client.files.create(
#     file=open('./training-files/file3.jsonl', "rb"),
#     purpose="fine-tune"
# )

# print(res)

rootDir = "./training-files/"
outputFile = open("./training-files/file3.jsonl", "a")

def append_to_file(c, d):
      outputFile.write("{\"messages\": [{\"role\": \"system\", \"content\": \"You are an experienced programer in delphi and C# who is able to translate code from one language into another.\"}, "
                        "{\"role\": \"assistant\", \"content\": \"Translate the following Delphi code in C#:"+ d +"\"}, "
                        "{\"role\": \"user\", \"content\": \"Here is the translation of the code from Delphi to C#:"+ c +"\"}]}\n")

def process_code(code):
    print(code)

def process_files():
    for dir in  os.listdir(rootDir):
        dir_path = os.path.join(rootDir, dir)

        if os.path.isdir(dir_path):
            csharp = ""
            delphi = ""
            for filename in os.listdir(rootDir + dir):

                if not filename.startswith("."):
                    file_path = os.path.join(dir_path, filename)

                    with open(file_path, "r") as file:
                        f = file.read()
                        lines = file.readlines()

                    file_ext = filename.split(".")[1]

                    if file_ext == "cs":
                         csharp = "".join(f.splitlines()).replace('"', '\\"')
                    elif file_ext == "dpr":
                        delphi = "".join(f.splitlines()).replace('"', '\\"')
            append_to_file(csharp, delphi)      

process_files()