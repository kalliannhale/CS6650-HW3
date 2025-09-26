# locustfile.py - A Grimoire Specifically for the Go/Gin Mansion

from locust import FastHttpUser, task, between
import random

class GoGinHaunter(FastHttpUser):
    # The spirits will wait 1 to 2 seconds between their spectral actions.
    wait_time = between(1, 2)

    # This spell defines what a spirit does upon first being summoned.
    # It's a one-time setup ritual.
    def on_start(self):
        """ On start, each user will log in and get a specific album. """
        # Let's start with a known list of album IDs from your Go code.
        self.album_ids = ["1", "2", "3"]
        print("A new spirit has manifested and learned the mansion's secrets.")

    # A GET request spell to retrieve a known album by its ID.
    # A reliable, predictable haunting.
    @task(3) # This task is 3 times more likely to be run than the POST task.
    def get_specific_album(self):
        # Choose one of the known IDs at random.
        album_id = random.choice(self.album_ids)
        
        # The endpoint must exactly match what's in your main.go: /albums/:id
        url = f"/albums/{album_id}"
        
        print(f"A spirit investigates the history of album ID: {album_id}")
        self.client.get(url, name="/albums/[id]") # We use a name for clean aggregation in the UI

    # A POST request spell. This spirit adds a new album to the collection.
    # An active, creative haunting, like Ayesha Erotica dropping a new track.
    @task(1)
    def create_new_album(self):
        # The JSON payload must EXACTLY match the `album` struct in your Go code.
        # Notice the keys: "id", "title", "artist", "price". Case-sensitive!
        # Your Go server will add this to its in-memory list.
        new_album_payload = {
            "id": f"locust-{random.randint(100, 999)}", # Create a semi-unique ID
            "title": "Vroom Vroom EP",
            "artist": "Charli XCX",
            "price": 13.37
        }

        # The endpoint for creating a new album is /albums
        url = "/albums"

        print(f"A spirit adds a new record to the collection: {new_album_payload['title']}")
        self.client.post(url, json=new_album_payload)
