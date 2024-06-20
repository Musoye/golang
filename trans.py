import os
import assemblyai as aai


API_KEY = '65e24ac9eca943b7919cd298ad783339'


def save_transcription(file_path, transcription):
    txt_file_path = os.path.splitext(file_path)[0] + ".txt"
    with open(txt_file_path, 'w') as f:
        f.write(transcription)
    print(f"Transcription saved to {txt_file_path}")

# Replace with your API key
aai.settings.api_key = ""


for i in range(4, 12):
    FILE_URL = f"./{i}.wav"

    transcriber = aai.Transcriber()
    transcript = transcriber.transcribe(FILE_URL)

    if transcript.status == aai.TranscriptStatus.error:
        print(transcript.error)
    else:
        url = f'{i}.txt'
        save_transcription(url, transcript.text)
        print(f'Transcription saved to {url}')

