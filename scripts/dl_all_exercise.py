import requests
import subprocess
from dotenv import load_dotenv
import os
from tqdm import tqdm

load_dotenv()

API_TOKEN = os.getenv("API_TOKEN")
USERNAME = os.getenv("USERNAME")

def fetch_solutions(page):
    url = f"https://exercism.org/api/v2/solutions?user_id={USERNAME}&page={page}"
    headers = {
        "Authorization": f"Bearer {API_TOKEN}"
    }
    response = requests.get(url, headers=headers)
    
    if response.status_code != 200:
        print(f"\nRequest failed with status code {response.status_code}")
        return None
    return response.json()

def download_exercise(exercise):
    if exercise['status'] == 'published':
        command = [
            "exercism", "download", 
            f"--track={exercise['track']['slug']}", 
            f"--exercise={exercise['exercise']['slug']}"
        ]
        subprocess.run(command, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)

def calculate_total_exercises(total_pages):
    total_exercises = 0
    for page in range(total_pages + 1):
        data = fetch_solutions(page)
        if data:
            total_exercises += len(data["results"])
    return total_exercises

def main():
    page = 0
    initial_data = fetch_solutions(page)
    
    if not initial_data:
        return
    
    total_pages = initial_data["meta"]["total_pages"]
    
    total_exercises = calculate_total_exercises(total_pages)
    
    with tqdm(total=total_exercises, desc="Downloading Exercises", unit="exercise") as progress_bar:
        while page <= total_pages:
            data = fetch_solutions(page)
            if data:
                for exercise in data["results"]:
                    download_exercise(exercise)
            page += 1

    print("\nDownload complete!")

if __name__ == "__main__":
    main()
