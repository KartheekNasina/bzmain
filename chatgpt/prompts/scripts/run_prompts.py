import openai
import os
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Set up the OpenAI API key from environment variable
openai.api_key = os.getenv('OPENAI_API_KEY')

def send_prompt_to_openai(prompt):
    """Send a prompt to OpenAI and get the response."""
    response = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": "You are a go lang coding master helping me write go lang code."},
            {"role": "user", "content": prompt}
        ]
    )
    return response.choices[0].message['content'].strip()

def main():
    # Read prompts from the input file
    with open('input_prompts.txt', 'r') as f:
        prompts = f.readlines()

    # Collect responses for each prompt
    responses = []
    for prompt in prompts:
        response = send_prompt_to_openai(prompt.strip())
        responses.append(response)

    # Write the responses to the output file
    with open('output_responses.txt', 'w') as f:
        for response in responses:
            f.write(response + "\n")

if __name__ == '__main__':
    main()
