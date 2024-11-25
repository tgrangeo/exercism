import subprocess
from dotenv import load_dotenv
import requests
import os

tracks = ["go", "java", "javascript", "php", "purescript", "python", "rust", "typescript"]
README_FILE = 'README.md'

load_dotenv()

API_TOKEN = os.getenv("API_TOKEN")
BASE_URL = "https://exercism.org/api/v2/tracks"

def get_tracks_data():
    headers = {
        "Authorization": f"Bearer {API_TOKEN}"
    }
    response = requests.get(BASE_URL, headers=headers)
    
    if response.status_code == 200:
        data = response.json()
        return data.get("tracks", [])
    else:
        print(f"Failed to retrieve tracks data with status code {response.status_code}")
        return []

def build_exercises_dict():
    exercises_dict = {}
    data = get_tracks_data()
    for track in data:
        slug = track.get("slug")
        num_exercises = track.get("num_exercises", 0)
        exercises_dict[slug] = num_exercises
    return exercises_dict

def remove_section_content():
    # Read the existing README content
    with open(README_FILE, 'r') as file:
        content = file.readlines()
    
    # Find the indexes of the sections
    start_idx = None
    end_idx = None
    for i, line in enumerate(content):
        if line.strip() == "## **Languages & Tracks**":
            start_idx = i
        if line.strip() == "## **Scripts**":
            end_idx = i
            break
    
    # If both sections are found, remove content in between (but keep the headers)
    if start_idx is not None and end_idx is not None:
        content = content[:start_idx + 1] + content[end_idx:]
    
    # Write the updated content back to the README
    with open(README_FILE, 'w') as file:
        file.writelines(content)

def update_languages_section(new_data):
    with open(README_FILE, 'r') as file:
        content = file.readlines()
    
    start_idx = None
    for i, line in enumerate(content):
        if line.strip() == "## **Languages & Tracks**":
            start_idx = i + 1
            break
    
    if start_idx is not None:
        end_idx = start_idx
        while end_idx < len(content) and content[end_idx].startswith('-'):
            end_idx += 1
        new_language_lines = [f"{line}\n" for line in new_data]
        content = content[:start_idx] + ["\n"] + new_language_lines+ ["\n"] + content[end_idx:]
    
    with open(README_FILE, 'w') as file:
        file.writelines(content)

def main():
    new_data = []
    d = build_exercises_dict()
    for t in tracks:
        res = subprocess.run(f"find {t} -type d -mindepth 1 -maxdepth 1 | wc -l", shell=True, capture_output=True, text=True)
        folder_count = int(res.stdout.strip())
        new_data.append(f"- **{t.capitalize()}**: {folder_count} / {d.get(t, 0)} ({round(folder_count * 100 / d.get(t, 0))}%)")
    remove_section_content()
    update_languages_section(new_data)

if __name__ == "__main__":
    main()
