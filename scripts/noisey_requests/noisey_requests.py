import requests
import time
import random

url = "http://localhost:8080/rfid/collect"

untracked_ids = [{"id": "d1"}, {"id": "f5"}, {"id": "g6"}, {"id": "h7"}, {"id": "j8"},
                 {"id": "k9"}, {"id": "c2"}]
tracked_ids = [{"id": "a3b2d1"}, {"id": "d6e4f7"}, {"id": "z1x3c8"}, {"id": "j5k7l2"},
               {"id": "m9n0o4"}, {"id": "p8q3r1"}, {"id": "u5v2w9"}]

num_requests = 400

interval = 0.02

def spam_server():
    for i in range(num_requests):
        
        untracked_id = random.choice(untracked_ids)
        params = untracked_id

        
        if random.random() < 0.1:  
            tracked_id = random.choice(tracked_ids)
            params = tracked_id

        try:
            
            response = requests.post(url, params=params)
            print(f"Request {i + 1} Status: {response.status_code}, Response: {response.text}")
        except requests.exceptions.RequestException as e:
            print(f"Error on request {i + 1}: {e}")
        
        
        time.sleep(interval)

if __name__ == "__main__":
    spam_server()
