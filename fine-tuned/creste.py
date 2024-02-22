from openai import OpenAI
client = OpenAI(
    api_key="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"
)

response=client.fine_tuning.jobs.create(
  training_file="file-ITYmkfbg3DVdF7GY59ZV14sb", 
  model="gpt-3.5-turbo"
)

print(response)